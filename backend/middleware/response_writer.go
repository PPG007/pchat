package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
	"pchat/utils"
)

type responseWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func (w responseWriter) Write(data []byte) (int, error) {
	w.Body.Write(data)
	return w.ResponseWriter.Write(data)
}

func responseWriterMiddleware(ctx *gin.Context) {
	writer := responseWriter{
		ResponseWriter: ctx.Writer,
		Body:           &bytes.Buffer{},
	}
	ctx.Writer = writer
	ctx.Next()
	if ctx.Writer.Status() >= http.StatusBadRequest {
		utils.SetResponseBody(ctx, writer.Body.String())
	}
}
