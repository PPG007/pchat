package user

import "pchat/repository/bson"

const (
	C_PERMISSION = "permission"
)

var (
	CPermission = &Permission{}
)

type Permission struct {
	Id   bson.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
}
