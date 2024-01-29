package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	RESPONSE_BODY_KEY = "responseBody"
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
		ctx.Set(RESPONSE_BODY_KEY, writer.Body.String())
	}
}
