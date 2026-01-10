package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type AppChallengeRankDaily struct {
	ID       uint64          `gorm:"primaryKey;autoIncrement;comment:排行ID"`
	RankDate time.Time       `gorm:"type:date;not null;uniqueIndex:uk_rank,priority:1"`
	RankType uint8           `gorm:"not null;uniqueIndex:uk_rank,priority:2;index:idx_rank_type,priority:1"`
	UserID   uint64          `gorm:"not null;uniqueIndex:uk_rank,priority:3"`
	Value    decimal.Decimal `gorm:"type:decimal(30,2);not null"`
	RankNo   int
}

func (AppChallengeRankDaily) TableName() string {
	return "app_challenge_rank_daily"
}
