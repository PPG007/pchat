package model

import (
	"context"
	"pchat/repository"
)

var (
	indexes = map[string][]repository.IndexOption{
		"permission": {
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
