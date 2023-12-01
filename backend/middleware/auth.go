package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pchat/model"
	pb_common "pchat/pb/common"
	"pchat/utils"
)

var (
	noAuthPaths = []string{
		"/users/login",
		"/users/register",
		"/users/registerApplications",
	}
)

const (
	TOKEN_HEADER   = "X-Access-Token"
	USER_ID_HEADER = "X-User-Id"
)

func auth(ctx *gin.Context) {
	if utils.StrInArray(ctx.FullPath(), &noAuthPaths) {
		ctx.Next()
		return
	}
	token := ctx.GetHeader(TOKEN_HEADER)
	userClaim, err := model.ValidToken(ctx, token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, pb_common.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	ctx.Request.Header.Set(USER_ID_HEADER, userClaim.UserId)
	ctx.Next()
}
