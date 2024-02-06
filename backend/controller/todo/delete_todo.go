package todo

import (
	"context"
	model_todo "pchat/model/todo"
	pb_common "pchat/pb/common"
	"pchat/repository/bson"
)

// @Description
// @Router		/todos/:id [delete]
// @Tags		待办管理
// @Summary	删除待办
// @Accept		json
// @Produce	json
// @Success	200		{object}	nil
// @Param		body	body		pb_common.DetailRequest	true	"body"
func deleteTodo(ctx context.Context, req *pb_common.DetailRequest) (*pb_common.EmptyResponse, error) {
	err := model_todo.CTodo.DeleteById(ctx, bson.NewObjectIdFromHex(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb_common.EmptyResponse{}, nil
}
