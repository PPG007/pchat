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

type PostPolicyOption struct {
	Size           *ObjectSizeRange
	ExpireDuration *time.Duration
}

type ObjectSizeRange struct {
	Min int64
	Max int64
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

func GetOSSClient(ctx context.Context, setting model_common.OSSSetting) (IOSS, error) {
	switch setting.Provider {
	case "minio":
		return newMinioOSSClient(ctx, setting)
	default:
		return nil, errors.New(errors.ERR_COMMON_UNSUPPORTED_OSS_PROVIDER, "unsupported oss provider")
	}
}
