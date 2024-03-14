package main

import (
	"context"
	"pchat/model"
	model_common "pchat/model/common"
	model_user "pchat/model/user"
	"pchat/utils/env"
	"pchat/utils/log"
)

func InitDefaultResources() {
	ctx := context.Background()
	if err := model_common.CSetting.CreateDefaultSetting(ctx); err != nil && env.IsDebug() {
		log.Error(ctx, "Failed to create default setting", log.Fields{
			"error": err,
		})
	}
	if err := model_user.CPermission.Init(ctx); err != nil && env.IsDebug() {
		log.Error(ctx, "Failed to init permissions", log.Fields{
			"error": err,
		})
	}
	if err := model.CreateIndexes(ctx); err != nil && env.IsDebug() {
		log.Error(ctx, "Failed to create indexes", log.Fields{
			"error": err,
		})
	}
}
