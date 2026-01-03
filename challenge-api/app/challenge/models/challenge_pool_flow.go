package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type ChallengePoolFlow struct {
	Id        int64           `json:"id" gorm:"primaryKey;autoIncrement;comment:奖池流水ID"`
	PoolId    int64           `json:"poolId" gorm:"column:pool_id;type:bigint;comment:奖池ID"`
	UserId    int64           `json:"userId" gorm:"column:user_id;type:bigint;comment:用户ID"`
	Amount    decimal.Decimal `json:"amount" gorm:"column:amount;type:decimal(30,2);comment:变动金额"`
	Type      int64           `json:"type" gorm:"column:type;type:tinyint;comment:类型 1报名 2失败 3平台补贴 4结算"`
	CreatedAt *time.Time      `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间戳"`

	Pool *ChallengePool `json:"pool" gorm:"foreignkey:pool_id"`
}

func (ChallengePoolFlow) TableName() string {
	return "app_challenge_pool_flow"
}
