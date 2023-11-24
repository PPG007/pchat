package utils

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"pchat/model"
	"time"
)

type UserClaim struct {
	createdAt time.Time
	expiredAt time.Time
	userId    string
	name      string
	email     string
	roleIds   []string
}

func (u UserClaim) GetExpirationTime() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{
		Time: u.expiredAt,
	}, nil
}

func (u UserClaim) GetIssuedAt() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{
		Time: u.createdAt,
	}, nil
}

func (u UserClaim) GetNotBefore() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{
		Time: u.createdAt,
	}, nil
}

func (u UserClaim) GetIssuer() (string, error) {
	return "pchat", nil
}

func (u UserClaim) GetSubject() (string, error) {
	return u.email, nil
}

func (u UserClaim) GetAudience() (jwt.ClaimStrings, error) {
	return []string{}, nil
}

func SignToken(ctx context.Context, user model.User) (string, error) {
	setting, err := model.CSetting.Get(ctx)
	if err != nil {
		return "", err
	}
	c := UserClaim{
		createdAt: time.Now(),
		expiredAt: time.Now().Add(time.Second * time.Duration(setting.AccessTokenSetting.ExpiredSecond)),
		userId:    user.Id.Hex(),
		name:      user.Name,
		email:     user.Email,
		roleIds:   ObjectIdsToStrArray(user.Roles),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return t.SignedString([]byte(setting.AccessTokenSetting.Key))
}

func ValidToken(ctx context.Context, token string) (*UserClaim, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		setting, err := model.CSetting.Get(ctx)
		if err != nil {
			return nil, err
		}
		return []byte(setting.AccessTokenSetting.Key), nil
	})
	if err != nil {
		return nil, err
	}
	if c, ok := t.Claims.(UserClaim); ok {
		return &c, nil
	}
	return nil, errors.New("invalid token")
}
