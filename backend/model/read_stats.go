package model

import (
	"pchat/repository/bson"
	"time"
)

const (
	C_READ_STATS = "readStats"
)

var (
	CReadStats = &ReadStats{}
)

type ReadStats struct {
	Id        bson.ObjectId `bson:"_id"`
	MessageId bson.ObjectId `bson:"messageId"`
	UserId    bson.ObjectId `bson:"userId"`
	CreatedAt time.Time     `bson:"createdAt"`
}
