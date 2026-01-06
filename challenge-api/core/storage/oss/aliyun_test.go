package oss

import (
	"challenge/core/config"
	"testing"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func TestAliyunOSS_GeneratePresignedURL(t *testing.T) {
	type fields struct {
		client *oss.Client
		bucket *oss.Bucket
		cfg    config.Oss
	}
	type args struct {
		objectKey string
		expire    time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AliyunOSS{
				client: tt.fields.client,
				bucket: tt.fields.bucket,
				cfg:    tt.fields.cfg,
			}
			got, err := a.GeneratePresignedURL(tt.args.objectKey, tt.args.expire)
			if (err != nil) != tt.wantErr {
				t.Errorf("GeneratePresignedURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GeneratePresignedURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAliyunOSS_Init(t *testing.T) {
	type fields struct {
		client *oss.Client
		bucket *oss.Bucket
		cfg    config.Oss
	}
	type args struct {
		cfg config.Oss
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AliyunOSS{
				client: tt.fields.client,
				bucket: tt.fields.bucket,
				cfg:    tt.fields.cfg,
			}
			if err := a.Init(tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAliyunOSS_Upload(t *testing.T) {
	type fields struct {
		client *oss.Client
		bucket *oss.Bucket
		cfg    config.Oss
	}
	type args struct {
		objectKey string
		localPath string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AliyunOSS{
				client: tt.fields.client,
				bucket: tt.fields.bucket,
				cfg:    tt.fields.cfg,
			}
			if err := a.Upload(tt.args.objectKey, tt.args.localPath); (err != nil) != tt.wantErr {
				t.Errorf("Upload() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAliyunOSS_UploadMultipart(t *testing.T) {
	type fields struct {
		client *oss.Client
		bucket *oss.Bucket
		cfg    config.Oss
	}
	type args struct {
		objectKey string
		localPath string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AliyunOSS{
				client: tt.fields.client,
				bucket: tt.fields.bucket,
				cfg:    tt.fields.cfg,
			}
			if err := a.UploadMultipart(tt.args.objectKey, tt.args.localPath); (err != nil) != tt.wantErr {
				t.Errorf("UploadMultipart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
