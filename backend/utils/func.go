package utils

import (
	"pchat/repository/bson"
)

func StrInArray(str string, arr *[]string) bool {
	if arr == nil {
		return false
	}
	for _, s := range *arr {
		if s == str {
			return true
		}
	}
	return false
}

func ObjectIdInArray(id bson.ObjectId, arr *[]bson.ObjectId) bool {
	if arr == nil {
		return false
	}
	for _, objectId := range *arr {
		if bson.IsObjectIdEqual(objectId, id) {
			return true
		}
	}
	return false
}

func StrArrUnique(arr []string) []string {
	m := make(map[string]struct{}, len(arr))
	for _, s := range arr {
		m[s] = struct{}{}
	}
	result := make([]string, 0, len(m))
	for k, _ := range m {
		result = append(result, k)
	}
	return result
}

func ObjectIdsToStrArray(ids []bson.ObjectId) []string {
	result := make([]string, 0, len(ids))
	for _, id := range ids {
		result = append(result, id.Hex())
	}
	return result
}

func StrArrToObjectIds(ids []string) []bson.ObjectId {
	result := make([]bson.ObjectId, 0, len(ids))
	for _, id := range ids {
		result = append(result, bson.NewObjectIdFromHex(id))
	}
	return result
}
