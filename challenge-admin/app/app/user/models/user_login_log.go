package models

import "time"

type UserLoginLog struct {
	Id         int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:登录日志ID"`
	UserId     int64      `json:"userId" gorm:"column:user_id;type:bigint;comment:用户ID"`
	LoginAt    *time.Time `json:"loginAt" gorm:"column:login_at;type:datetime;comment:登录时间"`
	LoginIp    string     `json:"loginIp" gorm:"column:login_ip;type:varchar(45);comment:登录IP"`
	DeviceFp   string     `json:"deviceFp" gorm:"column:device_fp;type:varchar(64);comment:设备指纹"`
	UserAgent  string     `json:"userAgent" gorm:"column:user_agent;type:varchar(500);comment:UA信息"`
	Status     int8       `json:"status" gorm:"column:status;type:tinyint;comment:登录状态 1成功 2失败 3风控拦截"`
	FailReason string     `json:"failReason" gorm:"column:fail_reason;type:varchar(255);comment:失败原因/拦截原因"`
	CreatedAt  *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:记录创建时间"`
}

func (UserLoginLog) TableName() string {
	return "app_user_login_log"
}
