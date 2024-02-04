package todo

import (
	"github.com/gin-gonic/gin"
	pb_todo "pchat/pb/todo"
)

func listTodoRecords(ctx *gin.Context, req *pb_todo.ListTodoRecordsRequest) (*pb_todo.ListTodoRecordsResponse, error) {
	return &pb_todo.ListTodoRecordsResponse{}, nil
}
