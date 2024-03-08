package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"runtime"
)

type Group struct {
	BasePath      string
	group         *gin.Engine
	ControllerMap map[string]*Controller
}

func NewGroup(basePath string) *Group {
	gin.SetMode(gin.ReleaseMode)
	return &Group{
		BasePath:      basePath,
		group:         gin.New(),
		ControllerMap: make(map[string]*Controller),
	}
}

func (g *Group) Register(controller *Controller) {
	handler := controller.Handler
	if controller.HandlerName == "" {
		controller.HandlerName = runtime.FuncForPC(reflect.ValueOf(controller.Handler).Pointer()).Name()
	}
	path := fmt.Sprintf("%s%s", g.BasePath, controller.Path)
	g.ControllerMap[path] = controller
	switch controller.Method {
	case http.MethodGet:
		g.group.GET(path, handler)
	case http.MethodPost:
		g.group.POST(path, handler)
	case http.MethodDelete:
		g.group.DELETE(path, handler)
	case http.MethodPut:
		g.group.PUT(path, handler)
	case http.MethodPatch:
		g.group.PATCH(path, handler)
	default:
		panic("unsupported method")
	}
}

func (g *Group) Engine() *gin.Engine {
	return g.group
}
