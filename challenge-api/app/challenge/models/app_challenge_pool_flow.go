package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type AppChallengePoolFlow struct {
	ID        uint64          `gorm:"primaryKey;autoIncrement;comment:奖池流水ID"`
	PoolID    uint64          `gorm:"not null;index:idx_pool"`
	UserID    uint64          `gorm:"not null;index:idx_user"`
	Amount    decimal.Decimal `gorm:"type:decimal(30,2);not null;default:0.00"`
	Type      uint8           `gorm:"not null;default:0;index:idx_type_time,priority:1"`
	CreatedAt time.Time       `gorm:"autoCreateTime;index:idx_type_time,priority:2;comment:创建时间"`
}

func (AppChallengePoolFlow) TableName() string {
	return "app_challenge_pool_flow"
}
