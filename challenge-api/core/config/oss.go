package config

type Oss struct {
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	Bucket          string

	// MinIO 专用
	UseSSL bool
}

var OssConfig = new(Oss)
