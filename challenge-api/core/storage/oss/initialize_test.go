package oss

import (
	"challenge/core/config"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		driver string
		cfg    config.Oss
	}
	tests := []struct {
		name    string
		args    args
		want    ObjectStorage
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.driver, tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}
