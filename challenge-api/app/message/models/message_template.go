package models

import "time"

// MessageTemplate 消息模板
type MessageTemplate struct {
	ID         uint64    `json:"id" gorm:"primaryKey;autoIncrement;comment:模板ID"`
	Code       string    `json:"code" gorm:"column:code;type:varchar(50);not null;uniqueIndex:uk_code;comment:模板编码"`
	MsgType    int8      `json:"msgType" gorm:"column:msg_type;type:tinyint;not null;default:1;index:idx_type_status,priority:1;comment:消息类型"`
	TitleTpl   string    `json:"titleTpl" gorm:"column:title_tpl;type:varchar(100);not null;comment:标题模板"`
	ContentTpl string    `json:"contentTpl" gorm:"column:content_tpl;type:text;not null;comment:内容模板"`
	SenderType int8      `json:"senderType" gorm:"column:sender_type;type:tinyint;not null;default:0;comment:发送者类型"`
	SenderID   uint64    `json:"senderId" gorm:"column:sender_id;type:bigint;not null;default:0;comment:发送者ID"`
	SenderName string    `json:"senderName" gorm:"column:sender_name;type:varchar(50);not null;default:'';comment:发送者名称"`
	BizType    string    `json:"bizType" gorm:"column:biz_type;type:varchar(50);not null;default:'';comment:业务类型"`
	Status     int8      `json:"status" gorm:"column:status;type:tinyint;not null;default:1;index:idx_type_status,priority:2;comment:状态 1启用 0禁用"`
	CreatedAt  time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
}

// TableName 指定表名
func (MessageTemplate) TableName() string {
	return "app_message_template"
}
