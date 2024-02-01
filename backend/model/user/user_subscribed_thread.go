package user

import "pchat/repository/bson"

const (
	C_USER_SUBSCRIBED_THREAD = "userSubscribedThread"
)

var (
	CUserSubscribedThread = &UserSubscribedThread{}
)

type UserSubscribedThread struct {
	Id       bson.ObjectId `bson:"_id"`
	ChatId   bson.ObjectId `bson:"chatId"`
	ThreadId bson.ObjectId `bson:"threadId"`
}
