package models

import (
	"time"

	"github.com/shopspring/decimal"
)

// AppUserInviteRelation 用户邀请关系表
type AppUserInviteRelation struct {
	ID            uint64          `gorm:"primaryKey;autoIncrement;column:id;comment:邀请关系ID" json:"id"`
	InviterUserID uint64          `gorm:"column:inviter_user_id;type:bigint unsigned;not null;comment:邀请人用户ID;index:idx_inviter" json:"inviterUserId"`
	InviteeUserID uint64          `gorm:"column:invitee_user_id;type:bigint unsigned;not null;uniqueIndex:uk_invitee;comment:被邀请人用户ID" json:"inviteeUserId"`
	InviteCode    string          `gorm:"column:invite_code;type:varchar(64);not null;index:idx_invite_code;comment:使用的邀请码" json:"inviteCode"`
	InviteReward  decimal.Decimal `gorm:"column:invite_reward;type:decimal(30,2);not null;default:0.00;comment:邀请奖励" json:"inviteReward"`
	InviteeReward decimal.Decimal `gorm:"column:invitee_reward;type:decimal(30,2);not null;default:0.00;comment:被邀请人奖励" json:"inviteeReward"`
	Status        string          `gorm:"column:status;type:char(1);not null;default:'1';comment:状态 1有效 2无效" json:"status"`
	CreatedAt     time.Time       `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP;index:idx_created_at;comment:邀请时间" json:"createdAt"`
}

// TableName 指定表名
func (AppUserInviteRelation) TableName() string {
	return "app_user_invite_relation"
}
