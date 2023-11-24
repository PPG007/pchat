package cron

import "github.com/robfig/cron/v3"

func Start() {
	c := cron.New()
	c.AddFunc("0 0 1 * *", SyncHolidays)
	c.Start()
}
