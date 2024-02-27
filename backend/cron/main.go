package cron

import (
	"context"
	"github.com/robfig/cron/v3"
	"pchat/utils/log"
)

type cronJobFn func(ctx context.Context) error

func cronWrapper(name string, fn cronJobFn) func() {
	return func() {
		ctx := context.Background()
		err := fn(ctx)
		if err != nil {
			log.Error(ctx, "Failed to run cron job", log.Fields{
				"jobName": name,
				"error":   err,
			})
		}
	}
}

type cronJob struct {
	Name string
	Spec string
	Fn   cronJobFn
}

var (
	jobs = make([]cronJob, 0)
)

func Start() {
	c := cron.New()
	for _, job := range jobs {
		c.AddFunc(job.Spec, cronWrapper(job.Name, job.Fn))
	}
	c.Start()
}
