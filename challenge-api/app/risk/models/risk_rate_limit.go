package models

import "time"

// RiskRateLimit 频率限制记录
type RiskRateLimit struct {
	ID            uint64     `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	Scene         string     `json:"scene" gorm:"column:scene;type:varchar(32);not null;index:idx_scene_identity,priority:1;uniqueIndex:uk_scene_identity_window,priority:1;comment:场景"`
	IdentityType  string     `json:"identityType" gorm:"column:identity_type;type:varchar(16);not null;index:idx_scene_identity,priority:2;uniqueIndex:uk_scene_identity_window,priority:2;comment:标识类型"`
	IdentityValue string     `json:"identityValue" gorm:"column:identity_value;type:varchar(128);not null;index:idx_scene_identity,priority:3;uniqueIndex:uk_scene_identity_window,priority:3;comment:标识值"`
	Count         int        `json:"count" gorm:"column:count;type:int;not null;default:1;comment:次数"`
	WindowStart   time.Time  `json:"windowStart" gorm:"column:window_start;type:datetime;not null;uniqueIndex:uk_scene_identity_window,priority:4;comment:窗口开始时间"`
	WindowEnd     time.Time  `json:"windowEnd" gorm:"column:window_end;type:datetime;not null;comment:窗口结束时间"`
	Blocked       string     `json:"blocked" gorm:"column:blocked;type:char(1);not null;default:'0';index:idx_blocked;comment:是否已拦截 0否 1是"`
	CreatedAt     *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
}

// TableName 指定表名
func (RiskRateLimit) TableName() string {
	return "app_risk_rate_limit"
}
