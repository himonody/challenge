package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type ChallengeConfig struct {
	Id            int64           `json:"id" gorm:"primaryKey;autoIncrement;comment:活动配置ID"`
	DayCount      int64           `json:"dayCount" gorm:"column:day_count;type:int;comment:挑战天数 1/7/21"`
	Amount        decimal.Decimal `json:"amount" gorm:"column:amount;type:decimal(30,2);comment:单人挑战金额"`
	CheckinStart  *time.Time      `json:"checkinStart" gorm:"column:checkin_start;type:datetime;comment:每日打卡开始时间"`
	CheckinEnd    *time.Time      `json:"checkinEnd" gorm:"column:checkin_end;type:datetime;comment:每日打卡结束时间"`
	PlatformBonus decimal.Decimal `json:"platformBonus" gorm:"column:platform_bonus;type:decimal(30,2);comment:平台补贴金额"`
	Status        int64           `json:"status" gorm:"column:status;type:tinyint;comment:状态 1启用 2停用"`
	CreatedAt     *time.Time      `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间戳"`
}

func (ChallengeConfig) TableName() string {
	return "app_challenge_config"
}
