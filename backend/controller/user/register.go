package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/qmgo"
	"pchat/controller/utils"
	"pchat/model"
	pb_user "pchat/pb/user"
)

func register(ctx *gin.Context, req *pb_user.RegisterRequest) (*pb_user.RegisterResponse, error) {
	_, err := model.CUser.GetByEmail(ctx, req.Email, false)
	if err == nil {
		return nil, errors.New("user already exists")
	}
	if !errors.Is(err, qmgo.ErrNoSuchDocuments) {
		return nil, err
	}
	setting, err := model.CSetting.Get(ctx)
	if err != nil {
		return nil, err
	}
	err = model.CUser.CreateNew(ctx, req.Email, req.Password, setting.ChatSetting.MustBeApprovedBeforeRegister)
	if err != nil {
		return nil, err
	}

	return &pb_user.RegisterResponse{
		NeedAudit: setting.ChatSetting.MustBeApprovedBeforeRegister,
	}, nil
}

var RegisterController = utils.NewGinController[*pb_user.RegisterRequest, *pb_user.RegisterResponse](register)
