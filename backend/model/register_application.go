package model

import (
	"context"
	"github.com/qiniu/qmgo"
	"pchat/repository"
	"pchat/repository/bson"
	"time"
)

const (
	C_REGISTER_APPLICATION = "registerApplication"

	STATUS_APPROVED = "approved"
	STATUS_REJECTED = "rejected"
	STATUS_PENDING  = "pending"
)

var (
	CRegisterApplication = &RegisterApplication{}
)

type RegisterApplication struct {
	Id           bson.ObjectId `bson:"_id"`
	CreatedAt    time.Time     `bson:"createdAt"`
	UpdatedAt    time.Time     `bson:"updatedAt"`
	Reason       string        `bson:"reason"`
	RejectReason string        `bson:"rejectReason"`
	Status       string        `bson:"status"`
	Email        string        `bson:"email"`
	UserId       bson.ObjectId `bson:"userId"`
}

func (*RegisterApplication) Upsert(ctx context.Context, user User, reason string) error {
	application := RegisterApplication{}
	condition := bson.M{
		"email": user.Email,
	}
	change := qmgo.Change{
		Upsert: true,
		Update: bson.M{
			"$set": bson.M{
				"status":    STATUS_PENDING,
				"reason":    reason,
				"updatedAt": time.Now(),
			},
			"$setOnInsert": bson.M{
				"createdAt": time.Now(),
				"userId":    user.Id,
			},
		},
	}
	return repository.FindAndApply(ctx, C_REGISTER_APPLICATION, condition, change, &application)
}

func (*RegisterApplication) Approve(ctx context.Context, ids []bson.ObjectId) error {
	condition := bson.M{
		"_id": bson.M{
			"$in": ids,
		},
	}
	updater := bson.M{
		"$set": bson.M{
			"status": STATUS_APPROVED,
		},
	}
	if err := repository.UpdateOne(ctx, C_REGISTER_APPLICATION, condition, updater); err != nil {
		return err
	}
	return CRegisterApplication.afterApprove(ctx, ids)
}

func (*RegisterApplication) Reject(ctx context.Context, ids []bson.ObjectId, reason string) error {
	condition := bson.M{
		"_id": bson.M{
			"$in": ids,
		},
	}
	updater := bson.M{
		"$set": bson.M{
			"status":       STATUS_REJECTED,
			"rejectReason": reason,
		},
	}
	return repository.UpdateOne(ctx, C_REGISTER_APPLICATION, condition, updater)
}

func (*RegisterApplication) GetByIds(ctx context.Context, ids []bson.ObjectId) ([]RegisterApplication, error) {
	condition := bson.M{
		"_id": bson.M{
			"$in": ids,
		},
	}
	var result []RegisterApplication
	err := repository.FindAll(ctx, C_REGISTER_APPLICATION, condition, &result)
	return result, err
}

func (r *RegisterApplication) afterApprove(ctx context.Context, ids []bson.ObjectId) error {
	registerApplications, err := r.GetByIds(ctx, ids)
	if err != nil {
		return err
	}
	userIds := make([]bson.ObjectId, 0, len(registerApplications))
	for _, application := range registerApplications {
		userIds = append(userIds, application.UserId)
	}
	return CUser.Activate(ctx, userIds)
}

func (*RegisterApplication) ListByPagination(ctx context.Context, pagination repository.Pagination) ([]RegisterApplication, int64, error) {
	var applications []RegisterApplication
	total, err := repository.FindAllWithPage(ctx, C_REGISTER_APPLICATION, pagination, &applications)
	return applications, total, err
}
