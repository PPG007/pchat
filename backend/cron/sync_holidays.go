package cron

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	model_common "pchat/model/common"
	"pchat/repository/bson"
	"pchat/utils/log"
	"time"
)

const (
	HOLIDAY_URL = "http://api.haoshenqi.top/holiday?date=%d"
)

const (
	STATUS_NORMAL_WORKING_DAY = iota
	STATUS_WEEKEND
	STATUS_WORKING_DAY
	STATUS_HOLIDAY
)

type Holiday struct {
	Status int    `json:"status"`
	Date   string `json:"date"`
}

func init() {
	jobs = append(jobs, cronJob{
		Name: "SyncHolidays",
		Spec: "0 0 1 * *",
		Fn:   SyncHolidays,
	})
	SyncHolidays(context.Background())
}

func SyncHolidays(ctx context.Context) error {
	var (
		err error
	)
	defer func() {
		if err != nil {
			log.Warn(ctx, "Failed to sync china holiday", log.Fields{
				"error": err,
			})
		}
	}()
	url := fmt.Sprintf(HOLIDAY_URL, time.Now().Year())
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var holidays []Holiday
	err = json.Unmarshal(bytes, &holidays)
	if err != nil {
		return err
	}
	var holidayModels []model_common.ChinaHoliday
	for _, holiday := range holidays {
		t, err := time.Parse("2006-01-02", holiday.Date)
		if err != nil {
			continue
		}
		holidayModels = append(holidayModels, model_common.ChinaHoliday{
			Id:           bson.NewObjectId(),
			IsWorkingDay: holiday.Status == STATUS_NORMAL_WORKING_DAY || holiday.Status == STATUS_WORKING_DAY,
			DateStr:      holiday.Date,
			Date:         t,
		})
	}
	return model_common.CChinaHoliday.BatchUpsert(ctx, holidayModels)
}
