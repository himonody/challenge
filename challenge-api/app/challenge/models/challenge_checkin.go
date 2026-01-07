package models

import (
	"time"
)

type ChallengeCheckin struct {
	Id          int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:打卡ID"`
	ChallengeId int64      `json:"challengeId" gorm:"column:challenge_id;type:bigint;uniqueIndex:uk_challenge_date,priority:1;comment:用户挑战ID"`
	UserId      int64      `json:"userId" gorm:"column:user_id;type:bigint;index:idx_user_date,priority:1;comment:用户ID"`
	CheckinDate *time.Time `json:"checkinDate" gorm:"column:checkin_date;type:date;uniqueIndex:uk_challenge_date,priority:2;index:idx_user_date,priority:2;comment:打卡日期 YYYYMMDD"`
	CheckinTime *time.Time `json:"checkinTime" gorm:"column:checkin_time;type:datetime;comment:打卡时间戳"`
	MoodCode    int64      `json:"moodCode" gorm:"column:mood_code;type:tinyint;comment:心情枚举 1开心 2平静 3一般 4疲惫 5低落 6爆棚"`
	MoodText    string     `json:"moodText" gorm:"column:mood_text;type:varchar(200);comment:用户心情文字描述（最多200字）"`
	ContentType int64      `json:"contentType" gorm:"column:content_type;type:tinyint;comment:打卡内容类型 1图片 2视频广告"`
	Status      int64      `json:"status" gorm:"column:status;type:tinyint;index:idx_status;comment:状态 1成功 2超时"`
	CreatedAt   *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:记录创建时间戳"`

	ChallengeUser *ChallengeUser           `json:"challengeUser" gorm:"foreignkey:challenge_id"`
	Images        []ChallengeCheckinImage  `json:"images" gorm:"foreignkey:checkin_id"`
	VideoAd       *ChallengeCheckinVideoAd `json:"videoAd" gorm:"foreignkey:checkin_id"`
}

func (ChallengeCheckin) TableName() string {
	return "app_challenge_checkin"
}
