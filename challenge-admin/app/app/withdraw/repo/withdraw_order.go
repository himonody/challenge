package repo

import (
	"challenge-admin/app/app/withdraw/models"

	"gorm.io/gorm"
)

func CreateWithdrawOrder(db *gorm.DB, withdrawOrder *models.WithdrawOrder) error {
	return db.Table("app_withdraw_order").Create(withdrawOrder).Error
}
func UpdateWithdrawOrder(db *gorm.DB, withdrawOrder *models.WithdrawOrder) error {
	return db.Table("app_withdraw_order").Where("id = ?", withdrawOrder.Id).Updates(withdrawOrder).Error
}
