package todo

import (
	"net/http"
	"pchat/controller/core"
)

var Group = core.NewGroup("/todos")

func init() {
	Group.Register(core.NewController("/records", http.MethodGet, listTodoRecords))
}
