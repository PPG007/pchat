package utils

import (
	"errors"
	"fmt"
	"github.com/PPG007/copier"
	"github.com/spf13/cast"
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
			if oid.IsZero() {
				return reflect.ValueOf(""), nil
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
			if id == "" {
				return reflect.ValueOf(bson.NilObjectId), nil
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

func SetStructFields(class interface{}, fields map[string]string) error {
	v := reflect.ValueOf(class)
	if v.Type().Kind() != reflect.Ptr {
		return errors.New("struct must be pointer")
	}
	v = v.Elem()
	t := v.Type()
	for name, value := range fields {
		name = UppercaseFirst(name)
		if field, ok := t.FieldByName(name); ok {
			fromV := reflect.ValueOf(value)
			target := v.FieldByName(name)
			if fromV.CanConvert(t) {
				target.Set(fromV.Convert(t))
				continue
			}
			switch target.Type().Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				v, err := cast.ToInt64E(value)
				if err != nil {
					return err
				}
				target.SetInt(v)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				v, err := cast.ToUint64E(value)
				if err != nil {
					return err
				}
				target.SetUint(v)
			default:
				return errors.New(fmt.Sprintf("cannot convert %s from string to %s", field.Name, target.Type().Kind().String()))
			}
		}
	}
	return nil
}
