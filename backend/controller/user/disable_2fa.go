package user

import (
	"github.com/gin-gonic/gin"
	"pchat/model"
	pb_common "pchat/pb/common"
	"pchat/utils"
)

// @Description
// @Router		/users/disable2FA [post]
// @Tags		账户
// @Summary	禁用双因素认证
// @Accept		json
// @Produce	json
// @Success	200	{object}	nil
func disable2FA(ctx *gin.Context, req *pb_common.EmptyRequest) (*pb_common.EmptyResponse, error) {
	err := model.CUser.DisableOTP(ctx, utils.GetUserIdAsObjectId(ctx))
	if err != nil {
		return nil, err
	}
	return &pb_common.EmptyResponse{}, nil
}
