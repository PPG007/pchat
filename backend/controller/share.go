package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"pchat/controller/core"
	"pchat/controller/todo"
	"pchat/controller/user"
	"pchat/controller/websocket"
	"pchat/utils"
)

var (
	ControllersMap = make(map[string]*core.Controller)
)

func AppendRoutes(root *gin.Engine, isDebug bool) {
	groups := []*core.Group{
		user.Group,
		todo.Group,
		websocket.Group,
	}
	for _, group := range groups {
		utils.MergeEngines(root, group.Engine())
		utils.MergeMaps(ControllersMap, group.ControllerMap)
	}
	if isDebug {
		printRoutes(root)
	}
}

func printRoutes(root *gin.Engine) {
	for _, info := range root.Routes() {
		log.Printf("[%s] %s -> %s\n", info.Method, info.Path, info.Handler)
	}
}
