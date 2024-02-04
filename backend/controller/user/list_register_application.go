package user

import (
	"context"
	model_user "pchat/model/user"
	pb_user "pchat/pb/user"
	"pchat/repository/bson"
	"pchat/utils"
)

// @Description
// @Router		/users/registerApplications [get]
// @Tags		账户
// @Summary	获取注册申请列表
// @Accept		json
// @Produce	json
// @Success	200		{object}	pb_user.ListRegisterApplicationResponse
// @Param		body	body		pb_user.ListRegisterApplicationRequest	true	"body"
func listRegisterApplications(ctx context.Context, req *pb_user.ListRegisterApplicationRequest) (*pb_user.ListRegisterApplicationResponse, error) {
	condition := bson.M{}
	if len(req.Status) > 0 {
		condition["status"] = bson.M{
			"$in": req.Status,
		}
	}
	applications, total, err := model_user.CRegisterApplication.ListByPagination(ctx, utils.FormatPagination(condition, req.ListCondition))
	if err != nil {
		return nil, err
	}
	var details []*pb_user.RegisterApplicationDetail
	utils.Copier().From(applications).To(&details)
	return &pb_user.ListRegisterApplicationResponse{
		Total: total,
		Items: details,
	}, nil
}
