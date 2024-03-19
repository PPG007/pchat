package common

import (
	"context"
	model_common "pchat/model/common"
	pb_common "pchat/pb/common"
	"pchat/utils/oss"
	"time"
)

// @Description
// @Router		/common/putObjectURL [get]
// @Tags		通用
// @Summary	获取上传到 oss 的链接
// @Accept		json
// @Produce	json
// @Param		body	body		pb_common.GetPutObjectURLRequest	true	"body"
// @Success	200		{object}	pb_common.GetPutObjectURLResponse
func getPutObjectURL(ctx context.Context, req *pb_common.GetPutObjectURLRequest) (*pb_common.GetPutObjectURLResponse, error) {
	setting, err := model_common.CSetting.GetWithCache(ctx)
	if err != nil {
		return nil, err
	}
	client, err := oss.GetOSSClient(ctx, setting.OSS)
	if err != nil {
		return nil, err
	}
	duration := time.Hour
	if req.ValidSecond > 0 {
		duration = time.Duration(req.ValidSecond) * time.Second
	}
	urlStr, err := client.SignPutObjectURL(ctx, req.Key, duration)
	if err != nil {
		return nil, err
	}
	return &pb_common.GetPutObjectURLResponse{
		Url: urlStr,
	}, nil
}
