package oss

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"net/url"
	"os"
	"pchat/errors"
	model_common "pchat/model/common"
	"time"
)

type minioOSSClient struct {
	bucketName string
	client     *minio.Client
}

func (m *minioOSSClient) SetBucket(ctx context.Context, bucketName string) error {
	exists, err := m.client.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New(errors.ERR_COMMON_BUCKETS_NOT_EXISTS, "bucket not exists")
	}
	m.bucketName = bucketName
	return nil
}

func (m *minioOSSClient) GetObject(ctx context.Context, key, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	object, err := m.client.GetObject(ctx, m.bucketName, key, minio.GetObjectOptions{})
	if err != nil {
		return err
	}
	_, err = io.Copy(file, object)
	return err
}

func (m *minioOSSClient) PutObject(ctx context.Context, key, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		return err
	}
	_, err = m.client.PutObject(ctx, m.bucketName, key, file, info.Size(), minio.PutObjectOptions{})
	return err
}

func (m *minioOSSClient) RemoveObject(ctx context.Context, key string) error {
	return m.client.RemoveObject(ctx, m.bucketName, key, minio.RemoveObjectOptions{})
}

func (m *minioOSSClient) HasObject(ctx context.Context, key string) (bool, error) {
	info, _ := m.client.StatObject(ctx, m.bucketName, key, minio.StatObjectOptions{})
	if info.Size > 0 {
		return true, nil
	}
	return false, nil
}

func (m *minioOSSClient) CopyObject(ctx context.Context, fromKey, toKey string) error {
	_, err := m.client.CopyObject(ctx, minio.CopyDestOptions{
		Bucket: m.bucketName,
		Object: toKey,
	}, minio.CopySrcOptions{
		Bucket: m.bucketName,
		Object: fromKey,
	})
	return err
}

func (m *minioOSSClient) RenameObject(ctx context.Context, fromKey, toKey string) error {
	err := m.CopyObject(ctx, fromKey, toKey)
	if err != nil {
		return err
	}
	return m.RemoveObject(ctx, fromKey)
}

func (m *minioOSSClient) StatObject(ctx context.Context, key string) (*ObjectMetaInfo, error) {
	info, err := m.client.StatObject(ctx, m.bucketName, key, minio.StatObjectOptions{})
	if err != nil {
		return nil, err
	}
	return &ObjectMetaInfo{
		Size:      info.Size,
		UpdatedAt: info.LastModified,
	}, nil
}

func (m *minioOSSClient) SignObjectURL(ctx context.Context, key string, expireDuration time.Duration) (string, error) {
	u, err := m.client.PresignedGetObject(ctx, m.bucketName, key, expireDuration, nil)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

func (m *minioOSSClient) SignPutObjectURL(ctx context.Context, key string, expireDuration time.Duration) (string, error) {
	u, err := m.client.PresignedPutObject(ctx, m.bucketName, key, expireDuration)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

func newMinioOSSClient(ctx context.Context, setting model_common.OSSSetting) (*minioOSSClient, error) {
	u, err := url.Parse(setting.Endpoint)
	if err != nil {
		return nil, err
	}
	client, err := minio.New(u.Host, &minio.Options{
		Creds:  credentials.NewStaticV4(setting.AccessKey, setting.SecretAccessKey, ""),
		Secure: u.Scheme == "https",
	})
	if err != nil {
		return nil, err
	}
	return &minioOSSClient{
		"",
		client,
	}, nil
}
