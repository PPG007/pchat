package todo

import (
	"context"
	"pchat/model"
	"pchat/repository"
	"pchat/repository/bson"
	"time"
)

const (
	C_TODO_RECORD = "todoRecord"
)

var (
	CTodoRecord = &TodoRecord{}
)

type TodoRecord struct {
	Id              bson.ObjectId `bson:"_id"`
	CreatedAt       time.Time     `bson:"createdAt"`
	UpdatedAt       time.Time     `bson:"updatedAt"`
	IsDeleted       bool          `bson:"isDeleted"`
	TodoId          bson.ObjectId `bson:"todoId"`
	UserId          bson.ObjectId `bson:"userId"`
	Content         string        `bson:"content"`
	Images          []string      `bson:"images,omitempty"`
	HasBeenDone     bool          `bson:"hasBeenDone"`
	HasBeenReminded bool          `bson:"hasBeenReminded"`
	DoneAt          time.Time     `bson:"doneAt,omitempty"`
	RemindAt        time.Time     `bson:"remindAt,omitempty"`
}

func (TodoRecord) ListByPagination(ctx context.Context, pagination repository.Pagination) (int64, []TodoRecord, error) {
	var records []TodoRecord
	total, err := repository.FindAllWithPage(ctx, C_TODO_RECORD, pagination, &records)
	return total, records, err
}

func (TodoRecord) ListAllByCondition(ctx context.Context, condition bson.M) ([]TodoRecord, error) {
	var records []TodoRecord
	err := repository.FindAll(ctx, C_TODO_RECORD, condition, &records)
	return records, err
}

func (r TodoRecord) Create(ctx context.Context) error {
	return repository.Insert(ctx, C_TODO_RECORD, r)
}

func (TodoRecord) GetById(ctx context.Context, id bson.ObjectId) (TodoRecord, error) {
	record := TodoRecord{}
	err := repository.FindOne(ctx, C_TODO_RECORD, model.GenIdCondition(id), &record)
	return record, err
}

func (TodoRecord) DeleteByTodoId(ctx context.Context, todoId bson.ObjectId, all bool) error {
	condition := bson.M{
		"todoId":      todoId,
		"hasBeenDone": false,
	}
	if all {
		delete(condition, "hasBeenDone")
	}
	_, err := repository.UpdateAll(ctx, C_TODO_RECORD, condition, bson.M{"$set": bson.M{"isDeleted": true}})
	return err
}

func (TodoRecord) MarkAsDone(ctx context.Context, id bson.ObjectId) error {
	condition := model.GenIdCondition(id)
	updater := bson.M{
		"$set": bson.M{
			"hasBeenDone": true,
			"doneAt":      time.Now(),
			"updatedAt":   time.Now(),
		},
	}
	return repository.UpdateOne(ctx, C_TODO_RECORD, condition, updater)
}

func (TodoRecord) MarkAsUndo(ctx context.Context, id bson.ObjectId) error {
	condition := model.GenIdCondition(id)
	updater := bson.M{
		"$set": bson.M{
			"hasBeenDone": false,
			"updatedAt":   time.Now(),
		},
	}
	return repository.UpdateOne(ctx, C_TODO_RECORD, condition, updater)
}
