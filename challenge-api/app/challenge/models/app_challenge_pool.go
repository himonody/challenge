package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type AppChallengePool struct {
	ID          uint64     `gorm:"primaryKey;autoIncrement;comment:奖池ID"`
	ConfigID    uint64     `gorm:"not null;uniqueIndex:uk_config_date,priority:1"`
	StartDate   *time.Time `gorm:"uniqueIndex:uk_config_date,priority:2"`
	EndDate     *time.Time
	TotalAmount decimal.Decimal `gorm:"type:decimal(30,2);not null;default:0.00"`
	Settled     int             `gorm:"not null;default:0;index:idx_settled"`
	CreatedAt   time.Time       `gorm:"autoCreateTime;comment:创建时间"`
}

func (AppChallengePool) TableName() string {
	return "app_challenge_pool"
}
