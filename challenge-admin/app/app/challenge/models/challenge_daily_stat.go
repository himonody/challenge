package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type ChallengeDailyStat struct {
	StatDate       time.Time       `json:"statDate" gorm:"column:stat_date;type:date;primaryKey;comment:统计日期"`
	JoinUserCnt    int             `json:"joinUserCnt" gorm:"column:join_user_cnt;type:int;comment:参与人数"`
	SuccessUserCnt int             `json:"successUserCnt" gorm:"column:success_user_cnt;type:int;comment:成功人数"`
	FailUserCnt    int             `json:"failUserCnt" gorm:"column:fail_user_cnt;type:int;comment:失败人数"`
	JoinAmount     decimal.Decimal `json:"joinAmount" gorm:"column:join_amount;type:decimal(30,2);comment:参与总金额"`
	SuccessAmount  decimal.Decimal `json:"successAmount" gorm:"column:success_amount;type:decimal(30,2);comment:成功金额"`
	FailAmount     decimal.Decimal `json:"failAmount" gorm:"column:fail_amount;type:decimal(30,2);comment:失败金额"`
	PlatformBonus  decimal.Decimal `json:"platformBonus" gorm:"column:platform_bonus;type:decimal(30,2);comment:平台补贴"`
	PoolAmount     decimal.Decimal `json:"poolAmount" gorm:"column:pool_amount;type:decimal(30,2);comment:奖池金额"`
	CreatedAt      *time.Time      `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
	UpdatedAt      *time.Time      `json:"updatedAt" gorm:"column:updated_at;type:datetime;comment:更新时间"`
}

func (ChallengeDailyStat) TableName() string { return "app_challenge_daily_stat" }
