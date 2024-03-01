package common

import (
	"context"
	"errors"
	"github.com/qiniu/qmgo"
	"pchat/repository"
	"pchat/repository/bson"
	"pchat/utils"
	"time"
)

const (
	C_SETTING = "setting"

	DEFAULT_ACCESS_TOKEN_EXPIRED_SECOND = 3600 * 24
)

var (
	CSetting = &Setting{}

	cachedSetting *Setting
)

type Setting struct {
	Id        bson.ObjectId  `bson:"_id"`
	UpdatedAt time.Time      `bson:"updatedAt"`
	SMTP      SMTPSetting    `bson:"smtp,omitempty"`
	OSS       OSSSetting     `bson:"oss,omitempty"`
	AI        AISetting      `bson:"ai,omitempty"`
	Chat      ChatSetting    `bson:"chat,omitempty"`
	Account   AccountSetting `bson:"account,omitempty"`
}

type AISetting struct {
	Provider  string `bson:"provider"`
	Proxy     string `bson:"proxy"`
	IsEnabled bool   `bson:"isEnabled"`
	Model     string `bson:"model"`
	Key       string `bson:"key"`
}

type SMTPSetting struct {
	Host       string `bson:"host"`
	Protocol   string `bson:"protocol"`
	Port       int    `bson:"port"`
	Username   string `bson:"username"`
	Password   string `bson:"password"`
	SenderName string `bson:"senderName"`
}

type AccountSetting struct {
	Register         RegisterSetting `bson:"register"`
	Password         PasswordSetting `bson:"password"`
	TokenValidSecond int64           `bson:"TokenValidSecond"`
	TokenKey         string          `bson:"tokenKey"`
}

type PasswordSetting struct {
	IsEnabled          bool  `bson:"isEnabled"`
	MinLength          int64 `bson:"minLength"`
	MaxLength          int64 `bson:"maxLength"`
	MustHasLowerCase   bool  `bson:"mustHasLowerCase"`
	MustHasUpperCase   bool  `bson:"mustHasUpperCase"`
	MustHasNumber      bool  `bson:"mustHasNumber"`
	MustHasSpecialCode bool  `bson:"mustHasSpecialCode"`
}

type RegisterSetting struct {
	MustBeApprovedBeforeRegister bool `bson:"mustBeApprovedBeforeRegister"`
}

type OSSSetting struct {
	Provider        string `bson:"provider"`
	PublicBucket    string `bson:"publicBucket"`
	PrivateBucket   string `bson:"privateBucket"`
	ValidSecond     int64  `bson:"validSecond"`
	AccessKey       string `bson:"accessKey"`
	SecretAccessKey string `bson:"secretAccessKey"`
	Endpoint        string `bson:"endpoint"`
}

type ChatSetting struct {
	ShowMessageReadStatus bool `bson:"showMessageReadStatus"`
	AllowRollback         bool `bson:"allowRollback"`
	SendEmailIfNotOnline  bool `bson:"sendEmailIfNotOnline"`
}

func (s *Setting) CreateDefaultSetting(ctx context.Context) error {
	_, err := s.Get(ctx)
	if errors.Is(err, qmgo.ErrNoSuchDocuments) {
		key, err := utils.GenerateRandomSecretKey(64)
		if err != nil {
			return err
		}
		setting := Setting{
			Id:        bson.NewObjectId(),
			UpdatedAt: time.Now(),
			Account: AccountSetting{
				Register: RegisterSetting{
					MustBeApprovedBeforeRegister: true,
				},
				TokenValidSecond: DEFAULT_ACCESS_TOKEN_EXPIRED_SECOND,
				TokenKey:         key,
			},
		}
		return repository.Insert(ctx, C_SETTING, setting)
	}
	return nil
}

func (*Setting) Get(ctx context.Context) (Setting, error) {
	s := Setting{}
	err := repository.FindOne(ctx, C_SETTING, nil, &s)
	return s, err
}

func (s *Setting) Update(ctx context.Context) error {
	condition := bson.M{
		"_id": s.Id,
	}
	updater := bson.M{
		"$set": bson.M{
			"updatedAt": time.Now(),
			"smtp":      s.SMTP,
			"oss":       s.OSS,
			"ai":        s.AI,
			"chat":      s.Chat,
			"account":   s.Account,
		},
	}
	return repository.UpdateOne(ctx, C_SETTING, condition, updater)
}

func (*Setting) GetWithCache(ctx context.Context) (*Setting, error) {
	if cachedSetting != nil {
		return cachedSetting, nil
	}
	setting, err := CSetting.Get(ctx)
	if err != nil {
		return nil, err
	}
	cachedSetting = &setting
	time.AfterFunc(time.Minute, func() {
		cachedSetting = nil
	})
	return cachedSetting, nil
}
