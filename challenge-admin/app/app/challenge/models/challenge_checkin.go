package models

import "time"

type ChallengeCheckin struct {
	Id          uint64     `json:"id" gorm:"column:id;primaryKey;autoIncrement;comment:打卡ID"`
	ChallengeId uint64     `json:"challengeId" gorm:"column:challenge_id;type:bigint;comment:用户挑战ID"`
	UserId      uint64     `json:"userId" gorm:"column:user_id;type:bigint;comment:用户ID"`
	CheckinDate time.Time  `json:"checkinDate" gorm:"column:checkin_date;type:date;comment:打卡日期"`
	CheckinTime *time.Time `json:"checkinTime" gorm:"column:checkin_time;type:datetime;comment:打卡时间"`
	MoodCode    int8       `json:"moodCode" gorm:"column:mood_code;type:tinyint;comment:心情枚举"`
	MoodText    string     `json:"moodText" gorm:"column:mood_text;type:varchar(200);comment:心情描述"`
	ContentType int8       `json:"contentType" gorm:"column:content_type;type:tinyint;comment:打卡内容类型"`
	Status      int8       `json:"status" gorm:"column:status;type:tinyint;comment:状态"`
	CreatedAt   *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
}

func (ChallengeCheckin) TableName() string { return "app_challenge_checkin" }
