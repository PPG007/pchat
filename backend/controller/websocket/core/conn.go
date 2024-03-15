package core

import (
	"context"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	pb_chat "pchat/pb/chat"
	"pchat/utils"
	"pchat/utils/env"
	"pchat/utils/log"
	"reflect"
	"strings"
)

type MessageHandler func(conn *Connection, clientId string, payload interface{})

type TypedMessageHandler[T proto.Message] func(conn *Connection, clientId string, payload T)

func NewHandler[T proto.Message](handler TypedMessageHandler[T]) MessageHandler {
	method := reflect.ValueOf(handler)
	return func(conn *Connection, clientId string, payload interface{}) {
		payloadValue := reflect.ValueOf(payload)
		if !payloadValue.IsValid() {
			return
		}
		if payloadValue.Type().ConvertibleTo(method.Type().In(2)) {
			method.Call([]reflect.Value{
				reflect.ValueOf(conn),
				reflect.ValueOf(clientId),
				payloadValue.Convert(method.Type().In(2)),
			})
		}
	}
}

type Connection struct {
	serverId         string
	clientId         string
	conn             *websocket.Conn
	ctx              context.Context
	handlers         map[pb_chat.MessageType]MessageHandler
	heartbeatSender  *heartbeatSender
	heartbeatChecker *heartbeatChecker
	message          chan *pb_chat.Message
}

var emptyHandler MessageHandler = func(conn *Connection, clientId string, payload interface{}) {
	return
}

var heartbeatHandler = NewHandler(func(conn *Connection, clientId string, payload *pb_chat.HeartbeatMessage) {
	conn.heartbeatChecker.beat()
})

func NewConnection(ctx context.Context, conn *websocket.Conn) *Connection {
	return &Connection{
		serverId:         env.ServerId,
		clientId:         uuid.NewString(),
		conn:             conn,
		ctx:              ctx,
		handlers:         make(map[pb_chat.MessageType]MessageHandler),
		heartbeatChecker: newChecker(ctx),
		heartbeatSender:  newSender(ctx, conn),
		message:          make(chan *pb_chat.Message, 100),
	}
}

func (conn *Connection) Start() {
	conn.startReading()
	conn.heartbeatSender.start()
	conn.heartbeatChecker.start()
LOOP:
	for {
		select {
		case <-conn.heartbeatChecker.timeout:
			break LOOP
		case message, ok := <-conn.message:
			if !ok {
				break LOOP
			}
			msgType, msg := parseMessage(message)
			conn.getHandler(msgType)(conn, message.ClientId, msg)
		}
	}
	conn.Stop()
}

func (conn *Connection) SetHandler(messageType pb_chat.MessageType, handler MessageHandler) {
	conn.handlers[messageType] = handler
}

func (conn *Connection) getHandler(messageType pb_chat.MessageType) MessageHandler {
	switch messageType {
	case pb_chat.MessageType_HEARTBEAT:
		return heartbeatHandler
	}
	if handler, ok := conn.handlers[messageType]; ok {
		return handler
	}
	return emptyHandler
}

func (conn *Connection) startReading() {
	utils.GO(conn.ctx, func(ctx context.Context) {
		for {
			message := &pb_chat.Message{}
			err := conn.conn.ReadJSON(message)
			if err != nil {
				return
			}
			conn.message <- message
		}
	})
}

func (conn *Connection) SendMessage(data any) error {
	return conn.conn.WriteJSON(data)
}

func (conn *Connection) Stop() {
	log.Info(conn.ctx, "Stopping connection...", log.Fields{})
	conn.heartbeatSender.stop()
	conn.heartbeatChecker.stop()
	log.Info(conn.ctx, "Connection stopped", log.Fields{})
}

func parseMessage(message *pb_chat.Message) (pb_chat.MessageType, interface{}) {
	rv := reflect.ValueOf(message).Elem()
	var msg interface{}
	for i := 0; i < rv.Type().NumField(); i++ {
		if strings.ToUpper(rv.Type().Field(i).Name) == pb_chat.MessageType_name[int32(message.Type)] {
			if !rv.Field(i).IsNil() {
				msg = rv.Field(i).Interface()
			}
		}
	}
	return message.Type, msg
}
