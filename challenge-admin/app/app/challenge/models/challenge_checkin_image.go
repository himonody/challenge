package models

import "time"

type ChallengeCheckinImage struct {
	Id        uint64     `json:"id" gorm:"column:id;primaryKey;autoIncrement;comment:图片ID"`
	CheckinId uint64     `json:"checkinId" gorm:"column:checkin_id;type:bigint;comment:打卡ID"`
	UserId    uint64     `json:"userId" gorm:"column:user_id;type:bigint;comment:用户ID"`
	ImageUrl  string     `json:"imageUrl" gorm:"column:image_url;type:varchar(500);comment:图片URL"`
	ImageHash string     `json:"imageHash" gorm:"column:image_hash;type:varchar(64);comment:图片Hash"`
	SortNo    int8       `json:"sortNo" gorm:"column:sort_no;type:tinyint;comment:图片顺序"`
	Status    int8       `json:"status" gorm:"column:status;type:tinyint;comment:状态"`
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:上传时间"`
}

func (ChallengeCheckinImage) TableName() string { return "app_challenge_checkin_image" }
