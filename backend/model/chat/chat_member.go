package chat

import (
	"pchat/repository/bson"
	"time"
)

const (
	C_CHAT_MEMBER = "chatMember"

	ROLE_OWNER = "owner"
	ROLE_ADMIN = "admin"
)

type ChatMember struct {
	Id        bson.ObjectId `bson:"_id"`
	CreatedAt time.Time     `bson:"createdAt"`
	UpdatedAt time.Time     `bson:"updatedAt"`
	UserId    bson.ObjectId `bson:"userId"`
	ChatId    bson.ObjectId `bson:"chatId"`
	Role      string        `bson:"role"`
}
