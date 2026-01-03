package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type ChallengeUser struct {
	Id              int64           `json:"id" gorm:"primaryKey;autoIncrement;comment:用户挑战ID"`
	UserId          int64           `json:"userId" gorm:"column:user_id;type:bigint;comment:用户ID"`
	ConfigId        int64           `json:"configId" gorm:"column:config_id;type:bigint;comment:活动配置ID"`
	PoolId          int64           `json:"poolId" gorm:"column:pool_id;type:bigint;comment:奖池ID"`
	ChallengeAmount decimal.Decimal `json:"challengeAmount" gorm:"column:challenge_amount;type:decimal(30,2);comment:用户挑战金额"`
	StartDate       int64           `json:"startDate" gorm:"column:start_date;type:int;comment:活动开始日期 YYYYMMDD"`
	EndDate         int64           `json:"endDate" gorm:"column:end_date;type:int;comment:活动结束日期 YYYYMMDD"`
	Status          int64           `json:"status" gorm:"column:status;type:tinyint;comment:状态 1进行中 2成功 3失败"`
	FailReason      int64           `json:"failReason" gorm:"column:fail_reason;type:tinyint;comment:失败原因 0无 1未打卡 2作弊"`
	CreatedAt       *time.Time      `json:"createdAt" gorm:"column:created_at;type:datetime;comment:报名时间戳"`
	FinishedAt      *time.Time      `json:"finishedAt" gorm:"column:finished_at;type:datetime;comment:完成时间戳"`

	Config   *ChallengeConfig   `json:"config" gorm:"foreignkey:config_id"`
	Pool     *ChallengePool     `json:"pool" gorm:"foreignkey:pool_id"`
	Checkins []ChallengeCheckin `json:"checkins" gorm:"foreignkey:challenge_id"`
}

func (ChallengeUser) TableName() string {
	return "app_challenge_user"
}
