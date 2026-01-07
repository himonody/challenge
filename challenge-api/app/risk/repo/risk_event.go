package repo

import (
	"challenge/app/risk/models"
	"gorm.io/gorm"
)

// CreateRiskEvent 创建风控事件
func CreateRiskEvent(db *gorm.DB, event *models.RiskEvent) error {
	return db.Create(event).Error
}

// ListEventsByUserID 获取用户风控事件列表
func ListEventsByUserID(db *gorm.DB, userID uint64, limit int) ([]models.RiskEvent, error) {
	var events []models.RiskEvent
	err := db.Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(limit).
		Find(&events).Error
	return events, err
}

// SumScoreByUserID 统计用户风险总分
func SumScoreByUserID(db *gorm.DB, userID uint64) (int64, error) {
	var totalScore int64
	err := db.Model(&models.RiskEvent{}).
		Where("user_id = ?", userID).
		Select("COALESCE(SUM(score), 0)").
		Scan(&totalScore).Error
	return totalScore, err
}
