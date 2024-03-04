package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pchat/controller"
	model_user "pchat/model/user"
	pb_common "pchat/pb/common"
	"pchat/utils"
)

func noAuth(path string) bool {
	if c := controller.ControllersMap[path]; c != nil {
		return c.NoAuth
	}
	return false
}

func allowUnauthorizedToken(path string) bool {
	if c := controller.ControllersMap[path]; c != nil {
		return c.AllowUnauthorizedToken
	}
	return false
}

func auth(ctx *gin.Context) {
	fullPath := ctx.FullPath()
	if noAuth(fullPath) {
		ctx.Next()
		return
	}
	token := utils.GetToken(ctx)
	userClaim, err := model_user.ValidToken(ctx, token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, pb_common.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	utils.SetUserId(ctx, userClaim.UserId)
	if !userClaim.IsAuthorized {
		if allowUnauthorizedToken(fullPath) {
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, pb_common.EmptyResponse{})
		}
		return
	}
	ctx.Next()
}
