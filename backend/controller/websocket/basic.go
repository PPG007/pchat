package websocket

import (
	"context"
	"github.com/gorilla/websocket"
	"pchat/controller/websocket/core"
)

func basic(ctx context.Context, wsConn *websocket.Conn) {
	conn := core.NewConnection(ctx, wsConn)
	conn.Start()
}
