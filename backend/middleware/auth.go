package middleware

import "github.com/gin-gonic/gin"

func auth(ctx *gin.Context) {
	ctx.Next()
}
