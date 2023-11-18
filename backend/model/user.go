package model

import (
	"pchat/repository/bson"
	"time"
)

const (
	C_USER = "user"

	C_USER_STATUS_ACTIVATED = "activated"
	C_USER_STATUS_BLOCKED   = "blocked"
	C_USER_STATUS_AUDITING  = "auditing"

	C_USER_CHAT_STATUS_ONLINE  = "online"
	C_USER_CHAT_STATUS_OFFLINE = "offline"
	C_USER_CHAT_STATUS_LEAVING = "leaving"
	C_USER_CHAT_STATUS_BUSY    = "busy"
)

var (
	CUser = &User{}
)

type User struct {
	Id         bson.ObjectId   `bson:"_id"`
	Name       string          `bson:"name"`
	Password   string          `bson:"password"`
	Email      string          `bson:"email"`
	Roles      []bson.ObjectId `bson:"roles"`
	CreatedAt  time.Time       `bson:"createdAt"`
	UpdatedAt  time.Time       `bson:"updatedAt"`
	Status     string          `bson:"status"`
	Avatar     string          `bson:"avatar"`
	ChatStatus string          `bson:"chatStatus"`
}
