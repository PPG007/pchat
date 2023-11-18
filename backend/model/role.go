package model

import "pchat/repository/bson"

const (
	C_ROLE = "role"
)

var (
	CRole = &Role{}
)

type Role struct {
	Id          bson.ObjectId `bson:"_id"`
	Permissions []string      `bson:"permissions"`
}
