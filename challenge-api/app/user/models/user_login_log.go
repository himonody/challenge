package models

import (
	"time"
)

// UserLoginLog 用户登录日志
type UserLoginLog struct {
	ID         uint64     `json:"id" gorm:"primaryKey;autoIncrement;comment:登录日志ID"`
	UserID     uint64     `json:"userId" gorm:"column:user_id;type:bigint;not null;default:0;index:idx_user_time,priority:1;comment:用户ID"`
	LoginAt    time.Time  `json:"loginAt" gorm:"column:login_at;type:datetime;not null;default:CURRENT_TIMESTAMP;index:idx_user_time,priority:2;comment:登录时间"`
	LoginIP    string     `json:"loginIp" gorm:"column:login_ip;type:varchar(45);index:idx_ip;comment:登录IP"`
	DeviceFP   string     `json:"deviceFp" gorm:"column:device_fp;type:varchar(64);comment:设备指纹"`
	UserAgent  string     `json:"userAgent" gorm:"column:user_agent;type:varchar(500);comment:UA信息"`
	Status     int8       `json:"status" gorm:"column:status;type:tinyint;not null;default:1;index:idx_status;comment:登录状态 1成功 2失败 3风控拦截"`
	FailReason string     `json:"failReason" gorm:"column:fail_reason;type:varchar(255);comment:失败原因/拦截原因"`
	CreatedAt  *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:记录创建时间戳"`
}

// TableName 指定表名
func (UserLoginLog) TableName() string {
	return "app_user_login_log"
}
