package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"pchat/controller/user"
	"pchat/controller/websocket"
	"pchat/utils"
)

func GetRoot(isDebug bool) *gin.Engine {
	root := gin.New()
	utils.MergeEngines(root,
		user.Group.Engine(),
		websocket.Group.Engine(),
	)
	if isDebug {
		printRoutes(root)
	}
	return root
}

func printRoutes(root *gin.Engine) {
	for _, info := range root.Routes() {
		log.Printf("[%s] %s -> %s\n", info.Method, info.Path, info.Handler)
	}
}
