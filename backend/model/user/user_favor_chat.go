package user

import "pchat/repository/bson"

const (
	C_USER_FAVOR_CHAT = "userFavorChat"
)

var (
	CUserFavorChat = &UserFavorChat{}
)

type UserFavorChat struct {
	Id     bson.ObjectId `bson:"_id"`
	ChatId bson.ObjectId `bson:"chatId"`
	UserId bson.ObjectId `bson:"userId"`
}
