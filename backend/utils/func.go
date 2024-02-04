package utils

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	pb_common "pchat/pb/common"
	"pchat/repository"
	"pchat/repository/bson"
	"pchat/utils/log"
	"runtime"
	"strings"
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

func GetFirstDayInYear(arg time.Time) time.Time {
	return time.Date(arg.Year(), time.January, 1, 0, 0, 0, 0, time.Local)
}

func GetLastDayInYear(arg time.Time) time.Time {
	return time.Date(arg.Year()+1, time.January, 1, 0, 0, 0, 0, time.Local).AddDate(0, 0, -1)
}

func FormatListCondition(listCondition *pb_common.ListCondition) *pb_common.ListCondition {
	if listCondition == nil {
		listCondition = &pb_common.ListCondition{}
	}
	if listCondition.Page == 0 {
		listCondition.Page = 1
	}
	if listCondition.PerPage == 0 {
		listCondition.PerPage = 10
	}
	if len(listCondition.OrderBy) == 0 {
		listCondition.OrderBy = []string{"-createdAt"}
	}
	return listCondition
}

func FormatPagination(condition bson.M, listCondition *pb_common.ListCondition) repository.Pagination {
	listCondition = FormatListCondition(listCondition)
	return repository.Pagination{
		Condition: condition,
		Page:      listCondition.Page,
		PerPage:   listCondition.PerPage,
		OrderBy:   listCondition.OrderBy,
	}
}

func GO(ctx context.Context, fn func()) {
	go func() {
		if r := recover(); r != nil {
			stack := make([]byte, 4096)
			stack = stack[:runtime.Stack(stack, false)]
			err := fmt.Sprintf("%v", r)
			log.ErrorTrace(ctx, "Panic in Goroutine", log.Fields{
				"error": err,
			}, stack)
		}
		fn()
	}()
}

func MarshalInterfaceToString(obj interface{}) string {
	b, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(b)
}

func FormatRegexStr(str string) string {
	replacers := []string{
		"\\", "\\\\",
		"*", "\\*",
		".", "\\.",
		"?", "\\?",
		"+", "\\+",
		"$", "\\$",
		"^", "\\^",
		"[", "\\[",
		"]", "\\]",
		"(", "\\(",
		")", "\\)",
		"{", "\\{",
		"}", "\\}",
		"|", "\\|",
		"/", "\\/",
	}
	return strings.NewReplacer(replacers...).Replace(str)
}

func GetFuzzySearchStrRegex(str string) bson.Regex {
	return bson.Regex{
		Pattern: FormatRegexStr(str),
		Options: "i",
	}
}
