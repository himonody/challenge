package repo

import (
	"challenge/app/risk/models"

	"gorm.io/gorm"
)

// ListActions 查询全部风险动作
func ListActions(db *gorm.DB) ([]models.RiskAction, error) {
	var list []models.RiskAction
	err := db.Find(&list).Error
	return list, err
}
