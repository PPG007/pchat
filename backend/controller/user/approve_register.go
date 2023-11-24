package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"pchat/controller/utils"
	"pchat/model"
	pb_common "pchat/pb/common"
	pb_user "pchat/pb/user"
)

func approveRegister(ctx *gin.Context, req *pb_user.ApproveRegisterRequest) (*pb_common.EmptyResponse, error) {
	user, err := model.CUser.GetByEmail(ctx, req.Email, false)
	if err != nil {
		return nil, err
	}
	if user.Status != model.USER_STATUS_AUDITING {
		return nil, errors.New("user don't need approve")
	}
	err = user.Activate(ctx)
	if err != nil {
		return nil, err
	}
	return &pb_common.EmptyResponse{}, nil
}

var ApproveRegisterController = utils.NewGinController[*pb_user.ApproveRegisterRequest, *pb_common.EmptyResponse](approveRegister)
