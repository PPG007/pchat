package user

import (
	"context"
	"github.com/PPG007/copier"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	c_utils "pchat/controller/utils"
	"pchat/model"
	pb_user "pchat/pb/user"
	"pchat/repository/bson"
	"pchat/utils"
)

//	@BasePath	/users

//	@Description
//	@Router		/login [post]
//	@Tags		账户
//	@Summary	登录
//	@Accept		json
//	@Produce	json
//	@Success	200		{object}	pb_user.LoginResponse
//	@Param		body	body		pb_user.LoginRequest	true	"body"
func login(ctx *gin.Context, req *pb_user.LoginRequest) (*pb_user.LoginResponse, error) {
	user, err := model.CUser.GetByEmail(ctx, req.Email, true)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	token, err := model.SignToken(ctx, user, !user.Is2FAEnabled)
	if err != nil {
		return nil, err
	}
	return formatLoginResponse(ctx, user, token, user.Is2FAEnabled), nil
}

func formatLoginResponse(ctx context.Context, user model.User, token string, need2FA bool) *pb_user.LoginResponse {
	resp := &pb_user.LoginResponse{
		Token:   token,
		Need2FA: need2FA,
	}
	if need2FA {
		return resp
	}
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
	return resp
}

var LoginController = c_utils.NewGinController[*pb_user.LoginRequest, *pb_user.LoginResponse](login)
