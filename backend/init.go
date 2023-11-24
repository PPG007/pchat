package main

import (
	"context"
	"pchat/model"
)

func InitDefaultResources() {
	ctx := context.Background()
	model.CSetting.CreateDefaultSetting(ctx)
}
