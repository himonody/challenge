package models

import "time"

// MessageUser 用户收件箱
type MessageUser struct {
	ID        uint64     `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	MessageID uint64     `json:"messageId" gorm:"column:message_id;type:bigint;not null;uniqueIndex:uk_user_message,priority:2;comment:消息ID"`
	UserID    uint64     `json:"userId" gorm:"column:user_id;type:bigint;not null;uniqueIndex:uk_user_message,priority:1;index:idx_user_read,priority:1;index:idx_user_time,priority:1;index:idx_user_deleted,priority:1;comment:接收用户ID"`
	IsRead    int8       `json:"isRead" gorm:"column:is_read;type:tinyint;not null;default:0;index:idx_user_read,priority:2;comment:是否已读 0未读 1已读"`
	IsDeleted int8       `json:"isDeleted" gorm:"column:is_deleted;type:tinyint;not null;default:0;index:idx_user_deleted,priority:2;comment:是否删除 0否 1是"`
	ReadAt    *time.Time `json:"readAt" gorm:"column:read_at;type:datetime;comment:阅读时间"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP;index:idx_user_time,priority:2;comment:接收时间"`
}

// TableName 指定表名
func (MessageUser) TableName() string {
	return "app_message_user"
}
