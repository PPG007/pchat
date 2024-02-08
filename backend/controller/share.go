package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"pchat/controller/todo"
	"pchat/controller/user"
	"pchat/controller/websocket"
	"pchat/utils"
)

func AppendRoutes(root *gin.Engine, isDebug bool) {
	utils.MergeEngines(root,
		user.Group.Engine(),
		todo.Group.Engine(),
		websocket.Group.Engine(),
	)
	if isDebug {
		printRoutes(root)
	}
}

func printRoutes(root *gin.Engine) {
	for _, info := range root.Routes() {
		log.Printf("[%s] %s -> %s\n", info.Method, info.Path, info.Handler)
	}
}
