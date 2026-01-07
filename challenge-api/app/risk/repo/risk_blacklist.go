package repo

import (
	"challenge/app/risk/models"

	"gorm.io/gorm"
)

// GetActiveBlacklist 查询命中的黑名单
func GetActiveBlacklist(db *gorm.DB, typ, value string) (*models.RiskBlacklist, error) {
	var item models.RiskBlacklist
	err := db.Where("type=? AND value=? AND status='1'", typ, value).First(&item).Error
	return &item, err
}
