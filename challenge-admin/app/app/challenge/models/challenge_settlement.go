package models

import "time"

type ChallengeSettlement struct {
	Id          uint64     `json:"id" gorm:"column:id;primaryKey;autoIncrement;comment:结算ID"`
	ChallengeId uint64     `json:"challengeId" gorm:"column:challenge_id;type:bigint;comment:用户挑战ID"`
	UserId      uint64     `json:"userId" gorm:"column:user_id;type:bigint;comment:用户ID"`
	Reward      string     `json:"reward" gorm:"column:reward;type:decimal(30,2);comment:最终获得金额"`
	CreatedAt   *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:结算时间"`
}

func (ChallengeSettlement) TableName() string { return "app_challenge_settlement" }
