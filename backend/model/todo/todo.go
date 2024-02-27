package todo

import (
	"context"
	"pchat/model"
	"pchat/model/common"
	"pchat/repository"
	"pchat/repository/bson"
	"pchat/utils"
	"time"
)

const (
	C_TODO = "todo"

	REPEAT_TYPE_DAILY   = "day"
	REPEAT_TYPE_WEEK    = "week"
	REPEAT_TYPE_MONTH   = "month"
	REPEAT_TYPE_YEAR    = "year"
	REPEAT_TYPE_HOLIDAY = "holiday"
	REPEAT_TYPE_WEEKDAY = "weekday"
)

var (
	CTodo = &Todo{}
)

type Todo struct {
	Id            bson.ObjectId `bson:"_id"`
	CreatedAt     time.Time     `bson:"createdAt"`
	UpdatedAt     time.Time     `bson:"updatedAt"`
	IsDeleted     bool          `bson:"isDeleted"`
	NeedRemind    bool          `bson:"needRemind"`
	Content       string        `bson:"content"`
	UserId        bson.ObjectId `bson:"userId"`
	Images        []string      `bson:"images,omitempty"`
	RemindSetting RemindSetting `bson:"remindSetting,omitempty"`
}

type RemindSetting struct {
	RemindAt         time.Time `bson:"remindAt"`
	IsRepeatable     bool      `bson:"isRepeatable"`
	LastRemindAt     time.Time `bson:"lastRemindAt,omitempty"`
	RepeatType       string    `bson:"repeatType"`
	RepeatDateOffset int64     `bson:"repeatDateOffset"`
}

func (Todo) GetById(ctx context.Context, id bson.ObjectId) (Todo, error) {
	condition := model.GenIdCondition(id)
	todo := Todo{}
	err := repository.FindOne(ctx, C_TODO, condition, &todo)
	return todo, err
}

func (Todo) ListByIds(ctx context.Context, ids []bson.ObjectId) ([]Todo, error) {
	var todos []Todo
	condition := bson.M{
		"_id": bson.M{
			"$in": ids,
		},
	}
	err := repository.FindAll(ctx, C_TODO, condition, &todos)
	return todos, err
}

func (t Todo) Create(ctx context.Context) error {
	t.Id = bson.NewObjectId()
	t.CreatedAt = time.Now()
	t.UpdatedAt = t.CreatedAt
	t.UserId = utils.GetUserIdAsObjectId(ctx)
	if !t.RemindSetting.RemindAt.IsZero() {
		t.NeedRemind = true
	}
	if err := repository.Insert(ctx, C_TODO, t); err != nil {
		return err
	}
	record := t.GenRecord(ctx)
	if err := record.Create(ctx); err != nil {
		return err
	}
	return t.UpdateLastRemindAt(ctx, record.RemindAt)
}

func (t Todo) GenRecord(ctx context.Context) TodoRecord {
	return TodoRecord{
		Id:        bson.ObjectId{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsDeleted: false,
		TodoId:    t.Id,
		UserId:    t.UserId,
		Content:   t.Content,
		Images:    t.Images,
		RemindAt:  t.RemindSetting.GetNextRemindAt(ctx),
	}
}

func (Todo) DeleteById(ctx context.Context, id bson.ObjectId) error {
	condition := bson.M{
		"_id": id,
	}
	if err := repository.UpdateOne(ctx, C_TODO, condition, bson.M{"$set": bson.M{"isDeleted": true}}); err != nil {
		return err
	}
	return CTodoRecord.DeleteByTodoId(ctx, id, false)
}

func (t Todo) Update(ctx context.Context) error {
	condition := model.GenIdCondition(t.Id)
	updater := bson.M{
		"$set": bson.M{
			"updatedAt":     time.Now(),
			"needRemind":    t.NeedRemind,
			"content":       t.Content,
			"images":        t.Images,
			"remindSetting": t.RemindSetting,
		},
	}
	return repository.UpdateOne(ctx, C_TODO, condition, updater)
}

func (t Todo) UpdateLastRemindAt(ctx context.Context, remindAt time.Time) error {
	condition := model.GenIdCondition(t.Id)
	updater := bson.M{
		"$set": bson.M{
			"remindSetting.lastRemindAt": remindAt,
		},
	}
	return repository.UpdateOne(ctx, C_TODO, condition, updater)
}

func (r RemindSetting) GetNextRemindAt(ctx context.Context) time.Time {
	from := r.LastRemindAt
	if from.IsZero() {
		return r.RemindAt
	}
	if r.IsRepeatable {
		switch r.RepeatType {
		case REPEAT_TYPE_DAILY:
			return from.AddDate(0, 0, int(r.RepeatDateOffset))
		case REPEAT_TYPE_WEEK:
			return from.AddDate(0, 0, int(7*r.RepeatDateOffset))
		case REPEAT_TYPE_MONTH:
			return from.AddDate(0, int(r.RepeatDateOffset), 0)
		case REPEAT_TYPE_YEAR:
			return from.AddDate(int(r.RepeatDateOffset), 0, 0)
		case REPEAT_TYPE_HOLIDAY:
			nextHoliday, _ := common.CChinaHoliday.GetNextHoliday(ctx, from)
			return nextHoliday.Date
		case REPEAT_TYPE_WEEKDAY:
			nextWorkingDay, _ := common.CChinaHoliday.GetNextWorkingDay(ctx, from)
			return nextWorkingDay.Date
		}
	}
	return time.Time{}
}
