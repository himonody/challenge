package repo

import (
	"challenge/app/risk/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ListStrategiesByScene 获取场景下所有策略（按优先级升序）
func ListStrategiesByScene(db *gorm.DB, scene string) ([]models.RiskStrategy, error) {
	var list []models.RiskStrategy
	err := db.Where("scene = ? AND status = 1", scene).
		Order("priority asc").
		Find(&list).Error
	return list, err
}

// UpsertStrategyCache 将策略写入缓存表
func UpsertStrategyCache(db *gorm.DB, caches []models.RiskStrategyCache) error {
	if len(caches) == 0 {
		return nil
	}
	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "scene"}, {Name: "identity_type"}, {Name: "rule_code"}},
		DoUpdates: clause.AssignmentColumns([]string{"window_seconds", "threshold", "action", "action_value"}),
	}).Create(&caches).Error
}

// ListStrategyCacheByScene 读取缓存表
func ListStrategyCacheByScene(db *gorm.DB, scene string) ([]models.RiskStrategyCache, error) {
	var list []models.RiskStrategyCache
	err := db.Where("scene = ?", scene).Find(&list).Error
	return list, err
}
