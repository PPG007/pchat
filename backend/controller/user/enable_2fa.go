package user

import (
	"github.com/gin-gonic/gin"
	c_utils "pchat/controller/utils"
	"pchat/model"
	pb_common "pchat/pb/common"
	pb_user "pchat/pb/user"
	"pchat/utils"
)

// @Description
// @Router		/users/enable2FA [post]
// @Tags		账户
// @Summary	开启双因素认证
// @Accept		json
// @Produce	json
// @Success	200	{object}	pb_user.Enable2FAResponse
func enable2FA(ctx *gin.Context, req *pb_common.EmptyRequest) (*pb_user.Enable2FAResponse, error) {
	url, recoveryCodes, err := model.CUser.Enable2FA(ctx, utils.GetUserIdAsObjectId(ctx))
	if err != nil {
		return nil, err
	}
	return &pb_user.Enable2FAResponse{
		Url:           url,
		RecoveryCodes: recoveryCodes,
	}, nil
}

var Enable2FA = c_utils.NewGinController(enable2FA)
