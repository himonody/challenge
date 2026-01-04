package repo

import (
	"challenge/app/user/models"

	"gorm.io/gorm"
)

func CreateUserOperLog(db *gorm.DB, log *models.AppUserOperLog) error {
	return db.Table("app_user_oper_log").Create(log).Error
}
