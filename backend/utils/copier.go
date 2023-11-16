package utils

import (
	"github.com/jinzhu/copier"
	"pchat/repository/bson"
	"time"
)

type Copier struct {
	from   any
	option copier.Option
}

func ObjectIdStringConverter() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: bson.ObjectId{},
		DstType: "",
		Fn: func(src interface{}) (dst interface{}, err error) {
			oid, _ := src.(bson.ObjectId)
			return oid.Hex(), nil
		},
	}
}

func TimeStringConverter() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: time.Time{},
		DstType: "",
		Fn: func(src interface{}) (dst interface{}, err error) {
			t, _ := src.(time.Time)
			return t.Format(time.RFC3339), nil
		},
	}
}

func StringObjectIdConverter() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: "",
		DstType: bson.ObjectId{},
		Fn: func(src interface{}) (dst interface{}, err error) {
			id, _ := src.(string)
			if bson.IsObjectIdHex(id) {
				return bson.NewObjectIdFromHex(id), nil
			}
			return bson.NilObjectId, nil
		},
	}
}

func StringTimeConverter() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: "",
		DstType: time.Time{},
		Fn: func(src interface{}) (dst interface{}, err error) {
			str, _ := src.(string)
			t, _ := time.Parse(time.RFC3339, str)
			return t, nil
		},
	}
}

func EmptyCopier() *Copier {
	return new(Copier)
}

func DefaultCopier() *Copier {
	c := EmptyCopier()
	c.RegisterConverter(ObjectIdStringConverter())
	c.RegisterConverter(StringObjectIdConverter())
	c.RegisterConverter(TimeStringConverter())
	c.RegisterConverter(StringTimeConverter())
	return c
}

func (c *Copier) From(from any) *Copier {
	c.from = from
	return c
}

func (c *Copier) To(to any) error {
	return copier.CopyWithOption(to, c.from, c.option)
}

func (c *Copier) RegisterDiffPair(diffPair []copier.FieldNameMapping) *Copier {
	c.option.FieldNameMapping = append(c.option.FieldNameMapping, diffPair...)
	return c
}

func (c *Copier) RegisterConverter(converter copier.TypeConverter) *Copier {
	c.option.Converters = append(c.option.Converters, converter)
	return c
}

func (c *Copier) IgnoreEmpty() *Copier {
	c.option.IgnoreEmpty = true
	return c
}

func (c *Copier) DeepCopy() *Copier {
	c.option.DeepCopy = true
	return c
}

func (c *Copier) CaseSensitive() *Copier {
	c.option.CaseSensitive = true
	return c
}
