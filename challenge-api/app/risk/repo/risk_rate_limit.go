package repo

import (
	"challenge/app/risk/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// UpsertRateLimit 写入/累计频控记录
func UpsertRateLimit(db *gorm.DB, rec *models.RiskRateLimit) error {
	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "scene"}, {Name: "identity_type"}, {Name: "identity_value"}, {Name: "window_start"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"count": gorm.Expr("count + 1"), "window_end": rec.WindowEnd, "blocked": rec.Blocked}),
	}).Create(rec).Error
}

// GetRateLimit 获取频控记录
func GetRateLimit(db *gorm.DB, scene, identityType, identityValue string, windowStart int64) (*models.RiskRateLimit, error) {
	var rec models.RiskRateLimit
	err := db.Where("scene=? AND identity_type=? AND identity_value=? AND UNIX_TIMESTAMP(window_start)=?", scene, identityType, identityValue, windowStart).
		First(&rec).Error
	return &rec, err
}
