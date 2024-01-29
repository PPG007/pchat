package websocket

import (
	"net/http"
	"pchat/controller/core"
)

var Group = core.NewGroup("/ws")

func init() {
	Group.Register(&core.Controller{
		Path:    "/chat",
		Method:  http.MethodGet,
		Handler: ChatHandler,
	})
}
