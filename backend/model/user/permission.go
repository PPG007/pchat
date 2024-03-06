package user

import (
	"context"
	"pchat/permissions"
	"pchat/repository"
	"pchat/repository/bson"
)

const (
	C_PERMISSION = "permission"
)

var (
	CPermission = &Permission{}
)

type Permission struct {
	Id   bson.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
}

func (*Permission) Init(ctx context.Context) error {
	for _, permission := range permissions.AllPermissions {
		condition := bson.M{
			"name": permission,
		}
		updater := bson.M{
			"$set": bson.M{
				"name": permission,
			},
		}
		if err := repository.Upsert(ctx, C_PERMISSION, condition, updater); err != nil {
			return err
		}
	}
	return nil
}
