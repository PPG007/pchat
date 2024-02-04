package user

import (
	"context"
	"errors"
	model_user "pchat/model/user"
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
func validOTP(ctx context.Context, req *pb_common.StringValue) (*pb_user.LoginResponse, error) {
	user, err := model_user.CUser.GetById(ctx, utils.GetUserIdAsObjectId(ctx))
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
	token, err := model_user.SignToken(ctx, user, true)
	return formatLoginResponse(ctx, user, token, false), nil
}
