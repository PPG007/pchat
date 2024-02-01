package user

import "pchat/repository/bson"

const (
	C_USER_UNREAD_MESSAGE = "userUnreadMessage"
)

var (
	CUserUnreadMessage = &UserUnreadMessage{}
)

type UserUnreadMessage struct {
	Id        bson.ObjectId `bson:"_id"`
	MessageId bson.ObjectId `bson:"messageId"`
	ChatId    bson.ObjectId `bson:"chatId"`
	UserId    bson.ObjectId `bson:"userId"`
}
