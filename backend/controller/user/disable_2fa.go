package user

import (
	"github.com/gin-gonic/gin"
	c_utils "pchat/controller/utils"
	"pchat/model"
	pb_common "pchat/pb/common"
	"pchat/utils"
)

func disable2FA(ctx *gin.Context, req *pb_common.EmptyRequest) (*pb_common.EmptyResponse, error) {
	err := model.CUser.DisableOTP(ctx, utils.GetUserIdAsObjectId(ctx))
	if err != nil {
		return nil, err
	}
	return &pb_common.EmptyResponse{}, nil
}

var Disable2FA = c_utils.NewGinController(disable2FA)
