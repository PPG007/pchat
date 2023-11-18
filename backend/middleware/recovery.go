package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pchat/utils/log"
	"runtime"
)

func recovery(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			stack := make([]byte, 4096)
			stack = stack[:runtime.Stack(stack, false)]
			err := fmt.Sprintf("%v", r)
			log.ErrorTrace(ctx, "Uncaught panic", log.Fields{
				"error": err,
			}, stack)
			ctx.JSON(http.StatusInternalServerError, ResponseErrorMessage(err))
		}
	}()
	ctx.Next()
}
