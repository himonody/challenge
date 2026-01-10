package models

import "time"

type AppChallengeCheckinImage struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement;comment:图片ID"`
	CheckinID uint64    `gorm:"not null;uniqueIndex:uk_checkin_hash,priority:1;index:idx_checkin"`
	UserID    uint64    `gorm:"not null;index:idx_user"`
	ImageURL  string    `gorm:"type:varchar(500);not null;default:''"`
	ImageHash string    `gorm:"type:varchar(64);not null;default:'';uniqueIndex:uk_checkin_hash,priority:2"`
	SortNo    uint8     `gorm:"not null;default:1"`
	Status    uint8     `gorm:"not null;default:1"`
	CreatedAt time.Time `gorm:"autoCreateTime;comment:上传时间"`
}

func (AppChallengeCheckinImage) TableName() string {
	return "app_challenge_checkin_image"
}
