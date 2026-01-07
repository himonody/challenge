package repo

import (
	"challenge/app/sse/models"
	"time"

	"gorm.io/gorm"
)

// CreateMessage 创建消息记录
func CreateMessage(db *gorm.DB, message *models.AppSSEMessage) error {
	return db.Create(message).Error
}

// GetMessageByEventID 根据事件ID获取消息
func GetMessageByEventID(db *gorm.DB, eventID string) (*models.AppSSEMessage, error) {
	var message models.AppSSEMessage
	err := db.Where("event_id = ?", eventID).First(&message).Error
	if err != nil {
		return nil, err
	}
	return &message, nil
}

// GetPendingMessages 获取待发送的消息
func GetPendingMessages(db *gorm.DB, receiverID string, receiverType string, limit int) ([]*models.AppSSEMessage, error) {
	var messages []*models.AppSSEMessage
	query := db.Where("receiver_id = ? AND receiver_type = ? AND status = ?",
		receiverID, receiverType, models.MessageStatusPending)

	// 未过期的消息
	query = query.Where("expire_at IS NULL OR expire_at > ?", time.Now())

	// 按优先级和创建时间排序
	query = query.Order("priority DESC, created_at ASC").Limit(limit)

	err := query.Find(&messages).Error
	return messages, err
}

// GetMessagesSince 获取指定时间之后的消息（用于重连恢复）
func GetMessagesSince(db *gorm.DB, receiverID string, receiverType string, lastEventID string, since time.Time, limit int) ([]*models.AppSSEMessage, error) {
	var messages []*models.AppSSEMessage
	query := db.Where("receiver_id = ? AND receiver_type = ?", receiverID, receiverType)

	if lastEventID != "" {
		// 如果有 lastEventID，从该事件之后开始
		query = query.Where("created_at > (SELECT created_at FROM app_sse_message WHERE event_id = ?)", lastEventID)
	} else {
		// 否则从指定时间开始
		query = query.Where("created_at > ?", since)
	}

	// 未过期的消息
	query = query.Where("expire_at IS NULL OR expire_at > ?", time.Now())

	query = query.Order("priority DESC, created_at ASC").Limit(limit)

	err := query.Find(&messages).Error
	return messages, err
}

// UpdateMessageStatus 更新消息状态
func UpdateMessageStatus(db *gorm.DB, eventID string, status int8) error {
	return db.Model(&models.AppSSEMessage{}).
		Where("event_id = ?", eventID).
		Update("status", status).Error
}

// BatchUpdateMessageStatus 批量更新消息状态
func BatchUpdateMessageStatus(db *gorm.DB, eventIDs []string, status int8) error {
	if len(eventIDs) == 0 {
		return nil
	}
	return db.Model(&models.AppSSEMessage{}).
		Where("event_id IN ?", eventIDs).
		Update("status", status).Error
}

// DeleteExpiredMessages 删除过期消息
func DeleteExpiredMessages(db *gorm.DB) error {
	return db.Where("expire_at IS NOT NULL AND expire_at < ?", time.Now()).
		Delete(&models.AppSSEMessage{}).Error
}

// DeleteOldMessages 删除旧消息（保留最近N天）
func DeleteOldMessages(db *gorm.DB, days int) error {
	return db.Where("created_at < ?", time.Now().AddDate(0, 0, -days)).
		Delete(&models.AppSSEMessage{}).Error
}

// GetMessageCountByReceiver 获取接收者的消息数量
func GetMessageCountByReceiver(db *gorm.DB, receiverID string, receiverType string, status *int8) (int64, error) {
	query := db.Model(&models.AppSSEMessage{}).
		Where("receiver_id = ? AND receiver_type = ?", receiverID, receiverType)

	if status != nil {
		query = query.Where("status = ?", *status)
	}

	var count int64
	err := query.Count(&count).Error
	return count, err
}
