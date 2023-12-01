package user

import (
	"github.com/gin-gonic/gin"
	c_utils "pchat/controller/utils"
	"pchat/model"
	pb_common "pchat/pb/common"
	pb_user "pchat/pb/user"
	"pchat/utils"
)

func approveRegister(ctx *gin.Context, req *pb_user.AuditRegisterApplicationRequest) (*pb_common.EmptyResponse, error) {
	if !req.IsApproved {
		err := model.CRegisterApplication.Reject(ctx, utils.StrArrToObjectIds(req.Ids), req.RejectReason)
		if err != nil {
			return nil, err
		}
		return &pb_common.EmptyResponse{}, nil
	}
	err := model.CRegisterApplication.Approve(ctx, utils.StrArrToObjectIds(req.Ids))
	if err != nil {
		return nil, err
	}
	return &pb_common.EmptyResponse{}, nil
}

var ApproveRegisterController = c_utils.NewGinController[*pb_user.AuditRegisterApplicationRequest, *pb_common.EmptyResponse](approveRegister)
