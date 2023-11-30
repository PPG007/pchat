package user

import (
	"github.com/PPG007/copier"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	c_utils "pchat/controller/utils"
	"pchat/model"
	pb_user "pchat/pb/user"
	"pchat/repository/bson"
	"pchat/utils"
)

func login(ctx *gin.Context, req *pb_user.LoginRequest) (*pb_user.LoginResponse, error) {
	user, err := model.CUser.GetByEmail(ctx, req.Email, true)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	token, err := model.SignToken(ctx, user)
	if err != nil {
		return nil, err
	}
	resp := &pb_user.LoginResponse{}
	resp.Token = token
	utils.Copier().RegisterDiffPairs([]copier.DiffPair{
		{
			Origin: "Roles",
			Target: []string{"Permissions"},
		},
	}).RegisterTransformer("Permissions", func(roles []bson.ObjectId) []string {
		permissions, err := model.CRole.GetPermissionsByIds(ctx, roles)
		if err != nil {
			return nil
		}
		return permissions
	}).From(user).To(resp)
	return resp, nil
}

var LoginController = c_utils.NewGinController[*pb_user.LoginRequest, *pb_user.LoginResponse](login)
