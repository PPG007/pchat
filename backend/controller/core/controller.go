package core

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/joncalhoun/qson"
	jsoniter "github.com/json-iterator/go"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
	"pchat/utils"
	"reflect"
	"runtime"
)

type Handler[Request, Response proto.Message] func(context.Context, Request) (Response, error)

type GinController = func(*gin.Context)

type Controller struct {
	Path                   string
	Method                 string
	Permission             string
	NoAuth                 bool
	AllowUnauthorizedToken bool
	Handler                func(ctx *gin.Context)
	HandlerName            string
}

type ControllerOption func(c *Controller)

func WithPermission(permission string) ControllerOption {
	return func(c *Controller) {
		c.Permission = permission
	}
}

func WithNoAuth() ControllerOption {
	return func(c *Controller) {
		c.NoAuth = true
	}
}

func WithAllowUnauthorizedToken() ControllerOption {
	return func(c *Controller) {
		c.AllowUnauthorizedToken = true
	}
}

func NewController[Request, Response proto.Message](path, method string, handler Handler[Request, Response], opts ...ControllerOption) *Controller {
	c := &Controller{
		Path:        path,
		Method:      method,
		Handler:     wrapController(handler),
		HandlerName: runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name(),
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func newReq[T proto.Message]() T {
	return proto.Clone(*(new(T))).ProtoReflect().New().Interface().(T)
}

func getRequestData[T proto.Message](ctx *gin.Context) (*T, error) {
	var (
		req = newReq[T]()
	)

	if query := ctx.Request.URL.RawQuery; query != "" {
		data, err := qson.ToJSON(query)
		if err != nil {
			return nil, err
		}
		if err := jsoniter.Unmarshal(data, req); err != nil {
			return nil, err
		}
	}
	if ctx.Request.ContentLength > 0 {
		if err := ctx.ShouldBindJSON(req); err != nil && !errors.Is(err, io.EOF) {
			return nil, err
		}
	}
	if params := ctx.Params; len(params) > 0 {
		data := make(map[string]string, len(params))
		for _, param := range params {
			data[param.Key] = param.Value
		}
		if err := utils.SetStructFields(req, data); err != nil {
			return nil, err
		}
	}
	return &req, nil
}

func wrapController[Request, Response proto.Message](handler Handler[Request, Response]) GinController {
	return func(ctx *gin.Context) {
		req, err := getRequestData[Request](ctx)
		if err != nil {
			utils.ResponseError(ctx, err)
			return
		}
		err = utils.ValidateRequest(*req)
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
