package repo

import (
	"challenge/app/withdraw/models"

	"gorm.io/gorm"
)

func CreateWithdrawOrder(db *gorm.DB, withdrawOrder *models.WithdrawOrder) error {
	return db.Table("withdraw_order").Create(withdrawOrder).Error
}
