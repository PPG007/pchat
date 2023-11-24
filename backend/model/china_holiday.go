package model

import (
	"context"
	"pchat/repository"
	"pchat/repository/bson"
	"time"
)

const (
	C_CHINA_HOLIDAY = "chinaHoliday"

	TIME_FORMATTER = "2006-01-02"
)

var (
	CChinaHoliday = &ChinaHoliday{}
)

type ChinaHoliday struct {
	Id           bson.ObjectId `bson:"_id"`
	IsWorkingDay bool          `bson:"isWorkingDay"`
	Date         time.Time     `bson:"date"`
	DateStr      string        `bson:"dateStr"`
}

func (*ChinaHoliday) GetNextWorkingDay(ctx context.Context) (ChinaHoliday, error) {
	condition := bson.M{
		"date": bson.M{
			"$gt": time.Now(),
		},
		"isWorkingDay": true,
	}
	result := ChinaHoliday{}
	err := repository.FindOne(ctx, C_CHINA_HOLIDAY, condition, &result)
	return result, err
}

func (*ChinaHoliday) GetNextHoliday(ctx context.Context) (ChinaHoliday, error) {
	condition := bson.M{
		"date": bson.M{
			"$gt": time.Now(),
		},
		"isWorkingDay": false,
	}
	result := ChinaHoliday{}
	err := repository.FindOne(ctx, C_CHINA_HOLIDAY, condition, &result)
	return result, err
}

func (*ChinaHoliday) BatchUpsert(ctx context.Context, holidays []ChinaHoliday) error {
	for _, holiday := range holidays {
		condition := bson.M{
			"dateStr": holiday.Date.Format(TIME_FORMATTER),
		}
		updater := bson.M{
			"date":         holiday.Date,
			"dateStr":      holiday.DateStr,
			"isWorkingDay": holiday.IsWorkingDay,
		}
		err := repository.Upsert(ctx, C_CHINA_HOLIDAY, condition, updater)
		if err != nil {
			return err
		}
	}
	return nil
}
