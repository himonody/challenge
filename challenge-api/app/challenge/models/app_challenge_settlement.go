package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type AppChallengeSettlement struct {
	ID          uint64          `gorm:"primaryKey;autoIncrement;comment:结算ID"`
	ChallengeID uint64          `gorm:"not null;uniqueIndex:uk_challenge_user,priority:1"`
	UserID      uint64          `gorm:"not null;uniqueIndex:uk_challenge_user,priority:2;index:idx_user"`
	Reward      decimal.Decimal `gorm:"type:decimal(30,2);not null;default:0.00"`
	CreatedAt   time.Time       `gorm:"autoCreateTime;comment:结算时间"`
}

func (AppChallengeSettlement) TableName() string {
	return "app_challenge_settlement"
}
