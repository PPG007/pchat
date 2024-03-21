package model

import (
	"context"
	"pchat/model/chat"
	"pchat/model/user"
	"pchat/repository"
)

var (
	indexes = map[string][]repository.IndexOption{
		user.C_PERMISSION: {
			{
				Fields: []repository.IndexField{
					{
						Name: "name",
						Desc: false,
					},
				},
				IsUnique:          true,
				PartialExpression: nil,
			},
		},
		chat.C_CHAT_MEMBER: {
			{
				Fields: []repository.IndexField{
					{
						Name: "userId",
						Desc: false,
					},
					{
						Name: "chatId",
						Desc: false,
					},
				},
				IsUnique: true,
			},
		},
	}
)

func CreateIndexes(ctx context.Context) error {
	for collection, options := range indexes {
		for _, option := range options {
			if err := repository.CreateIndex(ctx, collection, option); err != nil {
				return err
			}
		}
	}
	return nil
}
