package model

import (
	"pchat/repository/bson"
	"time"
)

const (
	C_CHAT = "chat"

	CHAT_TYPE_GROUP  = "group"
	CHAT_TYPE_DIRECT = "direct"
)

var (
	CChat = &Chat{}
)

type Chat struct {
	Id          bson.ObjectId   `bson:"_id"`
	IsDeleted   bool            `bson:"isDeleted"`
	CreatedAt   time.Time       `bson:"createdAt"`
	Type        string          `bson:"type"`
	Members     []bson.ObjectId `bson:"members"`
	Avatar      string          `bson:"avatar"`
	IsPrivate   bool            `bson:"isPrivate"`
	LastMessage BriefMessage    `bson:"lastMessage,omitempty"`
}

type BriefMessage struct {
	Id        bson.ObjectId `bson:"id"`
	CreatedAt time.Time     `bson:"createdAt"`
	Sender    bson.ObjectId `bson:"sender"`
	Content   string        `bson:"content"`
}
