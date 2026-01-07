package models

import "time"

// UserInviteCode 邀请码
type UserInviteCode struct {
	ID          uint64     `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	Code        string     `json:"code" gorm:"column:code;type:varchar(64);not null;uniqueIndex:uk_code;comment:邀请码"`
	OwnerUserID uint64     `json:"ownerUserId" gorm:"column:owner_user_id;type:bigint;not null;index:idx_owner;index:idx_owner_status,priority:1;comment:所属用户"`
	Status      string     `json:"status" gorm:"column:status;type:char(1);not null;default:'1';index:idx_status;index:idx_owner_status,priority:2;comment:1可用 2禁用"`
	TotalLimit  int        `json:"totalLimit" gorm:"column:total_limit;type:int;not null;default:0;comment:总次数 0不限制"`
	DailyLimit  int        `json:"dailyLimit" gorm:"column:daily_limit;type:int;not null;default:0;comment:每日次数 0不限制"`
	UsedTotal   int        `json:"usedTotal" gorm:"column:used_total;type:int;not null;default:0;comment:已使用总次数"`
	UsedToday   int        `json:"usedToday" gorm:"column:used_today;type:int;not null;default:0;index:idx_used_today;comment:今日已使用次数"`
	LastUsedAt  *time.Time `json:"lastUsedAt" gorm:"column:last_used_at;type:datetime;comment:最后使用时间"`
	CreatedAt   *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
}

// TableName 指定表名
func (UserInviteCode) TableName() string {
	return "app_user_invite_code"
}
