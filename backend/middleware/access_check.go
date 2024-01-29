package middleware

import (
	"github.com/gin-gonic/gin"
	"os"
	"pchat/utils"
	"pchat/utils/log"
)

var (
	permissionMap = make(map[string]string)
)

func init() {
	os.Open("../assets/permission.yaml")
}

func accessCheck(ctx *gin.Context) {
	// TODO:
	log.Info(ctx, "access check", log.Fields{
		"userId": utils.GetUserId(ctx),
	})
	ctx.Next()
}
