package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type AppChallengeTotalStat struct {
	ID int `gorm:"primaryKey;default:1;comment:固定主键"`

	TotalUserCnt    int `gorm:"not null;default:0;comment:累计用户数"`
	TotalJoinCnt    int `gorm:"not null;default:0;comment:累计参与人次"`
	TotalSuccessCnt int `gorm:"not null;default:0;comment:累计成功人次"`
	TotalFailCnt    int `gorm:"not null;default:0;comment:累计失败人次"`

	TotalJoinAmount    decimal.Decimal `gorm:"type:decimal(30,2);not null;default:0.00;comment:累计参与金额"`
	TotalSuccessAmount decimal.Decimal `gorm:"type:decimal(30,2);not null;default:0.00;comment:累计成功金额"`
	TotalFailAmount    decimal.Decimal `gorm:"type:decimal(30,2);not null;default:0.00;comment:累计失败金额"`

	TotalPlatformBonus decimal.Decimal `gorm:"type:decimal(30,2);not null;default:0.00;comment:累计平台补贴"`
	TotalPoolAmount    decimal.Decimal `gorm:"type:decimal(30,2);not null;default:0.00;comment:累计奖池金额"`

	UpdatedAt time.Time `gorm:"autoUpdateTime;comment:更新时间"`
}

func (AppChallengeTotalStat) TableName() string {
	return "app_challenge_total_stat"
}
