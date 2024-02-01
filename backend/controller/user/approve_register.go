package user

import (
	"github.com/gin-gonic/gin"
	model_user "pchat/model/user"
	pb_common "pchat/pb/common"
	pb_user "pchat/pb/user"
	"pchat/utils"
)

// @Description
// @Router		/users/approve [post]
// @Tags		账户
// @Summary	处理注册申请
// @Accept		json
// @Produce	json
// @Success	200		{object}	nil
// @Param		body	body		pb_user.AuditRegisterApplicationRequest	true	"body"
func approveRegister(ctx *gin.Context, req *pb_user.AuditRegisterApplicationRequest) (*pb_common.EmptyResponse, error) {
	if !req.IsApproved {
		err := model_user.CRegisterApplication.Reject(ctx, utils.StrArrToObjectIds(req.Ids), req.RejectReason)
		if err != nil {
			return nil, err
		}
		return &pb_common.EmptyResponse{}, nil
	}
	err := model_user.CRegisterApplication.Approve(ctx, utils.StrArrToObjectIds(req.Ids))
	if err != nil {
		return nil, err
	}
	return &pb_common.EmptyResponse{}, nil
}
