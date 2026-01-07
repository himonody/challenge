package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type ChallengeRankDaily struct {
	Id       int64           `json:"id" gorm:"primaryKey;autoIncrement;comment:排行ID"`
	RankDate *time.Time      `json:"rankDate" gorm:"column:rank_date;type:date;uniqueIndex:uk_rank,priority:1;index:idx_rank_type,priority:2;comment:排行日期"`
	RankType int64           `json:"rankType" gorm:"column:rank_type;type:tinyint;uniqueIndex:uk_rank,priority:2;index:idx_rank_type,priority:1;comment:1邀请 2收益 3毅力"`
	UserId   int64           `json:"userId" gorm:"column:user_id;type:bigint;uniqueIndex:uk_rank,priority:3;comment:用户ID"`
	Value    decimal.Decimal `json:"value" gorm:"column:value;type:decimal(30,2);comment:排行值"`
	RankNo   int64           `json:"rankNo" gorm:"column:rank_no;type:int;comment:排名"`
}

func (ChallengeRankDaily) TableName() string {
	return "app_challenge_rank_daily"
}
