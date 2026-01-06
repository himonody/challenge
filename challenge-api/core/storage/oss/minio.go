package oss

import (
	"challenge/core/config"
	"context"
	"fmt"
	"io"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioStorage struct {
	client *minio.Client
	cfg    config.Oss
}

func (m *MinioStorage) Init(cfg config.Oss) error {
	if cfg.Endpoint == "" ||
		cfg.AccessKeyID == "" ||
		cfg.AccessKeySecret == "" ||
		cfg.Bucket == "" {
		return ErrInvalidConfig
	}

	client, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKeyID, cfg.AccessKeySecret, ""),
		Secure: cfg.UseSSL,
	})
	if err != nil {
		return fmt.Errorf("init minio client failed: %w", err)
	}

	ctx := context.Background()
	exists, err := client.BucketExists(ctx, cfg.Bucket)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("bucket %s not exists", cfg.Bucket)
	}

	m.client = client
	m.cfg = cfg
	return nil
}

func (m *MinioStorage) Upload(objectKey, localPath string) error {
	_, err := m.client.FPutObject(
		context.Background(),
		m.cfg.Bucket,
		objectKey,
		localPath,
		minio.PutObjectOptions{},
	)
	return err
}
func (m *MinioStorage) UploadReader(objectKey string, r io.Reader, size int64, contentType string) error {
	_, err := m.client.PutObject(
		context.Background(),
		m.cfg.Bucket,
		objectKey,
		r,
		size,
		minio.PutObjectOptions{ContentType: contentType},
	)
	return err
}

// MinIO SDK 内置自动分片（推荐）
func (m *MinioStorage) UploadMultipart(objectKey, localPath string) error {
	_, err := m.client.FPutObject(
		context.Background(),
		m.cfg.Bucket,
		objectKey,
		localPath,
		minio.PutObjectOptions{
			PartSize: 64 * 1024 * 1024, // 64MB
		},
	)
	return err
}

func (m *MinioStorage) GeneratePresignedURL(objectKey string, expire time.Duration) (string, error) {
	url, err := m.client.PresignedGetObject(
		context.Background(),
		m.cfg.Bucket,
		objectKey,
		expire,
		nil,
	)
	if err != nil {
		return "", err
	}
	return url.String(), nil
}
