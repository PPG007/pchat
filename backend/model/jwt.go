package model

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"pchat/utils"
	"time"
)

type UserClaim struct {
	CreatedAt time.Time `json:"createdAt"`
	ExpiredAt time.Time `json:"expiredAt"`
	UserId    string    `json:"userId"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
}

func (u UserClaim) GetExpirationTime() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{
		Time: u.ExpiredAt,
	}, nil
}

func (u UserClaim) GetIssuedAt() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{
		Time: u.CreatedAt,
	}, nil
}

func (u UserClaim) GetNotBefore() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{
		Time: u.CreatedAt,
	}, nil
}

func (u UserClaim) GetIssuer() (string, error) {
	return "pchat", nil
}

func (u UserClaim) GetSubject() (string, error) {
	return u.Email, nil
}

func (u UserClaim) GetAudience() (jwt.ClaimStrings, error) {
	return []string{}, nil
}

func SignToken(ctx context.Context, user User) (string, error) {
	setting, err := CSetting.Get(ctx)
	if err != nil {
		return "", err
	}
	c := UserClaim{
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(time.Second * time.Duration(setting.AccessTokenSetting.ExpiredSecond)),
		UserId:    user.Id.Hex(),
		Name:      user.Name,
		Email:     user.Email,
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return t.SignedString(utils.ParseSecretString(setting.AccessTokenSetting.Key))
}

func ValidToken(ctx context.Context, token string) (*UserClaim, error) {
	t, err := jwt.ParseWithClaims(token, &UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		setting, err := CSetting.Get(ctx)
		if err != nil {
			return nil, err
		}
		return utils.ParseSecretString(setting.AccessTokenSetting.Key), nil
	})
	if err != nil {
		return nil, err
	}
	if c, ok := t.Claims.(*UserClaim); ok {
		return c, nil
	}
	return nil, errors.New("invalid token")
}
