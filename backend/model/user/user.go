package user

import (
	"context"
	"crypto/rand"
	"encoding/base32"
	"github.com/PPG007/gotp"
	"github.com/google/uuid"
	"github.com/qiniu/qmgo"
	"golang.org/x/crypto/bcrypt"
	"pchat/model"
	"pchat/repository"
	"pchat/repository/bson"
	"pchat/utils"
	"pchat/utils/env"
	"time"
)

const (
	C_USER = "user"

	USER_STATUS_ACTIVATED = "activated"
	USER_STATUS_BLOCKED   = "blocked"
	USER_STATUS_AUDITING  = "auditing"

	USER_CHAT_STATUS_ONLINE  = "online"
	USER_CHAT_STATUS_OFFLINE = "offline"
	USER_CHAT_STATUS_LEAVING = "leaving"
	USER_CHAT_STATUS_BUSY    = "busy"
)

var (
	CUser = &User{}
)

type User struct {
	Id            bson.ObjectId   `bson:"_id"`
	Name          string          `bson:"name"`
	Password      string          `bson:"password"`
	Email         string          `bson:"email"`
	Roles         []bson.ObjectId `bson:"roles"`
	CreatedAt     time.Time       `bson:"createdAt"`
	UpdatedAt     time.Time       `bson:"updatedAt"`
	Status        string          `bson:"status"`
	Avatar        string          `bson:"avatar"`
	ChatStatus    string          `bson:"chatStatus"`
	Is2FAEnabled  bool            `bson:"is2FAEnabled"`
	OTPSecret     string          `bson:"otpSecret"`
	RecoveryCodes []string        `bson:"recoveryCodes"`
}

func (*User) GetById(ctx context.Context, id bson.ObjectId) (User, error) {
	var user User
	condition := model.GenIdCondition(id)
	err := repository.FindOne(ctx, C_USER, condition, &user)
	return user, err
}

func (*User) GetByEmail(ctx context.Context, email string, onlyActivated bool) (User, error) {
	var user User
	condition := bson.M{
		"email": email,
	}
	if onlyActivated {
		condition["status"] = USER_STATUS_ACTIVATED
	}
	err := repository.FindOne(ctx, C_USER, condition, &user)
	return user, err
}

func (*User) CreateNew(ctx context.Context, email, password, reason string, needAudit bool) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := User{
		Id:       bson.NewObjectId(),
		Name:     "user",
		Password: string(hashed),
		Email:    email,
		// TODO: add default roles
		Roles:     nil,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status: func() string {
			if needAudit {
				return USER_STATUS_AUDITING
			}
			return USER_STATUS_ACTIVATED
		}(),
	}
	err = repository.Insert(ctx, C_USER, user)
	if err != nil {
		return err
	}
	if needAudit {
		return CRegisterApplication.Upsert(ctx, user, reason)
	}
	return nil
}

func (*User) UpdateById(ctx context.Context, id bson.ObjectId, updater bson.M) error {
	return repository.UpdateOne(ctx, C_USER, model.GenIdCondition(id), updater)
}

func (*User) Online(ctx context.Context) (User, error) {
	condition := model.GenIdCondition(utils.GetUserIdAsObjectId(ctx))
	change := qmgo.Change{
		Update: bson.M{
			"$set": bson.M{
				"chatStatus": USER_CHAT_STATUS_ONLINE,
			},
		},
		ReturnNew: true,
	}
	user := User{}
	err := repository.FindAndApply(ctx, C_USER, condition, change, &user)
	return user, err
}

func (*User) Offline(ctx context.Context) (User, error) {
	condition := model.GenIdCondition(utils.GetUserIdAsObjectId(ctx))
	change := qmgo.Change{
		Update: bson.M{
			"$set": bson.M{
				"chatStatus": USER_CHAT_STATUS_OFFLINE,
			},
		},
	}
	user := User{}
	err := repository.FindAndApply(ctx, C_USER, condition, change, &user)
	return user, err
}

