package utils

import (
	"github.com/asaskevich/govalidator"
	"github.com/spf13/cast"
	"pchat/repository/bson"
)

func init() {
	govalidator.CustomTypeTagMap.Set("objectId", func(i interface{}, o interface{}) bool {
		return bson.IsObjectIdHex(cast.ToString(i))
	})
	govalidator.CustomTypeTagMap.Set("objectIdList", func(i interface{}, o interface{}) bool {
		strs := cast.ToStringSlice(i)
		for _, str := range strs {
			if !bson.IsObjectIdHex(str) {
				return false
			}
		}
		return true
	})
}

func ValidateRequest(req any) error {
	_, err := govalidator.ValidateStruct(req)
	return err
}
