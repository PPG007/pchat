package controller

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"net/http"
	"pchat/controller/admin"
	"pchat/controller/core"
	"pchat/controller/todo"
	"pchat/controller/user"
	"pchat/controller/websocket"
	"pchat/utils"
	"pchat/utils/env"
)

var (
	ControllersMap = make(map[string]*core.Controller)
)

func AppendRoutes(root *gin.Engine) {
	groups := []*core.Group{
		user.Group,
		todo.Group,
		websocket.Group,
		admin.Group,
	}
	for _, group := range groups {
		utils.MergeEngines(root, group.Engine())
		utils.MergeMaps(ControllersMap, group.ControllerMap)
	}
	if env.IsDebug() {
		printRoutes(groups)
	}
}

func printRoutes(groups []*core.Group) {
	methodColor := func(method string) string {
		col := color.BgHiGreen
		switch method {
		case http.MethodPost:
			col = color.BgHiMagenta
		case http.MethodPut:
			col = color.BgHiCyan
		case http.MethodDelete:
			col = color.BgHiRed
		}
		return utils.GetColorString(method, col)
	}
	pathColor := utils.ColorStringFn(color.FgHiBlue)
	handlerColor := utils.ColorStringFn(color.FgBlue)
	for _, group := range groups {
		for p, c := range group.ControllerMap {
			fmt.Printf("[%s] %s -> %s\n", methodColor(c.Method), pathColor(p), handlerColor(c.HandlerName))
		}
	}
}
