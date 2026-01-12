package models

import (
	"time"
)

type ChallengePool struct {
	Id          uint64     `json:"id" gorm:"column:id;primaryKey;autoIncrement;comment:奖池ID"`
	ConfigId    uint64     `json:"configId" gorm:"column:config_id;type:bigint;comment:活动配置ID"`
	StartDate   *time.Time `json:"startDate" gorm:"column:start_date;type:datetime;comment:开始日期"`
	EndDate     *time.Time `json:"endDate" gorm:"column:end_date;type:datetime;comment:结束日期"`
	TotalAmount string     `json:"totalAmount" gorm:"column:total_amount;type:decimal(30,2);comment:奖池金额"`
	Settled     int8       `json:"settled" gorm:"column:settled;type:tinyint;comment:是否已结算"`
	CreatedAt   *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
}

func (ChallengePool) TableName() string { return "app_challenge_pool" }
