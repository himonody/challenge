package models

import "time"

type RiskRateLimit struct {
	Id            int64      `json:"id" gorm:"primaryKey;autoIncrement"`
	Scene         string     `json:"scene" gorm:"column:scene;type:varchar(32);comment:场景"`
	IdentityType  string     `json:"identityType" gorm:"column:identity_type;type:varchar(16);comment:标识类型"`
	IdentityValue string     `json:"identityValue" gorm:"column:identity_value;type:varchar(128);comment:标识值"`
	Count         int        `json:"count" gorm:"column:count;type:int;comment:次数"`
	WindowStart   *time.Time `json:"windowStart" gorm:"column:window_start;type:datetime;comment:窗口开始时间"`
	WindowEnd     *time.Time `json:"windowEnd" gorm:"column:window_end;type:datetime;comment:窗口结束时间"`
	Blocked       string     `json:"blocked" gorm:"column:blocked;type:char(1);comment:是否已拦截"`
	CreatedAt     *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:记录时间"`
}

func (RiskRateLimit) TableName() string { return "app_risk_rate_limit" }
