package repo

import (
	"challenge/app/sse/models"

	"gorm.io/gorm"
)

// CreateSubscription 创建订阅
func CreateSubscription(db *gorm.DB, sub *models.AppSSESubscription) error {
	return db.Create(sub).Error
}

// GetSubscription 获取订阅
func GetSubscription(db *gorm.DB, userID string, groupName string) (*models.AppSSESubscription, error) {
	var sub models.AppSSESubscription
	err := db.Where("user_id = ? AND group_name = ?", userID, groupName).First(&sub).Error
	if err != nil {
		return nil, err
	}
	return &sub, nil
}

// GetUserSubscriptions 获取用户的所有订阅
func GetUserSubscriptions(db *gorm.DB, userID string) ([]*models.AppSSESubscription, error) {
	var subs []*models.AppSSESubscription
	err := db.Where("user_id = ? AND status = ?", userID, models.SubscriptionStatusEnabled).Find(&subs).Error
	return subs, err
}

// GetGroupSubscribers 获取分组的所有订阅者
func GetGroupSubscribers(db *gorm.DB, groupName string) ([]*models.AppSSESubscription, error) {
	var subs []*models.AppSSESubscription
	err := db.Where("group_name = ? AND status = ?", groupName, models.SubscriptionStatusEnabled).Find(&subs).Error
	return subs, err
}

// UpdateSubscription 更新订阅
func UpdateSubscription(db *gorm.DB, userID string, groupName string, eventTypes string, status int8) error {
	updates := map[string]interface{}{
		"event_types": eventTypes,
		"status":      status,
	}
	return db.Model(&models.AppSSESubscription{}).
		Where("user_id = ? AND group_name = ?", userID, groupName).
		Updates(updates).Error
}

// DeleteSubscription 删除订阅
func DeleteSubscription(db *gorm.DB, userID string, groupName string) error {
	return db.Where("user_id = ? AND group_name = ?", userID, groupName).
		Delete(&models.AppSSESubscription{}).Error
}

// UpsertSubscription 创建或更新订阅
func UpsertSubscription(db *gorm.DB, sub *models.AppSSESubscription) error {
	var existing models.AppSSESubscription
	err := db.Where("user_id = ? AND group_name = ?", sub.UserID, sub.GroupName).First(&existing).Error

	if err == gorm.ErrRecordNotFound {
		// 不存在，创建
		return db.Create(sub).Error
	}

	if err != nil {
		return err
	}

	// 已存在，更新
	return db.Model(&existing).Updates(map[string]interface{}{
		"event_types": sub.EventTypes,
		"status":      sub.Status,
	}).Error
}
