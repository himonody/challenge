package models

import "time"

type ChallengeUser struct {
	Id              uint64     `json:"id" gorm:"column:id;primaryKey;autoIncrement;comment:用户挑战ID"`
	UserId          uint64     `json:"userId" gorm:"column:user_id;type:bigint;comment:用户ID"`
	ConfigId        uint64     `json:"configId" gorm:"column:config_id;type:bigint;comment:活动配置ID"`
	PoolId          uint64     `json:"poolId" gorm:"column:pool_id;type:bigint;comment:奖池ID"`
	ChallengeAmount string     `json:"challengeAmount" gorm:"column:challenge_amount;type:decimal(30,2);comment:用户挑战金额"`
	StartDate       int        `json:"startDate" gorm:"column:start_date;type:int;comment:开始日期"`
	EndDate         int        `json:"endDate" gorm:"column:end_date;type:int;comment:结束日期"`
	Status          int8       `json:"status" gorm:"column:status;type:tinyint;comment:状态"`
	FailReason      int8       `json:"failReason" gorm:"column:fail_reason;type:tinyint;comment:失败原因"`
	CreatedAt       *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:报名时间"`
	FinishedAt      *time.Time `json:"finishedAt" gorm:"column:finished_at;type:datetime;comment:完成时间"`
}

func (ChallengeUser) TableName() string { return "app_challenge_user" }
