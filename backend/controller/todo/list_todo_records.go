package todo

import (
	"context"
	pb_todo "pchat/pb/todo"
)

func listTodoRecords(ctx context.Context, req *pb_todo.ListTodoRecordsRequest) (*pb_todo.ListTodoRecordsResponse, error) {
	return &pb_todo.ListTodoRecordsResponse{}, nil
}
