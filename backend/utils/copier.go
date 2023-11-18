package utils

import (
	"github.com/PPG007/copier"
	"pchat/repository/bson"
	"reflect"
)

var (
	objectIdStringConverter = copier.Converter{
		Origin: reflect.TypeOf(bson.ObjectId{}),
		Target: reflect.TypeOf(""),
		Fn: func(fromValue reflect.Value, toType reflect.Type) (toValue reflect.Value, err error) {
			oid, ok := fromValue.Interface().(bson.ObjectId)
			if !ok {
				return fromValue, nil
			}
			return reflect.ValueOf(oid.Hex()), nil
		},
	}
	stringObjectIdConverter = copier.Converter{
		Origin: reflect.TypeOf(""),
		Target: reflect.TypeOf(bson.ObjectId{}),
		Fn: func(fromValue reflect.Value, toType reflect.Type) (toValue reflect.Value, err error) {
			id, ok := fromValue.Interface().(string)
			if !ok {
				return fromValue, nil
			}
			return reflect.ValueOf(bson.NewObjectIdFromHex(id)), nil
		},
	}
)

func Copier() *copier.Copier {
	c := copier.New(true).
		RegisterConverter(copier.TimeStringConverter).
		RegisterConverter(copier.StringTimeConverter).
		RegisterConverter(objectIdStringConverter).
		RegisterConverter(stringObjectIdConverter)
	return c
}
