package main

import (
	"context"
	model_common "pchat/model/common"
	model_user "pchat/model/user"
)

func InitDefaultResources() {
	ctx := context.Background()
	model_common.CSetting.CreateDefaultSetting(ctx)
	model_user.CPermission.Init(ctx)
}
