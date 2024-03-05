package admin

import (
	"context"
	model_common "pchat/model/common"
	pb_admin "pchat/pb/admin"
	pb_common "pchat/pb/common"
	"pchat/utils"
	"strings"
)

// @Description
// @Router		/admin/setting [put]
// @Tags		设置
// @Summary	修改设置
// @Accept		json
// @Produce	json
// @Success	200		{object}	pb_common.EmptyResponse
// @Param		body	body		pb_admin.UpdateSettingRequest	true	"body"
func updateSetting(ctx context.Context, req *pb_admin.UpdateSettingRequest) (*pb_common.EmptyResponse, error) {
	setting, err := model_common.CSetting.Get(ctx)
	if err != nil {
		return nil, err
	}
	if req.Smtp != nil {
		if err := utils.Copier().RegisterTransformer("Protocol", func(p pb_admin.UpdateSMTPSettingRequest_Protocol) string {
			return strings.ToLower(pb_admin.UpdateSMTPSettingRequest_Protocol_name[int32(p)])
		}).From(req.Smtp).To(&setting.SMTP); err != nil {
			return nil, err
		}
	}
	if req.Oss != nil {
		if err := utils.Copier().RegisterTransformer("Provider", func(p pb_admin.UpdateOSSSettingRequest_OSSProvider) string {
			return pb_admin.UpdateOSSSettingRequest_OSSProvider_name[int32(p)]
		}).From(req.Oss).To(&setting.OSS); err != nil {
			return nil, err
		}
	}
	if req.Ai != nil {
		if err := utils.Copier().RegisterTransformer("Provider", func(p pb_admin.UpdateAIRequest_AIProvider) string {
			return pb_admin.UpdateSMTPSettingRequest_Protocol_name[int32(p)]
		}).From(req.Ai).To(&setting.AI); err != nil {
			return nil, err
		}
	}
	if req.Chat != nil {
		if err := utils.Copier().From(req.Chat).To(&setting.Chat); err != nil {
			return nil, err
		}
	}
	if req.Account != nil {
		if err := utils.Copier().From(req.Account).To(&setting.Account); err != nil {
			return nil, err
		}
	}
	if err := setting.Update(ctx); err != nil {
		return nil, err
	}
	return &pb_common.EmptyResponse{}, nil
}
