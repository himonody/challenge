package models

import (
	"time"

	"gorm.io/datatypes"
)

// Message 站内信消息主体
type Message struct {
	ID         uint64         `json:"id" gorm:"primaryKey;autoIncrement;comment:消息ID"`
	MsgType    int8           `json:"msgType" gorm:"column:msg_type;type:tinyint;not null;default:1;index:idx_type_time,priority:1;comment:消息类型 1系统通知 2站内信 3私信"`
	Title      string         `json:"title" gorm:"column:title;type:varchar(100);not null;default:'';comment:消息标题"`
	Content    string         `json:"content" gorm:"column:content;type:text;not null;comment:消息内容"`
	SenderID   uint64         `json:"senderId" gorm:"column:sender_id;type:bigint;not null;default:0;index:idx_sender,priority:2;comment:发送者ID"`
	SenderType int8           `json:"senderType" gorm:"column:sender_type;type:tinyint;not null;default:0;index:idx_sender,priority:1;comment:发送者类型 0系统 1用户 2管理员"`
	SenderName string         `json:"senderName" gorm:"column:sender_name;type:varchar(50);not null;default:'';comment:发送者名称"`
	BizType    string         `json:"bizType" gorm:"column:biz_type;type:varchar(50);not null;default:'';index:idx_biz,priority:1;comment:业务类型"`
	BizID      uint64         `json:"bizId" gorm:"column:biz_id;type:bigint;not null;default:0;index:idx_biz,priority:2;comment:业务ID"`
	Extra      datatypes.JSON `json:"extra" gorm:"column:extra;type:json;comment:扩展数据"`
	CreatedAt  time.Time      `json:"createdAt" gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP;index:idx_type_time,priority:2;comment:创建时间"`
}

// TableName 指定表名
func (Message) TableName() string {
	return "app_message"
}
