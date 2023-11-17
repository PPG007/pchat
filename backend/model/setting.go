package model

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
)

type Setting struct {
	Id                 bson.ObjectId      `bson:"_id"`
	UpdatedAt          time.Time          `bson:"updatedAt"`
	EmailSetting       EmailSetting       `bson:"emailSetting,omitempty"`
	AccessTokenSetting AccessTokenSetting `bson:"accessTokenSetting,omitempty"`
	OSSSetting         OSSSetting         `bson:"ossSetting,omitempty"`
	OpenAISetting      OpenAISetting      `bson:"openAISetting,omitempty"`
	ChatSetting        ChatSetting        `bson:"chatSetting,omitempty"`
}

type OpenAISetting struct {
	Key       string `bson:"key"`
	Proxy     string `bson:"proxy"`
	IsEnabled bool   `bson:"isEnabled"`
}

type EmailSetting struct {
	Server               string `bson:"server"`
	Port                 int    `bson:"port"`
	Username             string `bson:"username"`
	Password             string `bson:"password"`
	SendEmailIfNotOnline bool   `bson:"sendEmailIfNotOnline"`
}

type AccessTokenSetting struct {
	Key           string `bson:"key"`
	ExpiredSecond int    `bson:"expiredSecond"`
}

type OSSSetting struct {
	Provider      string `bson:"provider"`
	Bucket        string `bson:"bucket"`
	ExpiredSecond int    `bson:"expiredSecond"`
	Url           string `bson:"url"`
}

type ChatSetting struct {
	ShowMessageReadStatus        bool `bson:"showMessageReadStatus"`
	AllowRollback                bool `bson:"allowRollback"`
	MustBeApprovedBeforeRegister bool `bson:"mustBeApprovedBeforeRegister"`
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
			AccessTokenSetting: AccessTokenSetting{
				Key:           key,
				ExpiredSecond: DEFAULT_ACCESS_TOKEN_EXPIRED_SECOND,
			},
			ChatSetting: ChatSetting{
				MustBeApprovedBeforeRegister: true,
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
			"updatedAt":          time.Now(),
			"emailSetting":       s.EmailSetting,
			"accessTokenSetting": s.AccessTokenSetting,
			"ossSetting":         s.OSSSetting,
			"openAISetting":      s.OpenAISetting,
			"chatSetting":        s.ChatSetting,
		},
	}
	return repository.UpdateOne(ctx, C_SETTING, condition, updater)
}
