package middleware

import "github.com/gin-gonic/gin"

func ResponseErrorMessage(msg string) map[string]string {
	return map[string]string{
		"message": msg,
	}
}

func RegisterMiddlewares(e *gin.Engine) {
	e.Use(recovery)
	e.Use(auth)
	e.Use(accessCheck)
}
