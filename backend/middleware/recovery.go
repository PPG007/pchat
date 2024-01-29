package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	pb_common "pchat/pb/common"
	"pchat/utils/log"
	"runtime"
)

func recovery(ctx *gin.Context) {
	accessLog := initAccessLog(ctx)
	defer func() {
		if r := recover(); r != nil {
			stack := make([]byte, 4096)
			stack = stack[:runtime.Stack(stack, false)]
			err := fmt.Sprintf("%v", r)
			accessLog.Record(ctx)
			log.ErrorTrace(ctx, "Uncaught panic", log.Fields{
				"error": err,
			}, stack)
			ctx.JSON(http.StatusInternalServerError, pb_common.ErrorResponse{
				Message: err,
			})
		}
	}()
	ctx.Next()
}
