package oss

import (
	"challenge/core/config"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type AliyunOSS struct {
	client *oss.Client
	bucket *oss.Bucket
	cfg    config.Oss
}

func (a *AliyunOSS) Init(cfg config.Oss) error {
	if cfg.Endpoint == "" ||
		cfg.AccessKeyID == "" ||
		cfg.AccessKeySecret == "" ||
		cfg.Bucket == "" {
		return ErrInvalidConfig
	}

	client, err := oss.New(cfg.Endpoint, cfg.AccessKeyID, cfg.AccessKeySecret)
	if err != nil {
		return fmt.Errorf("init aliyun oss client failed: %w", err)
	}

	bucket, err := client.Bucket(cfg.Bucket)
	if err != nil {
		return fmt.Errorf("get aliyun oss bucket failed: %w", err)
	}

	a.client = client
	a.bucket = bucket
	a.cfg = cfg
	return nil
}

func (a *AliyunOSS) Upload(objectKey, localPath string) error {
	return a.bucket.PutObjectFromFile(objectKey, localPath)
}
func (a *AliyunOSS) UploadReader(objectKey string, r io.Reader, size int64, contentType string) error {
	opts := []oss.Option{}
	if contentType != "" {
		opts = append(opts, oss.ContentType(contentType))
	}
	return a.bucket.PutObject(objectKey, r, opts...)
}
func (a *AliyunOSS) UploadMultipart(objectKey, localPath string) error {
	chunks, err := oss.SplitFileByPartNum(localPath, 5)
	if err != nil {
		return err
	}

	fd, err := os.Open(localPath)
	if err != nil {
		return err
	}
	defer fd.Close()

	imur, err := a.bucket.InitiateMultipartUpload(objectKey)
	if err != nil {
		return err
	}

	var parts []oss.UploadPart
	for _, chunk := range chunks {
		if _, err := fd.Seek(chunk.Offset, 0); err != nil {
			return err
		}

		part, err := a.bucket.UploadPart(imur, fd, chunk.Size, chunk.Number)
		if err != nil {
			return err
		}
		parts = append(parts, part)
	}

	_, err = a.bucket.CompleteMultipartUpload(imur, parts)
	return err
}

func (a *AliyunOSS) GeneratePresignedURL(objectKey string, expire time.Duration) (string, error) {
	return a.bucket.SignURL(objectKey, oss.HTTPGet, int64(expire.Seconds()))
}
