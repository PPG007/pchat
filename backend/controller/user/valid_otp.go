package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	c_utils "pchat/controller/utils"
	"pchat/model"
	pb_common "pchat/pb/common"
	pb_user "pchat/pb/user"
	"pchat/utils"
)

// @Description
// @Router		/users/validOTP [post]
// @Tags		账户
// @Summary	验证双因素密码
// @Accept		json
// @Produce	json
// @Param		body	body		pb_common.StringValue	true	"body"
// @Success	200		{object}	pb_user.LoginResponse
func validOTP(ctx *gin.Context, req *pb_common.StringValue) (*pb_user.LoginResponse, error) {
	user, err := model.CUser.GetById(ctx, utils.GetUserIdAsObjectId(ctx))
	if err != nil {
		return nil, err
	}
	ok, err := user.VerifyOTP(ctx, req.Value)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("invalid OTP")
	}
	token, err := model.SignToken(ctx, user, true)
	return formatLoginResponse(ctx, user, token, false), nil
}

var ValidOTP = c_utils.NewGinController[*pb_common.StringValue, *pb_user.LoginResponse](validOTP)
