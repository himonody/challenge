package oss

import (
	"challenge/core/config"
	"errors"
	"io"
	"time"
)

var (
	ErrUnsupportedDriver = errors.New("unsupported storage driver")
	ErrInvalidConfig     = errors.New("invalid storage config")
)

const (
	AliYun = "aliyun"
	MinIO  = "minio"
)

// ObjectStorage 对象存储统一接口
type ObjectStorage interface {
	Init(cfg config.Oss) error

	// Upload 普通上传
	Upload(objectKey, localPath string) error

	UploadReader(objectKey string, r io.Reader, size int64, contentType string) error // 新增

	// UploadMultipart 分片/大文件上传
	UploadMultipart(objectKey, localPath string) error

	// GeneratePresignedURL 获取访问链接
	GeneratePresignedURL(objectKey string, expire time.Duration) (string, error)
}

// New 创建对象存储实例
func New(driver string, cfg config.Oss) (ObjectStorage, error) {
	var s ObjectStorage

	switch driver {
	case AliYun:
		s = &AliyunOSS{}
	case MinIO:
		s = &MinioStorage{}
	default:
		return nil, ErrUnsupportedDriver
	}

	return s, s.Init(cfg)
}
