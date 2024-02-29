package oss

import (
	"context"
	"pchat/errors"
	model_common "pchat/model/common"
	"time"
)

type ObjectMetaInfo struct {
	Size      int64
	UpdatedAt time.Time
}

type IOSS interface {
	SetBucket(ctx context.Context, bucketName string) error
	GetObject(ctx context.Context, key, filePath string) error
	PutObject(ctx context.Context, key, filePath string) error
	RemoveObject(ctx context.Context, key string) error
	HasObject(ctx context.Context, key string) (bool, error)
	CopyObject(ctx context.Context, fromKey, toKey string) error
	RenameObject(ctx context.Context, fromKey, toKey string) error
	StatObject(ctx context.Context, key string) (*ObjectMetaInfo, error)
	SignObjectURL(ctx context.Context, key string, expireDuration time.Duration) (string, error)
	SignPutObjectURL(ctx context.Context, key string, expireDuration time.Duration) (string, error)
}

func GetOSSProvider(ctx context.Context) (IOSS, error) {
	setting, err := model_common.CSetting.GetWithCache(ctx)
	if err != nil {
		return nil, err
	}
	switch setting.OSS.Provider {
	case "minio":
		return newMinioOSSClient(ctx, setting.OSS)
	default:
		return nil, errors.New(errors.ERR_COMMON_UNSUPPORTED_OSS_PROVIDER, "unsupported oss provider")
	}
}
