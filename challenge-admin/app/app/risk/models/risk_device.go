package models

import "time"

type RiskDevice struct {
	Id        int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:记录ID"`
	DeviceFp  string     `json:"deviceFp" gorm:"column:device_fp;type:varchar(64);comment:设备指纹"`
	UserId    int64      `json:"userId" gorm:"column:user_id;type:bigint;comment:用户ID"`
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:记录时间"`
}

func (RiskDevice) TableName() string { return "app_risk_device" }
