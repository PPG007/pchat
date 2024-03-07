package admin

import (
	"net/http"
	"pchat/controller/core"
	"pchat/permissions"
)

var (
	Group = core.NewGroup("/admin")
)

func init() {
	Group.Register(core.NewController("/setting", http.MethodPut, updateSetting, core.WithPermission(permissions.SETTING_EDIT)))
}
