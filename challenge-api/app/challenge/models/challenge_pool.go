package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type ChallengePool struct {
	Id          int64           `json:"id" gorm:"primaryKey;autoIncrement;comment:奖池ID"`
	ConfigId    int64           `json:"configId" gorm:"column:config_id;type:bigint;comment:活动配置ID"`
	StartDate   *time.Time      `json:"startDate" gorm:"column:start_date;type:datetime;comment:活动开始日期"`
	EndDate     *time.Time      `json:"endDate" gorm:"column:end_date;type:datetime;comment:活动结束日期"`
	TotalAmount decimal.Decimal `json:"totalAmount" gorm:"column:total_amount;type:decimal(30,2);comment:奖池当前总金额"`
	Settled     int64           `json:"settled" gorm:"column:settled;type:tinyint;comment:是否已结算 0否 1是"`
	CreatedAt   *time.Time      `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间戳"`

	Config *ChallengeConfig `json:"config" gorm:"foreignkey:config_id"`
}

func (ChallengePool) TableName() string {
	return "app_challenge_pool"
}
