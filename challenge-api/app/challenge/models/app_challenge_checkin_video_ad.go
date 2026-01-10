package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type AppChallengeCheckinVideoAd struct {
	ID            uint64          `gorm:"primaryKey;autoIncrement;comment:视频广告打卡ID"`
	CheckinID     uint64          `gorm:"not null;uniqueIndex:uk_checkin"`
	UserID        uint64          `gorm:"not null;index:idx_user"`
	AdPlatform    string          `gorm:"type:varchar(50);not null;default:''"`
	AdUnitID      string          `gorm:"type:varchar(100);not null;default:''"`
	AdOrderNo     string          `gorm:"type:varchar(100);not null;default:'';uniqueIndex:uk_ad_order"`
	VideoDuration int             `gorm:"not null;default:0"`
	WatchDuration int             `gorm:"not null;default:0"`
	RewardAmount  decimal.Decimal `gorm:"type:decimal(30,2);not null;default:0.00"`
	VerifyStatus  int             `gorm:"not null;default:0;index:idx_verify_status"`
	CreatedAt     time.Time       `gorm:"autoCreateTime;comment:观看完成时间"`
	VerifiedAt    *time.Time      `gorm:"comment:校验完成时间"`
}

func (AppChallengeCheckinVideoAd) TableName() string {
	return "app_challenge_checkin_video_ad"
}
