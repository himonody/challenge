package models

import "time"

type ChallengePoolFlow struct {
	Id        uint64     `json:"id" gorm:"column:id;primaryKey;autoIncrement;comment:奖池流水ID"`
	PoolId    uint64     `json:"poolId" gorm:"column:pool_id;type:bigint;comment:奖池ID"`
	UserId    uint64     `json:"userId" gorm:"column:user_id;type:bigint;comment:用户ID"`
	Amount    string     `json:"amount" gorm:"column:amount;type:decimal(30,2);comment:变动金额"`
	Type      int8       `json:"type" gorm:"column:type;type:tinyint;comment:类型"`
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
}

func (ChallengePoolFlow) TableName() string { return "app_challenge_pool_flow" }
