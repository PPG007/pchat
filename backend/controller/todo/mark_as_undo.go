package todo

import (
	"context"
	model_todo "pchat/model/todo"
	pb_common "pchat/pb/common"
	"pchat/repository/bson"
)

func markAsUndo(ctx context.Context, req *pb_common.DetailRequest) (*pb_common.EmptyResponse, error) {
	err := model_todo.CTodoRecord.MarkAsUndo(ctx, bson.NewObjectIdFromHex(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb_common.EmptyResponse{}, nil
}
