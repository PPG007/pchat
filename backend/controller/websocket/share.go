package websocket

import "github.com/gin-gonic/gin"

func RegisterWS(engin *gin.Engine) {
	group := engin.Group("/ws")
	group.GET("/chat", ChatHandler)
}
