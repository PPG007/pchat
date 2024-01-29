package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"pchat/utils"
)

func rewrite(ctx *gin.Context) {
	utils.SetRequestId(ctx, uuid.NewString())
	ctx.Next()
}
