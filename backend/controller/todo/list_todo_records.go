package todo

import (
	"context"
	"pchat/model"
	model_todo "pchat/model/todo"
	pb_todo "pchat/pb/todo"
	"pchat/repository/bson"
	"pchat/utils"
	"time"
)

// @Description
// @Router		/todos/records [get]
// @Tags		待办管理
// @Summary	获取待办列表
// @Accept		json
// @Produce	json
// @Success	200		{object}	pb_todo.ListTodoRecordsResponse
// @Param		body	body		pb_todo.ListTodoRecordsRequest	true	"body"
func listTodoRecords(ctx context.Context, req *pb_todo.ListTodoRecordsRequest) (*pb_todo.ListTodoRecordsResponse, error) {
	condition := genListTodoRecordsCondition(ctx, req)
	total, records, err := model_todo.CTodoRecord.ListByPagination(ctx, utils.FormatPagination(condition, req.ListCondition))
	if err != nil {
		return nil, err
	}
	details, err := formatTodoRecordDetails(ctx, records)
	if err != nil {
		return nil, err
	}
	return &pb_todo.ListTodoRecordsResponse{
		Total: total,
		Items: details,
	}, nil
}

func genListTodoRecordsCondition(ctx context.Context, req *pb_todo.ListTodoRecordsRequest) bson.M {
	condition := model.GenDefaultUserIdCondition(ctx)
	if req.HasBeenDone != nil {
		condition["hasBeenDone"] = req.HasBeenDone.Value
	}
	if req.SearchKey != "" {
		condition["content"] = utils.GetFuzzySearchStrRegex(req.SearchKey)
	}
	return condition
}

func formatTodoRecordDetails(ctx context.Context, records []model_todo.TodoRecord) ([]*pb_todo.TodoRecordDetail, error) {
	todoIds := make([]bson.ObjectId, 0, len(records))
	for _, record := range records {
		todoIds = append(todoIds, record.TodoId)
	}
	todos, err := model_todo.CTodo.ListByIds(ctx, todoIds)
	if err != nil {
		return nil, err
	}
	todoMap := make(map[bson.ObjectId]model_todo.Todo, len(todos))
	for _, todo := range todos {
		todoMap[todo.Id] = todo
	}
	result := make([]*pb_todo.TodoRecordDetail, 0, len(records))
	for _, record := range records {
		result = append(result, formatTodoRecordDetail(record, todoMap[record.TodoId]))
	}
	return result, nil
}

func formatTodoRecordDetail(record model_todo.TodoRecord, todo model_todo.Todo) *pb_todo.TodoRecordDetail {
	detail := &pb_todo.TodoRecordDetail{}
	utils.Copier().From(record).To(detail)
	detail.RemindSetting = &pb_todo.RemindSetting{
		RemindAt:         record.RemindAt.Format(time.RFC3339),
		IsRepeatable:     todo.RemindSetting.IsRepeatable,
		RepeatType:       todo.RemindSetting.RepeatType,
		RepeatDateOffset: todo.RemindSetting.RepeatDateOffset,
	}
	return detail
}
