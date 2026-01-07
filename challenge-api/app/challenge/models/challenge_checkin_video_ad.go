package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type ChallengeCheckinVideoAd struct {
	Id            int64           `json:"id" gorm:"primaryKey;autoIncrement;comment:视频广告打卡ID"`
	CheckinId     int64           `json:"checkinId" gorm:"column:checkin_id;type:bigint;uniqueIndex:uk_checkin;comment:关联打卡ID"`
	UserId        int64           `json:"userId" gorm:"column:user_id;type:bigint;index:idx_user;comment:用户ID"`
	AdPlatform    string          `json:"adPlatform" gorm:"column:ad_platform;type:varchar(50);comment:广告平台 如：csj、gdt、unity"`
	AdUnitId      string          `json:"adUnitId" gorm:"column:ad_unit_id;type:varchar(100);comment:广告位ID"`
	AdOrderNo     string          `json:"adOrderNo" gorm:"column:ad_order_no;type:varchar(100);uniqueIndex:uk_ad_order;comment:广告联盟返回的订单号（唯一）"`
	VideoDuration int64           `json:"videoDuration" gorm:"column:video_duration;type:int;comment:视频时长（秒）"`
	WatchDuration int64           `json:"watchDuration" gorm:"column:watch_duration;type:int;comment:实际观看时长（秒）"`
	RewardAmount  decimal.Decimal `json:"rewardAmount" gorm:"column:reward_amount;type:decimal(30,2);comment:该广告产生的收益"`
	VerifyStatus  int64           `json:"verifyStatus" gorm:"column:verify_status;type:tinyint;index:idx_verify_status;comment:校验状态 0待校验 1成功 2失败"`
	CreatedAt     *time.Time      `json:"createdAt" gorm:"column:created_at;type:datetime;comment:观看完成时间戳"`
	VerifiedAt    *time.Time      `json:"verifiedAt" gorm:"column:verified_at;type:datetime;comment:校验完成时间戳"`

	Checkin *ChallengeCheckin `json:"checkin" gorm:"foreignkey:checkin_id"`
}

func (ChallengeCheckinVideoAd) TableName() string {
	return "app_challenge_checkin_video_ad"
}
