package cron

import (
	"context"
	model_todo "pchat/model/todo"
	"pchat/utils"
	"sync"
)

func init() {
	jobs = append(jobs, cronJob{
		Name: "SendTodoRemind",
		Spec: "@every 20s",
		Fn:   SendTodoRemind,
	})
}

func SendTodoRemind(ctx context.Context) error {
	records, err := model_todo.CTodoRecord.ListNeedRemindRecords(ctx)
	if err != nil {
		return err
	}
	wg := &sync.WaitGroup{}
	pool, err := utils.NewGoroutinePoolWithPanicHandler(10)
	if err != nil {
		return err
	}
	defer pool.Release()
	for _, record := range records {
		temp := record
		wg.Add(1)
		err := pool.Submit(func() {
			defer wg.Done()
			temp.SendRemindMessage(ctx)
		})
		if err != nil {
			wg.Done()
		}
	}
	wg.Wait()
	return nil
}
