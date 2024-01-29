package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	pb_common "pchat/pb/common"
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
