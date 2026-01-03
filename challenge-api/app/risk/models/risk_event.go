package models

import (
	"time"
)

type RiskEvent struct {
	Id        int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:事件ID"`
	UserId    int64      `json:"userId" gorm:"column:user_id;type:bigint;comment:用户ID"`
	EventType int64      `json:"eventType" gorm:"column:event_type;type:tinyint;comment:事件类型"`
	Detail    string     `json:"detail" gorm:"column:detail;type:varchar(255);comment:事件详情"`
	Score     int64      `json:"score" gorm:"column:score;type:int;comment:风险分"`
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:发生时间"`
}

func (RiskEvent) TableName() string {
	return "app_risk_event"
}
