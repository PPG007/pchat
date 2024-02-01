package user

import (
	"context"
	"pchat/repository"
	"pchat/repository/bson"
	"pchat/utils"
)

const (
	C_ROLE = "role"
)

var (
	CRole = &Role{}
)

type Role struct {
	Id          bson.ObjectId `bson:"_id"`
	Permissions []string      `bson:"permissions"`
}

func (*Role) GetPermissionsByIds(ctx context.Context, ids []bson.ObjectId) ([]string, error) {
	condition := bson.M{
		"_id": bson.M{
			"$in": ids,
		},
	}
	var roles []Role
	err := repository.FindAll(ctx, C_ROLE, condition, &roles)
	if err != nil {
		return nil, err
	}
	permissions := make([]string, 0, len(roles))
	for _, role := range roles {
		permissions = append(permissions, role.Permissions...)
	}
	return utils.StrArrUnique(permissions), nil
}
