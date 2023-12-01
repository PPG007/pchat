package user

import (
	"github.com/gin-gonic/gin"
	c_utils "pchat/controller/utils"
	"pchat/model"
	pb_user "pchat/pb/user"
	"pchat/repository/bson"
	"pchat/utils"
)

func listRegisterApplication(ctx *gin.Context, req *pb_user.ListRegisterApplicationRequest) (*pb_user.ListRegisterApplicationResponse, error) {
	condition := bson.M{}
	if len(req.Status) > 0 {
		condition["status"] = bson.M{
			"$in": req.Status,
		}
	}
	applications, total, err := model.CRegisterApplication.ListByPagination(ctx, utils.FormatPagination(condition, req.ListCondition))
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

var ListRegisterApplicationController = c_utils.NewGinController[*pb_user.ListRegisterApplicationRequest, *pb_user.ListRegisterApplicationResponse](listRegisterApplication)
