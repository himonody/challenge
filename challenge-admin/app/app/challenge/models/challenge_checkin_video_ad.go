package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type ChallengeCheckinVideoAd struct {
	Id            uint64          `json:"id" gorm:"column:id;primaryKey;autoIncrement;comment:视频广告打卡ID"`
	CheckinId     uint64          `json:"checkinId" gorm:"column:checkin_id;type:bigint;comment:打卡ID"`
	UserId        uint64          `json:"userId" gorm:"column:user_id;type:bigint;comment:用户ID"`
	AdPlatform    string          `json:"adPlatform" gorm:"column:ad_platform;type:varchar(50);comment:广告平台"`
	AdUnitId      string          `json:"adUnitId" gorm:"column:ad_unit_id;type:varchar(100);comment:广告位ID"`
	AdOrderNo     string          `json:"adOrderNo" gorm:"column:ad_order_no;type:varchar(100);comment:广告订单号"`
	VideoDuration int             `json:"videoDuration" gorm:"column:video_duration;type:int;comment:视频时长"`
	WatchDuration int             `json:"watchDuration" gorm:"column:watch_duration;type:int;comment:观看时长"`
	RewardAmount  decimal.Decimal `json:"rewardAmount" gorm:"column:reward_amount;type:decimal(30,2);comment:收益金额"`
	VerifyStatus  int8            `json:"verifyStatus" gorm:"column:verify_status;type:tinyint;comment:校验状态"`
	CreatedAt     *time.Time      `json:"createdAt" gorm:"column:created_at;type:datetime;comment:观看完成时间"`
	VerifiedAt    *time.Time      `json:"verifiedAt" gorm:"column:verified_at;type:datetime;comment:校验完成时间"`
}

func (ChallengeCheckinVideoAd) TableName() string { return "app_challenge_checkin_video_ad" }
