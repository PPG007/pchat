package websocket

import (
	"net/http"
	"pchat/controller/core"
)

var Group = core.NewGroup("/ws")

func init() {
	Group.Register(&core.Controller{
		Path:    "",
		Method:  http.MethodGet,
		Handler: core.WrapWSHandler(basic),
		NoAuth:  true,
	})
}
