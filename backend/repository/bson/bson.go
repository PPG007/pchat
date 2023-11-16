package bson

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type M = primitive.M

type ObjectId = primitive.ObjectID

var NilObjectId = NewObjectIdFromHex("000000000000000000000000")

func IsObjectIdEqual(a, b ObjectId) bool {
	return a.Hex() == b.Hex()
}

func NewObjectIdFromHex(id string) ObjectId {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(fmt.Sprintf("invalid objectId: %s", id))
	}
	return oid
}

func NewObjectId() ObjectId {
	return primitive.NewObjectID()
}

func NewObjectIdFromTime(t time.Time) ObjectId {
	return primitive.NewObjectIDFromTimestamp(t)
}

func IsObjectIdHex(str string) bool {
	return primitive.IsValidObjectID(str)
}
