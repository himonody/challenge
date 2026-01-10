package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type AppChallengeUser struct {
	ID              uint64          `gorm:"primaryKey;autoIncrement;comment:用户挑战ID"`
	UserID          uint64          `gorm:"not null;uniqueIndex:uk_user_active,priority:1;comment:用户ID"`
	ConfigID        uint64          `gorm:"not null;comment:活动配置ID"`
	PoolID          uint64          `gorm:"not null;index:idx_pool;comment:奖池ID"`
	ChallengeAmount decimal.Decimal `gorm:"type:decimal(30,2);not null;default:0.00"`
	StartDate       int             `gorm:"not null;default:0;comment:开始日期"`
	EndDate         int             `gorm:"not null;default:0;comment:结束日期"`
	Status          uint8           `gorm:"not null;default:1;uniqueIndex:uk_user_active,priority:2;index:idx_status"`
	FailReason      uint8           `gorm:"not null;default:0"`
	CreatedAt       time.Time       `gorm:"autoCreateTime;comment:报名时间"`
	FinishedAt      *time.Time      `gorm:"comment:完成时间"`
}

func (AppChallengeUser) TableName() string {
	return "app_challenge_user"
}
