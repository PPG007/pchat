package main

import (
	"context"
	model_common "pchat/model/common"
)

func InitDefaultResources() {
	ctx := context.Background()
	model_common.CSetting.CreateDefaultSetting(ctx)
}
