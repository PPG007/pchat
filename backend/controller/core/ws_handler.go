package core

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"pchat/utils"
)

type WSHandler = func(ctx context.Context, conn *websocket.Conn)

var (
	upgrader = websocket.Upgrader{}
)

func WrapWSHandler(handler WSHandler) GinController {
	return func(ctx *gin.Context) {
		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			utils.ResponseError(ctx, err)
			return
		}
		defer conn.Close()
		handler(utils.CopyContext(ctx), conn)
	}
}
