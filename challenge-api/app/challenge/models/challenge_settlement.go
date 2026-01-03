package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type ChallengeSettlement struct {
	Id          int64           `json:"id" gorm:"primaryKey;autoIncrement;comment:结算ID"`
	ChallengeId int64           `json:"challengeId" gorm:"column:challenge_id;type:bigint;comment:用户挑战ID"`
	UserId      int64           `json:"userId" gorm:"column:user_id;type:bigint;comment:用户ID"`
	Reward      decimal.Decimal `json:"reward" gorm:"column:reward;type:decimal(30,2);comment:最终获得金额"`
	CreatedAt   *time.Time      `json:"createdAt" gorm:"column:created_at;type:datetime;comment:结算时间戳"`

	ChallengeUser *ChallengeUser `json:"challengeUser" gorm:"foreignkey:challenge_id"`
}

func (ChallengeSettlement) TableName() string {
	return "app_challenge_settlement"
}
