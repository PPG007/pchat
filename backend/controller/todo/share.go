package todo

import (
	"net/http"
	"pchat/controller/core"
)

var Group = core.NewGroup("/todos")

func init() {
	Group.Register(core.NewController("/records", http.MethodGet, listTodoRecords))
	Group.Register(core.NewController("/records/:id", http.MethodGet, getTodoRecord))
	Group.Register(core.NewController("/upsert", http.MethodPost, upsertTodo))
	Group.Register(core.NewController("/:id", http.MethodDelete, deleteTodo))
	Group.Register(core.NewController("/:id/done", http.MethodPost, markAsDone))
	Group.Register(core.NewController("/:id/undo", http.MethodPost, markAsUndo))
}
