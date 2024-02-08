package todo

import (
	"context"
	"errors"
	model_todo "pchat/model/todo"
	pb_common "pchat/pb/common"
	pb_todo "pchat/pb/todo"
	"pchat/repository/bson"
	"pchat/utils"
)

// @Description
// @Router		/todos [post]
// @Tags		待办管理
// @Summary	创建、修改待办
// @Accept		json
// @Produce	json
// @Success	200		{object}	nil
// @Param		body	body		pb_todo.UpsertTodoRequest	true	"body"
func upsertTodo(ctx context.Context, req *pb_todo.UpsertTodoRequest) (*pb_common.EmptyResponse, error) {
	if !bson.IsObjectIdHex(req.Id) {
		todo := model_todo.Todo{}
		if err := utils.Copier().From(req).To(&todo); err != nil {
			return nil, err
		}
		if err := todo.Create(ctx); err != nil {
			return nil, err
		}
		return &pb_common.EmptyResponse{}, nil
	}
	todo, err := model_todo.CTodo.GetById(ctx, bson.NewObjectIdFromHex(req.Id))
	if err != nil {
		return nil, err
	}
	condition := bson.M{
		"todoId":      todo.Id,
		"hasBeenDone": false,
		"isDeleted":   false,
	}
	records, err := model_todo.CTodoRecord.ListAllByCondition(ctx, condition)
	if err != nil {
		return nil, err
	}
	if len(records) > 1 {
		return nil, errors.New("invalid todo")
	}
	if err := utils.Copier().From(req).To(&todo); err != nil {
		return nil, err
	}
	if err := todo.Update(ctx); err != nil {
		return nil, err
	}
	// TODO: update record
	return &pb_common.EmptyResponse{}, nil
}
