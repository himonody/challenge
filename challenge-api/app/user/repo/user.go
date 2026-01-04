package repo

import (
	"challenge/app/user/models"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *models.AppUser) error {
	return db.Table("app_user").Create(user).Error
}
