package common

import (
	"net/http"
	"pchat/controller/core"
)

var Group = core.NewGroup("/common")

func init() {
	Group.Register(core.NewController("/putObjectURL", http.MethodGet, getPutObjectURL))
}
