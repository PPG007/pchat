package middleware

import "github.com/gin-gonic/gin"

var (
	permissionMap = make(map[string]string)
)

func accessCheck(ctx *gin.Context) {
	ctx.Next()
}
