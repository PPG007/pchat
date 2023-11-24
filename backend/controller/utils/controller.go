package utils

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
	"net/http"
	pb_common "pchat/pb/common"
	"pchat/utils"
)

type Handler[Request, Response proto.Message] func(*gin.Context, Request) (Response, error)

type GinController = func(*gin.Context)

func NewGinController[Request, Response proto.Message](fn Handler[Request, Response]) GinController {
	return func(ctx *gin.Context) {
		req := new(Request)
		var err error
		if ctx.Request.Method == http.MethodGet {
			err = ctx.ShouldBindQuery(req)
		} else {
			err = ctx.ShouldBindJSON(req)
		}
		if err != nil {
			ResponseError(ctx, err)
			return
		}
		err = utils.ValidateRequest(req)
		if err != nil {
			ResponseError(ctx, err)
			return
		}
		resp, err := fn(ctx, *req)
		if err != nil {
			ResponseError(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func ResponseError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, pb_common.ErrorResponse{
		Message: err.Error(),
	})
}
