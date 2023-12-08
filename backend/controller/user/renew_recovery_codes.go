package user

import (
	"github.com/gin-gonic/gin"
	c_utils "pchat/controller/utils"
	"pchat/model"
	pb_common "pchat/pb/common"
	"pchat/utils"
)

func renewRecoveryCodes(ctx *gin.Context, req *pb_common.EmptyRequest) (*pb_common.StringArrayValue, error) {
	codes, err := model.CUser.GenerateRecoveryCodes(ctx, utils.GetUserIdAsObjectId(ctx), true)
	if err != nil {
		return nil, err
	}
	return &pb_common.StringArrayValue{
		Values: codes,
	}, nil
}

var RenewRecoveryCodes = c_utils.NewGinController(renewRecoveryCodes)
