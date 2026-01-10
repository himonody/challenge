package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type AppChallengeDailyStat struct {
	StatDate time.Time `gorm:"type:date;primaryKey;comment:统计日期 YYYYMMDD"`

	JoinUserCnt    int `gorm:"not null;default:0;comment:参与人数"`
	SuccessUserCnt int `gorm:"not null;default:0;comment:成功人数"`
	FailUserCnt    int `gorm:"not null;default:0;comment:失败人数"`

	JoinAmount    decimal.Decimal `gorm:"type:decimal(30,2);not null;default:0.00;comment:参与总金额"`
	SuccessAmount decimal.Decimal `gorm:"type:decimal(30,2);not null;default:0.00;comment:成功金额"`
	FailAmount    decimal.Decimal `gorm:"type:decimal(30,2);not null;default:0.00;comment:失败金额"`

	PlatformBonus decimal.Decimal `gorm:"type:decimal(30,2);not null;default:0.00;comment:平台补贴"`
	PoolAmount    decimal.Decimal `gorm:"type:decimal(30,2);not null;default:0.00;comment:奖池金额"`

	CreatedAt time.Time `gorm:"autoCreateTime;comment:创建时间"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;comment:更新时间"`
}

func (AppChallengeDailyStat) TableName() string {
	return "app_challenge_daily_stat"
}
