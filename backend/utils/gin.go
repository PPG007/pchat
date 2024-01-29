package utils

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	pb_common "pchat/pb/common"
	"pchat/repository/bson"
)

func ResponseError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, pb_common.ErrorResponse{
		Message: err.Error(),
	})
}

func MergeEngine(main, sub *gin.Engine) {
	for _, info := range sub.Routes() {
		main.Handle(info.Method, info.Path, info.HandlerFunc)
	}
}

func MergeEngines(main *gin.Engine, subEngines ...*gin.Engine) {
	for _, engine := range subEngines {
		MergeEngine(main, engine)
	}
}

const (
	USER_ID_KEY       = "userId"
	REQUEST_ID_KEY    = "reqId"
	RESPONSE_BODY_KEY = "responseBody"
	ACCESS_TOKEN_KEY  = "X-Access-Token"
)

func GetToken(ctx *gin.Context) string {
	return ctx.GetHeader(ACCESS_TOKEN_KEY)
}

func SetUserId(ctx *gin.Context, userId string) {
	ctx.Set(USER_ID_KEY, userId)
}

func GetUserId(ctx context.Context) string {
	return ExtractValueFromContext(ctx, USER_ID_KEY)
}

func GetUserIdAsObjectId(ctx context.Context) bson.ObjectId {
	return bson.NewObjectIdFromHex(ExtractValueFromContext(ctx, USER_ID_KEY))
}

func SetResponseBody(ctx *gin.Context, body string) {
	ctx.Set(RESPONSE_BODY_KEY, body)
}

func GetResponseBody(ctx *gin.Context) string {
	return ctx.GetString(RESPONSE_BODY_KEY)
}

func SetRequestId(ctx *gin.Context, id string) {
	ctx.Set(REQUEST_ID_KEY, id)
	ctx.Header(REQUEST_ID_KEY, id)
}

func GetRequestId(ctx context.Context) string {
	return ExtractValueFromContext(ctx, REQUEST_ID_KEY)
}

func ExtractValueFromContext(ctx context.Context, key string) string {
	if ginCtx, ok := ctx.(*gin.Context); ok {
		return ginCtx.GetString(key)
	}
	return cast.ToString(ctx.Value(key))
}
