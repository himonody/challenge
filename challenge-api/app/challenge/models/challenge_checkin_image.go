package models

import (
	"time"
)

type ChallengeCheckinImage struct {
	Id        int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:图片ID"`
	CheckinId int64      `json:"checkinId" gorm:"column:checkin_id;type:bigint;index:idx_checkin;uniqueIndex:uk_checkin_hash,priority:1;comment:打卡ID"`
	UserId    int64      `json:"userId" gorm:"column:user_id;type:bigint;index:idx_user;comment:用户ID"`
	ImageUrl  string     `json:"imageUrl" gorm:"column:image_url;type:varchar(500);comment:图片URL"`
	ImageHash string     `json:"imageHash" gorm:"column:image_hash;type:varchar(64);uniqueIndex:uk_checkin_hash,priority:2;comment:图片Hash（防重复）"`
	SortNo    int64      `json:"sortNo" gorm:"column:sort_no;type:tinyint;comment:图片顺序"`
	Status    int64      `json:"status" gorm:"column:status;type:tinyint;comment:状态 1正常 2屏蔽 3审核中"`
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:上传时间戳"`

	Checkin *ChallengeCheckin `json:"checkin" gorm:"foreignkey:checkin_id"`
}

func (ChallengeCheckinImage) TableName() string {
	return "app_challenge_checkin_image"
}
