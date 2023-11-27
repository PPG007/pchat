package controller

import (
	"github.com/gin-gonic/gin"
	"pchat/controller/user"
	"pchat/controller/websocket"
)

func RegisterControllers(engin *gin.Engine) {
	user.RegisterRoutes(engin)
	websocket.RegisterWS(engin)
}
