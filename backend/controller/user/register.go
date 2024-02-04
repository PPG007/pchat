package user

import (
	"context"
	"errors"
	"github.com/qiniu/qmgo"
	"pchat/model/common"
	model_user "pchat/model/user"
	pb_user "pchat/pb/user"
)

// @Description
// @Router		/users/register [post]
// @Tags		账户
// @Summary	注册
// @Accept		json
// @Produce	json
// @Param		body	body		pb_user.RegisterRequest	true	"body"
// @Success	200		{object}	pb_user.RegisterResponse
func register(ctx context.Context, req *pb_user.RegisterRequest) (*pb_user.RegisterResponse, error) {
	_, err := model_user.CUser.GetByEmail(ctx, req.Email, false)
	if err == nil {
		return nil, errors.New("user already exists")
	}
	if !errors.Is(err, qmgo.ErrNoSuchDocuments) {
		return nil, err
	}
	setting, err := common.CSetting.Get(ctx)
	if err != nil {
		return nil, err
	}
	err = model_user.CUser.CreateNew(ctx, req.Email, req.Password, req.Reason, setting.ChatSetting.MustBeApprovedBeforeRegister)
	if err != nil {
		return nil, err
	}

	return &pb_user.RegisterResponse{
		NeedAudit: setting.ChatSetting.MustBeApprovedBeforeRegister,
	}, nil
}
