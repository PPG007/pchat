package model

import (
	"context"
	"github.com/qiniu/qmgo"
	"golang.org/x/crypto/bcrypt"
	"pchat/repository"
	"pchat/repository/bson"
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
	Id         bson.ObjectId   `bson:"_id"`
	Name       string          `bson:"name"`
	Password   string          `bson:"password"`
	Email      string          `bson:"email"`
	Roles      []bson.ObjectId `bson:"roles"`
	CreatedAt  time.Time       `bson:"createdAt"`
	UpdatedAt  time.Time       `bson:"updatedAt"`
	Status     string          `bson:"status"`
	Avatar     string          `bson:"avatar"`
	ChatStatus string          `bson:"chatStatus"`
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

func (*User) CreateNew(ctx context.Context, email, password string, needAudit bool) error {
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
	return repository.Insert(ctx, C_USER, user)
}

func (u *User) Activate(ctx context.Context) error {
	condition := bson.M{
		"_id": u.Id,
	}
	change := qmgo.Change{
		Update: bson.M{
			"$set": bson.M{
				"status":    USER_STATUS_ACTIVATED,
				"updatedAt": time.Now(),
			},
		},
	}
	err := repository.FindAndApply(ctx, C_USER, condition, change, u)
	return err
}
