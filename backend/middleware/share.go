package middleware

import "github.com/gin-gonic/gin"

func RegisterMiddlewares(e *gin.Engine) {
	e.Use(recovery)
	e.Use(responseWriterMiddleware)
	e.Use(accessLog)
	e.Use(rewrite)
	e.Use(auth)
	e.Use(accessCheck)
}
