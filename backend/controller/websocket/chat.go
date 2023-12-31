package websocket

import (
	"github.com/PPG007/pubsub"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	c_utils "pchat/controller/utils"
	"pchat/model"
	pb_chat "pchat/pb/chat"
	"pchat/utils"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var bus = pubsub.New()

const (
	MESSAGE_TOPIC = "message"
)

func ChatHandler(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		c_utils.ResponseError(ctx, err)
		return
	}
	defer conn.Close()
	user, err := model.CUser.Online(ctx)
	if err != nil {
		c_utils.ResponseError(ctx, err)
		return
	}
	id := bus.Subscribe(MESSAGE_TOPIC, func(args ...interface{}) {
		message := args[0].(model.Message)
		if message.Sender.Hex() == utils.GetUserId(ctx) {
			return
		}
		detail := &pb_chat.MessageDetail{}
		utils.Copier().From(message).To(detail)
		conn.WriteJSON(detail)
	})
	defer bus.Unsubscribe(MESSAGE_TOPIC, id)
	defer user.Offline(ctx)
	readMessages(ctx, conn)
}

func readMessages(ctx *gin.Context, conn *websocket.Conn) {
	for {
		message := &pb_chat.NewMessage{}
		err := conn.ReadJSON(message)
		if err != nil {
			if websocket.IsCloseError(err) || websocket.IsUnexpectedCloseError(err) {
				return
			}
			c_utils.ResponseError(ctx, err)
			return
		}
		m, err := model.CMessage.CreateFromPb(ctx, message)
		if err != nil {
			c_utils.ResponseError(ctx, err)
			return
		}
		bus.Publish(MESSAGE_TOPIC, m)
	}
}
