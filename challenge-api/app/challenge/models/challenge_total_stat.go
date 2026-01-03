package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type ChallengeTotalStat struct {
	Id                 int64           `json:"id" gorm:"primaryKey;column:id;type:tinyint;comment:固定主键"`
	TotalUserCnt       int64           `json:"totalUserCnt" gorm:"column:total_user_cnt;type:int;comment:累计用户数"`
	TotalJoinCnt       int64           `json:"totalJoinCnt" gorm:"column:total_join_cnt;type:int;comment:累计参与人次"`
	TotalSuccessCnt    int64           `json:"totalSuccessCnt" gorm:"column:total_success_cnt;type:int;comment:累计成功人次"`
	TotalFailCnt       int64           `json:"totalFailCnt" gorm:"column:total_fail_cnt;type:int;comment:累计失败人次"`
	TotalJoinAmount    decimal.Decimal `json:"totalJoinAmount" gorm:"column:total_join_amount;type:decimal(30,2);comment:累计参与金额"`
	TotalSuccessAmount decimal.Decimal `json:"totalSuccessAmount" gorm:"column:total_success_amount;type:decimal(30,2);comment:累计成功金额"`
	TotalFailAmount    decimal.Decimal `json:"totalFailAmount" gorm:"column:total_fail_amount;type:decimal(30,2);comment:累计失败金额"`
	TotalPlatformBonus decimal.Decimal `json:"totalPlatformBonus" gorm:"column:total_platform_bonus;type:decimal(30,2);comment:累计平台补贴"`
	TotalPoolAmount    decimal.Decimal `json:"totalPoolAmount" gorm:"column:total_pool_amount;type:decimal(30,2);comment:累计奖池金额"`
	UpdatedAt          *time.Time      `json:"updatedAt" gorm:"column:updated_at;type:datetime;comment:更新时间戳"`
}

func (ChallengeTotalStat) TableName() string {
	return "app_challenge_total_stat"
}
