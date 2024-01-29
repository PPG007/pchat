package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	REQUEST_ID_HEADER = "X-Request-Id"
)

func rewrite(ctx *gin.Context) {
	reqId := uuid.NewString()
	ctx.Header(REQUEST_ID_HEADER, reqId)
	ctx.Request.Header.Set(REQUEST_ID_HEADER, reqId)
	ctx.Next()
}
