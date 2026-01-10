package models

import "time"

type AppChallengeCheckin struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement;comment:打卡ID"`
	ChallengeID uint64 `gorm:"not null;uniqueIndex:uk_challenge_date,priority:1;comment:用户挑战ID"`
	UserID      uint64 `gorm:"not null;index:idx_user_date,priority:1;comment:用户ID"`

	CheckinDate time.Time  `gorm:"type:date;not null;uniqueIndex:uk_challenge_date,priority:2;index:idx_user_date,priority:2;comment:打卡日期"`
	CheckinTime *time.Time `gorm:"comment:打卡时间"`

	MoodCode    uint8  `gorm:"not null;default:0;comment:心情枚举"`
	MoodText    string `gorm:"type:varchar(200);not null;default:'';comment:心情描述"`
	ContentType uint8  `gorm:"not null;default:1;comment:内容类型"`

	Status    uint8     `gorm:"not null;default:1;index:idx_status;comment:状态"`
	CreatedAt time.Time `gorm:"autoCreateTime;comment:记录创建时间"`
}

func (AppChallengeCheckin) TableName() string {
	return "app_challenge_checkin"
}
