package admin

import (
	"net/http"
	"pchat/controller/core"
)

var (
	Group = core.NewGroup("/admin")
)

func init() {
	Group.Register(core.NewController("/setting", http.MethodPut, updateSetting))
}
