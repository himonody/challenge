package repo

import (
	"challenge/app/challenge/models"

	"gorm.io/gorm"
)

func GetChallengeConfig(db *gorm.DB) (*models.AppChallengeConfig, error) {
	config := new(models.AppChallengeConfig)
	err := db.Table("app_challenge_config").Where("status = 1").Order("sort DESC").Last(&config).Error
	return config, err
}
