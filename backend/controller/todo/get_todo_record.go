package todo

import (
	"context"
	model_todo "pchat/model/todo"
	pb_common "pchat/pb/common"
	pb_todo "pchat/pb/todo"
	"pchat/repository/bson"
)

// @Description
// @Router		/todos/records/:id [get]
// @Tags		待办管理
// @Summary	获取一条待办内容
// @Accept		json
// @Produce	json
// @Success	200		{object}	pb_todo.TodoRecordDetail
// @Param		body	body		pb_common.DetailRequest	true	"body"
func getTodoRecord(ctx context.Context, req *pb_common.DetailRequest) (*pb_todo.TodoRecordDetail, error) {
	record, err := model_todo.CTodoRecord.GetById(ctx, bson.NewObjectIdFromHex(req.Id))
	if err != nil {
		return nil, err
	}
	todo, err := model_todo.CTodo.GetById(ctx, record.TodoId)
	if err != nil {
		return nil, err
	}
	return formatTodoRecordDetail(record, todo), nil
}
