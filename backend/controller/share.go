package controller

import (
	"github.com/gin-gonic/gin"
	"pchat/controller/user"
)

func RegisterControllers(engin *gin.Engine) {
	user.RegisterRoutes(engin)
}
