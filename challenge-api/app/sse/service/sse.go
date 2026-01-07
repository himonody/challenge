package service

import (
	"challenge/app/sse/models"
	"challenge/app/sse/repo"
	"challenge/app/sse/storage"
	"challenge/core/runtime"
	"challenge/core/sse"
	"challenge/core/utils/log"
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// SSEService SSE 业务服务
type SSEService struct {
	Orm  *gorm.DB
	Run  runtime.Runtime
	Lang string
}

// NewSSEService 创建 SSE 服务
func NewSSEService(orm *gorm.DB, run runtime.Runtime, lang string) *SSEService {
	return &SSEService{
		Orm:  orm,
		Run:  run,
		Lang: lang,
	}
}

// SendToUser 发送消息给指定用户
func (s *SSEService) SendToUser(userID string, eventType string, data interface{}, options ...MessageOption) error {
	// 应用选项
	opt := &messageOptions{
		priority: models.PriorityNormal,
		persist:  true,
	}
	for _, o := range options {
		o(opt)
	}

	// 生成事件ID
	eventID := opt.eventID
	if eventID == "" {
		eventID = fmt.Sprintf("evt_%s_%d", userID, time.Now().UnixNano())
	}

	// 创建事件
	event := sse.NewEventWithID(eventID, eventType, data)

	// 立即发送（如果用户在线）
	sse.SSEManager.SendToUser(userID, event)

	// 持久化消息（用于离线推送和重连恢复）
	if opt.persist {
		if err := s.persistMessage(eventID, eventType, userID, models.ReceiverTypeUser, "", data, opt); err != nil {
			log.Errorf("Failed to persist message: %v", err)
			// 不影响发送流程
		}
	}

	// 更新未读计数
	cache := s.Run.GetCacheAdapter()
	if err := storage.IncrUnreadCount(cache, userID); err != nil {
		log.Errorf("Failed to increment unread count: %v", err)
	}

	return nil
}

// SendToGroup 发送消息到指定分组
func (s *SSEService) SendToGroup(group string, eventType string, data interface{}, options ...MessageOption) error {
	// 应用选项
	opt := &messageOptions{
		priority: models.PriorityNormal,
		persist:  false, // 分组消息默认不持久化
	}
	for _, o := range options {
		o(opt)
	}

	// 生成事件ID
	eventID := opt.eventID
	if eventID == "" {
		eventID = fmt.Sprintf("evt_group_%s_%d", group, time.Now().UnixNano())
	}

	// 创建事件
	event := sse.NewEventWithID(eventID, eventType, data)

	// 发送到分组
	sse.SSEManager.SendToGroup(group, event)

	// 持久化（可选）
	if opt.persist {
		// 获取分组订阅者
		subscribers, err := repo.GetGroupSubscribers(s.Orm, group)
		if err != nil {
			log.Errorf("Failed to get group subscribers: %v", err)
		} else {
			// 为每个订阅者持久化消息
			for _, sub := range subscribers {
				if err := s.persistMessage(eventID, eventType, sub.UserID, models.ReceiverTypeUser, group, data, opt); err != nil {
					log.Errorf("Failed to persist message for user %s: %v", sub.UserID, err)
				}
			}
		}
	}

	return nil
}

// Broadcast 广播消息到所有在线客户端
func (s *SSEService) Broadcast(eventType string, data interface{}) error {
	// 生成事件ID
	eventID := fmt.Sprintf("evt_broadcast_%d", time.Now().UnixNano())

	// 创建事件
	event := sse.NewEventWithID(eventID, eventType, data)

	// 广播
	sse.SSEManager.Broadcast(event)

	return nil
}

// GetPendingMessages 获取用户的待发送消息（用于重连恢复）
func (s *SSEService) GetPendingMessages(userID string, lastEventID string, limit int) ([]*models.AppSSEMessage, error) {
	// 从最近5分钟开始恢复
	since := time.Now().Add(-5 * time.Minute)
	return repo.GetMessagesSince(s.Orm, userID, models.ReceiverTypeUser, lastEventID, since, limit)
}

// MarkMessageRead 标记消息已读
func (s *SSEService) MarkMessageRead(eventID string, userID string) error {
	// 更新数据库
	if err := repo.UpdateMessageStatus(s.Orm, eventID, models.MessageStatusRead); err != nil {
		return err
	}

	// 减少未读计数
	cache := s.Run.GetCacheAdapter()
	if err := storage.DecrUnreadCount(cache, userID, 1); err != nil {
		log.Errorf("Failed to decrease unread count: %v", err)
	}

	return nil
}

// GetUnreadCount 获取未读消息数量
func (s *SSEService) GetUnreadCount(userID string) (int, error) {
	cache := s.Run.GetCacheAdapter()

	// 先从缓存获取
	count, err := storage.GetUnreadCount(cache, userID)
	if err != nil {
		// 缓存失败，从数据库获取
		status := int8(models.MessageStatusPending)
		count64, err := repo.GetMessageCountByReceiver(s.Orm, userID, models.ReceiverTypeUser, &status)
		if err != nil {
			return 0, err
		}
		count = int(count64)
	}

	return count, nil
}

// Subscribe 订阅分组
func (s *SSEService) Subscribe(userID string, group string, eventTypes []string) error {
	sub := &models.AppSSESubscription{
		UserID:     userID,
		GroupName:  group,
		EventTypes: fmt.Sprintf(",%s,", joinStrings(eventTypes, ",")), // 前后加逗号便于查询
		Status:     models.SubscriptionStatusEnabled,
	}

	return repo.UpsertSubscription(s.Orm, sub)
}

// Unsubscribe 取消订阅
func (s *SSEService) Unsubscribe(userID string, group string) error {
	return repo.DeleteSubscription(s.Orm, userID, group)
}

// GetUserSubscriptions 获取用户的订阅列表
func (s *SSEService) GetUserSubscriptions(userID string) ([]*models.AppSSESubscription, error) {
	return repo.GetUserSubscriptions(s.Orm, userID)
}

// CleanupExpiredMessages 清理过期消息（定时任务）
func (s *SSEService) CleanupExpiredMessages() error {
	return repo.DeleteExpiredMessages(s.Orm)
}

// CleanupOldMessages 清理旧消息（定时任务，保留最近N天）
func (s *SSEService) CleanupOldMessages(days int) error {
	return repo.DeleteOldMessages(s.Orm, days)
}

// ========== 私有方法 ==========

// persistMessage 持久化消息
func (s *SSEService) persistMessage(eventID string, eventType string, receiverID string, receiverType string, group string, data interface{}, opt *messageOptions) error {
	// 序列化数据
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// 计算过期时间
	var expireAt *time.Time
	if opt.ttl > 0 {
		t := time.Now().Add(time.Duration(opt.ttl) * time.Second)
		expireAt = &t
	}

	// 创建消息记录
	message := &models.AppSSEMessage{
		EventID:      eventID,
		EventType:    eventType,
		ReceiverID:   receiverID,
		ReceiverType: receiverType,
		GroupName:    group,
		Priority:     opt.priority,
		Data:         string(dataJSON),
		Status:       models.MessageStatusPending,
		ExpireAt:     expireAt,
	}

	return repo.CreateMessage(s.Orm, message)
}

// ========== 消息选项 ==========

type messageOptions struct {
	eventID  string // 自定义事件ID
	priority int    // 优先级
	persist  bool   // 是否持久化
	ttl      int    // 过期时间（秒）
}

type MessageOption func(*messageOptions)

// WithEventID 设置事件ID
func WithEventID(eventID string) MessageOption {
	return func(o *messageOptions) {
		o.eventID = eventID
	}
}

// WithPriority 设置优先级
func WithPriority(priority int) MessageOption {
	return func(o *messageOptions) {
		o.priority = priority
	}
}

// WithPersist 设置是否持久化
func WithPersist(persist bool) MessageOption {
	return func(o *messageOptions) {
		o.persist = persist
	}
}

// WithTTL 设置过期时间（秒）
func WithTTL(ttl int) MessageOption {
	return func(o *messageOptions) {
		o.ttl = ttl
	}
}

// ========== 工具函数 ==========

func joinStrings(strs []string, sep string) string {
	result := ""
	for i, s := range strs {
		if i > 0 {
			result += sep
		}
		result += s
	}
	return result
}
