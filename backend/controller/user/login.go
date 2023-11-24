package user

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	c_utils "pchat/controller/utils"
	"pchat/model"
	pb_user "pchat/pb/user"
)

func login(ctx *gin.Context, req *pb_user.LoginRequest) (*pb_user.LoginResponse, error) {
	user, err := model.CUser.GetByEmail(ctx, req.Email, true)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	token, err := model.SignToken(ctx, user)
	if err != nil {
		return nil, err
	}
	return &pb_user.LoginResponse{
		Token: token,
	}, nil
}

var LoginController = c_utils.NewGinController[*pb_user.LoginRequest, *pb_user.LoginResponse](login)
