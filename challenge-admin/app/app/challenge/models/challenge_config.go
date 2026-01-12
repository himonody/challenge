package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type ChallengeConfig struct {
	Id            uint64          `json:"id" gorm:"column:id;primaryKey;autoIncrement;comment:活动配置ID"`
	DayCount      int             `json:"dayCount" gorm:"column:day_count;type:int;comment:挑战天数"`
	Amount        decimal.Decimal `json:"amount" gorm:"column:amount;type:decimal(30,2);comment:单人挑战金额"`
	CheckinStart  *time.Time      `json:"checkinStart" gorm:"column:checkin_start;type:datetime;comment:每日打卡开始时间"`
	CheckinEnd    *time.Time      `json:"checkinEnd" gorm:"column:checkin_end;type:datetime;comment:每日打卡结束时间"`
	PlatformBonus decimal.Decimal `json:"platformBonus" gorm:"column:platform_bonus;type:decimal(30,2);comment:平台补贴金额"`
	Status        int8            `json:"status" gorm:"column:status;type:tinyint;comment:状态"`
	Sort          int8            `json:"sort" gorm:"column:sort;type:tinyint;comment:排序"`
	CreatedAt     *time.Time      `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
}

func (ChallengeConfig) TableName() string { return "app_challenge_config" }
