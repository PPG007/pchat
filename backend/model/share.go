package model

import "pchat/repository/bson"

func GenIdCondition(id bson.ObjectId) bson.M {
	return bson.M{
		"_id": id,
	}
}
