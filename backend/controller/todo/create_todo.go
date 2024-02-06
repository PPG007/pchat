package todo

import (
	"context"
	model_todo "pchat/model/todo"
	pb_common "pchat/pb/common"
	pb_todo "pchat/pb/todo"
	"pchat/utils"
)

// @Description
// @Router		/todos [post]
// @Tags		待办管理
// @Summary	创建待办
// @Accept		json
// @Produce	json
// @Success	200		{object}	nil
// @Param		body	body		pb_todo.CreateTodoRequest	true	"body"
func createTodo(ctx context.Context, req *pb_todo.CreateTodoRequest) (*pb_common.EmptyResponse, error) {
	todo := model_todo.Todo{}
	err := utils.Copier().From(req).To(&todo)
	if err != nil {
		return nil, err
	}
	return &pb_common.EmptyResponse{}, nil
}
