package models

import "time"

// AppSSEMessage SSE 消息记录表（用于重连恢复）
type AppSSEMessage struct {
	ID           uint64     `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement;comment:主键ID" json:"id"`
	EventID      string     `gorm:"column:event_id;type:varchar(64);not null;uniqueIndex:idx_event_id;comment:事件ID" json:"event_id"`
	EventType    string     `gorm:"column:event_type;type:varchar(64);not null;index:idx_event_type;comment:事件类型" json:"event_type"`
	ReceiverID   string     `gorm:"column:receiver_id;type:varchar(64);not null;index:idx_receiver_created;comment:接收者ID（用户ID或客户端ID）" json:"receiver_id"`
	ReceiverType string     `gorm:"column:receiver_type;type:varchar(20);not null;comment:接收者类型（user/client/group）" json:"receiver_type"`
	GroupName    string     `gorm:"column:group_name;type:varchar(64);default:'';index:idx_group_created;comment:分组名称" json:"group_name"`
	Priority     int        `gorm:"column:priority;type:tinyint;default:0;comment:优先级（0-普通 1-高 2-紧急）" json:"priority"`
	Data         string     `gorm:"column:data;type:text;comment:消息数据（JSON）" json:"data"`
	Status       int8       `gorm:"column:status;type:tinyint;default:0;comment:状态（0-待发送 1-已发送 2-已读）" json:"status"`
	ExpireAt     *time.Time `gorm:"column:expire_at;type:datetime;index;comment:过期时间" json:"expire_at"`
	CreatedAt    time.Time  `gorm:"column:created_at;type:datetime;not null;index:idx_receiver_created,priority:2;index:idx_group_created,priority:2;comment:创建时间" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"column:updated_at;type:datetime;comment:更新时间" json:"updated_at"`
}

// TableName 表名
func (AppSSEMessage) TableName() string {
	return "app_sse_message"
}

// 消息状态常量
const (
	MessageStatusPending = 0 // 待发送
	MessageStatusSent    = 1 // 已发送
	MessageStatusRead    = 2 // 已读
)

// 接收者类型常量
const (
	ReceiverTypeUser   = "user"   // 用户
	ReceiverTypeClient = "client" // 客户端
	ReceiverTypeGroup  = "group"  // 分组
)

// 消息优先级常量
const (
	PriorityNormal = 0 // 普通
	PriorityHigh   = 1 // 高
	PriorityUrgent = 2 // 紧急
)
