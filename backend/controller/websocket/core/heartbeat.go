package core

import (
	"context"
	"github.com/gorilla/websocket"
	pb_chat "pchat/pb/chat"
	"pchat/utils"
	"pchat/utils/log"
	"time"
)

const (
	heartbeatDuration = time.Second * 5
	heartbeatTimeout  = time.Second * 10
)

type heartbeatSender struct {
	ctx    context.Context
	ticker *time.Ticker
	conn   *websocket.Conn
}

func newSender(ctx context.Context, conn *websocket.Conn) *heartbeatSender {
	return &heartbeatSender{
		ctx:    ctx,
		ticker: time.NewTicker(heartbeatDuration),
		conn:   conn,
	}
}

func (sender *heartbeatSender) start() {
	utils.GO(sender.ctx, func(ctx context.Context) {
		for {
			select {
			case _, ok := <-sender.ticker.C:
				if !ok {
					return
				}
				message := &pb_chat.Message{
					Type: pb_chat.MessageType_HEARTBEAT,
					Heartbeat: &pb_chat.HeartbeatMessage{
						Time: time.Now().Format(time.RFC3339),
					},
				}
				sender.conn.WriteJSON(message)
			}
		}
	})
}

func (sender *heartbeatSender) stop() {
	log.Info(sender.ctx, "Stopping heartbeat sender", log.Fields{})
	sender.ticker.Stop()
	log.Info(sender.ctx, "Heartbeat sender stopped", log.Fields{})
}

type heartbeatChecker struct {
	ctx     context.Context
	ticker  *time.Ticker
	in      chan struct{}
	timeout chan struct{}
}

func newChecker(ctx context.Context) *heartbeatChecker {
	return &heartbeatChecker{
		ctx:     ctx,
		ticker:  time.NewTicker(heartbeatTimeout),
		in:      make(chan struct{}),
		timeout: make(chan struct{}, 1),
	}
}

func (checker *heartbeatChecker) start() {
	utils.GO(checker.ctx, func(ctx context.Context) {
		for {
			select {
			case <-checker.in:
				checker.ticker.Reset(heartbeatTimeout)
			case _, ok := <-checker.ticker.C:
				if !ok {
					return
				}
				checker.timeout <- struct{}{}
			}
		}
	})
}

func (checker *heartbeatChecker) beat() {
	checker.in <- struct{}{}
}

func (checker *heartbeatChecker) stop() {
	log.Info(checker.ctx, "Stopping heartbeat checker", log.Fields{})
	close(checker.in)
	checker.ticker.Stop()
	log.Info(checker.ctx, "Heartbeat checker stopped", log.Fields{})
}
