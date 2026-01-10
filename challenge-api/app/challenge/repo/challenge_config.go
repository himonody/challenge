package repo

import (
	"challenge/app/challenge/models"

	"gorm.io/gorm"
)

func GetChallengeConfig(db *gorm.DB) ([]*models.AppChallengeConfig, error) {
	configs := make([]*models.AppChallengeConfig, 0)
	err := db.Table("app_challenge_config").Where("status = 1").Order("sort DESC").Find(&configs).Error
	return configs, err
}
