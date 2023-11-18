package controller

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
	"net/http"
)

type handler[Request, Response proto.Message] func(*gin.Context, Request) (Response, error)

type ginController = func(*gin.Context)

func newController[Request, Response proto.Message](fn handler[Request, Response]) ginController {
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
		resp, err := fn(ctx, *req)
		if err != nil {
			ResponseError(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func ResponseError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, map[string]string{
		"message": err.Error(),
	})
}

func RegisterControllers(engin *gin.Engine) {
	engin.GET("/demo", demoController)
	engin.POST("/demo", demoController)
}
