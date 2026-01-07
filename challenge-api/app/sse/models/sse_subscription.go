package models

import "time"

// AppSSESubscription SSE 订阅关系表
type AppSSESubscription struct {
	ID         uint64    `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement;comment:主键ID" json:"id"`
	UserID     string    `gorm:"column:user_id;type:varchar(64);not null;uniqueIndex:idx_user_group;comment:用户ID" json:"user_id"`
	GroupName  string    `gorm:"column:group_name;type:varchar(64);not null;uniqueIndex:idx_user_group;index:idx_group_status;comment:订阅组名" json:"group_name"`
	EventTypes string    `gorm:"column:event_types;type:varchar(255);default:'';comment:订阅的事件类型（逗号分隔）" json:"event_types"`
	Status     int8      `gorm:"column:status;type:tinyint;default:1;index:idx_group_status,priority:2;comment:状态（0-禁用 1-启用）" json:"status"`
	CreatedAt  time.Time `gorm:"column:created_at;type:datetime;not null;comment:创建时间" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:datetime;comment:更新时间" json:"updated_at"`
}

// TableName 表名
func (AppSSESubscription) TableName() string {
	return "app_sse_subscription"
}

// 订阅状态常量
const (
	SubscriptionStatusDisabled = 0 // 禁用
	SubscriptionStatusEnabled  = 1 // 启用
)
