package utils

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"pchat/repository/bson"
	"time"
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

func GenerateRandomSecretKey(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(randomBytes), nil
}

func ParseSecretString(key string) []byte {
	byteKey, _ := hex.DecodeString(key)
	return byteKey
}

const (
	USER_ID_HEADER = "X-User-Id"
)

func GetUserId(ctx context.Context) string {
	if ginCtx, ok := ctx.(*gin.Context); ok {
		return ginCtx.GetHeader(USER_ID_HEADER)
	}
	return ""
}

func GetFirstDayInYear(arg time.Time) time.Time {
	return time.Date(arg.Year(), time.January, 1, 0, 0, 0, 0, time.Local)
}

func GetLastDayInYear(arg time.Time) time.Time {
	return time.Date(arg.Year()+1, time.January, 1, 0, 0, 0, 0, time.Local).AddDate(0, 0, -1)
}
