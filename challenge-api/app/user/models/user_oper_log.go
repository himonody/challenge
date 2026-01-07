package models

import (
	"time"
)

// AppUserOperLog 用户关键行为日志
type AppUserOperLog struct {
	ID int `gorm:"column:id;primaryKey;autoIncrement;comment:日志编码" json:"id"`

	UserID int `gorm:"column:user_id;not null;default:0;index:idx_user_created,priority:1;comment:用户编号" json:"user_id"`

	ActionType string `gorm:"column:action_type;type:char(2);not null;default:'';index:idx_action_type;comment:用户行为类型" json:"action_type"`
	OperateIP  string `gorm:"column:operate_ip;type:varchar(45);comment:操作IP" json:"operate_ip"`

	ByType string `gorm:"column:by_type;type:char(2);not null;default:'1';comment:更新用户类型 1-app用户 2-后台用户" json:"by_type"`

	Status string `gorm:"column:status;type:char(1);not null;default:'1';index:idx_status;comment:状态(1-正常 2-异常)" json:"status"`

	CreateBy int `gorm:"column:create_by;not null;default:0;comment:创建者" json:"create_by"`
	UpdateBy int `gorm:"column:update_by;not null;default:0;comment:更新者" json:"update_by"`

	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP;index:idx_user_created,priority:2;comment:创建时间" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`

	Remark string `gorm:"column:remark;type:varchar(255);comment:备注信息" json:"remark"`
}

// TableName 指定表名
func (AppUserOperLog) TableName() string {
	return "app_user_oper_log"
}
