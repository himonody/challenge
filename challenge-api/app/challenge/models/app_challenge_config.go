package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type AppChallengeConfig struct {
	ID            uint64          `gorm:"primaryKey;autoIncrement;comment:活动配置ID"`
	DayCount      int             `gorm:"not null;default:0;uniqueIndex:uk_day_amount,priority:1;comment:挑战天数"`
	Amount        decimal.Decimal `gorm:"type:decimal(30,2);not null;default:0.00;uniqueIndex:uk_day_amount,priority:2;comment:单人挑战金额"`
	CheckinStart  *time.Time      `gorm:"comment:每日打卡开始时间"`
	CheckinEnd    *time.Time      `gorm:"comment:每日打卡结束时间"`
	PlatformBonus decimal.Decimal `gorm:"type:decimal(30,2);not null;default:0.00;comment:平台补贴金额"`
	Status        int             `gorm:"not null;default:1;index:idx_status;comment:状态"`
	Sort          int             `gorm:"not null;default:1;index:idx_sort;comment:排序"`
	CreatedAt     time.Time       `gorm:"autoCreateTime;comment:创建时间"`
}

func (AppChallengeConfig) TableName() string {
	return "app_challenge_config"
}
