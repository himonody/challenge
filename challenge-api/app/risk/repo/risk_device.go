package repo

import (
	"challenge/app/risk/models"
	"gorm.io/gorm"
)

// CreateRiskDevice 创建设备记录
func CreateRiskDevice(db *gorm.DB, device *models.RiskDevice) error {
	return db.Create(device).Error
}

// GetDeviceByFP 通过设备指纹获取设备记录
func GetDeviceByFP(db *gorm.DB, deviceFP string) (*models.RiskDevice, error) {
	var device models.RiskDevice
	err := db.Where("device_fp = ?", deviceFP).First(&device).Error
	return &device, err
}

// ListDevicesByUserID 获取用户所有设备
func ListDevicesByUserID(db *gorm.DB, userID uint64) ([]models.RiskDevice, error) {
	var devices []models.RiskDevice
	err := db.Where("user_id = ?", userID).Find(&devices).Error
	return devices, err
}

// CountUsersByDeviceFP 统计设备关联用户数
func CountUsersByDeviceFP(db *gorm.DB, deviceFP string) (int64, error) {
	var count int64
	err := db.Model(&models.RiskDevice{}).Where("device_fp = ?", deviceFP).
		Distinct("user_id").Count(&count).Error
	return count, err
}