func (*User) Activate(ctx context.Context, ids []bson.ObjectId) error {
	condition := bson.M{
		"_id": bson.M{
			"$in": ids,
		},
		"status": bson.M{
			"$in": []string{USER_STATUS_BLOCKED, USER_STATUS_AUDITING},
		},
	}
	updater := bson.M{
		"$set": bson.M{
			"updatedAt": time.Now(),
			"status":    USER_STATUS_ACTIVATED,
		},
	}
	_, err := repository.UpdateAll(ctx, C_USER, condition, updater)
	// TODO: send notification
	return err
}

func (*User) Enable2FA(ctx context.Context, id bson.ObjectId) (string, []string, error) {
	rawSecret := make([]byte, 10)
	n, err := rand.Read(rawSecret)
	if err != nil {
		return "", nil, err
	}
	rawSecret = rawSecret[:n]
	secret := base32.StdEncoding.EncodeToString(rawSecret)
	condition := model.GenIdCondition(id)
	codes, _ := CUser.GenerateRecoveryCodes(ctx, id, false)
	change := qmgo.Change{
		ReturnNew: true,
		Update: bson.M{
			"$set": bson.M{
				"otpSecret":     secret,
				"is2FAEnabled":  true,
				"recoveryCodes": codes,
			},
		},
	}
	var user User
	err = repository.FindAndApply(ctx, C_USER, condition, change, &user)
	if err != nil {
		return "", nil, err
	}
	return gotp.NewTOTP(
		gotp.WithSecret(rawSecret),
		gotp.WithLabel(user.Email),
		gotp.WithIssuer(env.GetAppName()),
	).SignURL(), codes, nil
}

func (user *User) VerifyOTP(ctx context.Context, input string) (bool, error) {
	if !user.Is2FAEnabled {
		return true, nil
	}
	if utils.StrInArray(input, &user.RecoveryCodes) {
		err := user.DisableRecoveryCode(ctx, user.Id, input)
		if err != nil {
			return false, err
		}
		if len(user.RecoveryCodes) == 1 {
			utils.GO(ctx, func(innerCtx context.Context) {
				// TODO: send notification
				user.GenerateRecoveryCodes(innerCtx, user.Id, true)
			})
		}
		return true, nil
	}
	secret, err := base32.StdEncoding.DecodeString(user.OTPSecret)
	if err != nil {
		return false, err
	}
	return gotp.NewTOTP(
		gotp.WithSecret(secret),
	).SignPassword() == input, nil
}

func (*User) GenerateRecoveryCodes(ctx context.Context, id bson.ObjectId, doUpdate bool) ([]string, error) {
	length := 10
	codes := make([]string, 0, length)
	for i := 0; i < length; i++ {
		codes = append(codes, uuid.NewString())
	}
	if !doUpdate {
		return codes, nil
	}
	condition := model.GenIdCondition(id)
	updater := bson.M{
		"$set": bson.M{
			"recoveryCodes": codes,
		},
	}
	err := repository.UpdateOne(ctx, C_USER, condition, updater)
	return codes, err
}

func (*User) DisableRecoveryCode(ctx context.Context, id bson.ObjectId, code string) error {
	condition := model.GenIdCondition(id)
	updater := bson.M{
		"$pull": bson.M{
			"recoveryCodes": code,
		},
	}
	return repository.UpdateOne(ctx, C_USER, condition, updater)
}

func (*User) DisableOTP(ctx context.Context, id bson.ObjectId) error {
	condition := model.GenIdCondition(id)
	updater := bson.M{
		"$unset": bson.M{
			"recoveryCodes": "",
			"otpSecret":     "",
		},
		"$set": bson.M{
			"is2FAEnabled": false,
		},
	}
	return repository.UpdateOne(ctx, C_USER, condition, updater)
}

func (*User) GetPermissions(ctx context.Context, id bson.ObjectId) ([]string, error) {
	user, err := CUser.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	if len(user.Roles) == 0 {
		return nil, nil
	}
	return CRole.GetPermissionsByIds(ctx, user.Roles)
}
