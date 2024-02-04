package model

import (
	"context"
	"pchat/repository/bson"
	"pchat/utils"
)

func GenIdCondition(id bson.ObjectId) bson.M {
	return bson.M{
		"_id": id,
	}
}

func GenDefaultUserIdCondition(ctx context.Context) bson.M {
	return bson.M{
		"userId": utils.GetUserIdAsObjectId(ctx),
	}
}
