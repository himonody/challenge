package repo

import (
	"challenge/app/risk/models"
	"gorm.io/gorm"
)

// GetRiskUser 获取风控用户信息
func GetRiskUser(db *gorm.DB, userID uint64) (*models.RiskUser, error) {
	var riskUser models.RiskUser
	err := db.Where("user_id = ?", userID).First(&riskUser).Error
	return &riskUser, err
}

// UpsertRiskUser 创建或更新风控用户
func UpsertRiskUser(db *gorm.DB, riskUser *models.RiskUser) error {
	return db.Save(riskUser).Error
}

// UpdateRiskLevel 更新风险等级
func UpdateRiskLevel(db *gorm.DB, userID uint64, riskLevel int8, riskScore int64, reason string) error {
	return db.Model(&models.RiskUser{}).
		Where("user_id = ?", userID).
		Updates(map[string]interface{}{
			"risk_level": riskLevel,
			"risk_score": riskScore,
			"reason":     reason,
		}).Error
}

// IncrementRiskScore 增加风险分数
func IncrementRiskScore(db *gorm.DB, userID uint64, delta int64) error {
	return db.Model(&models.RiskUser{}).
		Where("user_id = ?", userID).
		Update("risk_score", gorm.Expr("risk_score + ?", delta)).Error
}
