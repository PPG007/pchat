package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"pchat/controller"
	"pchat/errors"
	model_user "pchat/model/user"
	"pchat/utils"
)

var (
	permissionMap = make(map[string]string)
)

func init() {
	os.Open("../assets/permission.yaml")
}

func accessCheck(ctx *gin.Context) {
	c := controller.ControllersMap[ctx.FullPath()]
	if c == nil {
		ctx.Abort()
		return
	}
	// TODO: use cache
	if c.Permission != "" {
		permissions, err := model_user.CUser.GetPermissions(ctx, utils.GetUserIdAsObjectId(ctx))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.New(errors.ERR_COMMON_UNKNOWN, err.Error()))
			return
		}
		if !utils.StrInArray(c.Permission, &permissions) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, errors.New(errors.ERR_COMMON_PERMISSION_DENIED, "Permission Denied"))
			return
		}
	}
	ctx.Next()
}
