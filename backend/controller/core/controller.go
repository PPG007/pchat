package core

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joncalhoun/qson"
	"google.golang.org/protobuf/proto"
	"net/http"
	pb_common "pchat/pb/common"
	"pchat/utils"
)

type Handler[Request, Response proto.Message] func(context.Context, Request) (Response, error)

type GinController = func(*gin.Context)

type Controller struct {
	Path    string
	Method  string
	Handler func(ctx *gin.Context)
}

func NewController[Request, Response proto.Message](path, method string, handler Handler[Request, Response]) *Controller {
	return &Controller{
		Path:    path,
		Method:  method,
		Handler: wrapController(handler),
	}
}

func getRequestData[T proto.Message](ctx *gin.Context) (*T, error) {
	var (
		err error
		req = new(T)
	)
	if ctx.Request.Method == http.MethodGet {
		err = qson.Unmarshal(req, ctx.Request.URL.RawQuery)
	} else {
		err = ctx.ShouldBindJSON(req)
	}
	if err != nil {
		return nil, err
	}
	return req, nil
}

func wrapController[Request, Response proto.Message](handler Handler[Request, Response]) GinController {
	return func(ctx *gin.Context) {
		req, err := getRequestData[Request](ctx)
		if err != nil {
			utils.ResponseError(ctx, err)
			return
		}
		err = utils.ValidateRequest(req)
		if err != nil {
			utils.ResponseError(ctx, err)
			return
		}
		resp, err := handler(ctx, *req)
		if err != nil {
			utils.ResponseError(ctx, err)
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
