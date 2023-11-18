package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"pchat/pb"
)

func demo(ctx *gin.Context, req *pb.DemoRequest) (*pb.DemoResponse, error) {
	if req.Value == "test" {
		return nil, errors.New("test")
	}
	return &pb.DemoResponse{
		Value: req.Value,
	}, nil
}

var demoController = newController[*pb.DemoRequest, *pb.DemoResponse](demo)
