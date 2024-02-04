package todo

import (
	"context"
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
