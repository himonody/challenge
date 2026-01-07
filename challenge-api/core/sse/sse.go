package sse

import (
	"challenge/core/utils/log"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
)

// Manager SSE 连接管理器（优化版）
type Manager struct {
	clients     sync.Map               // clientID -> *Client（线程安全）
	groups      sync.Map               // groupName -> sync.Map(clientID -> *Client)
	clientCount int64                  // 客户端计数（原子操作）
	groupCount  int64                  // 分组计数（原子操作）
	register    chan *Client           // 注册通道
	unregister  chan *Client           // 注销通道
	broadcast   chan *BroadcastMessage // 广播消息
	groupcast   chan *GroupMessage     // 组播消息
	unicast     chan *UnicastMessage   // 单播消息

	// 配置项
	HeartbeatInterval time.Duration // 心跳间隔
	ClientTimeout     time.Duration // 客户端超时时间
	SendTimeout       time.Duration // 发送超时时间
}

// Client SSE 客户端（优化版）
type Client struct {
	ID           string             // 客户端唯一标识
	UserID       string             // 用户ID（业务层关联）
	Group        string             // 所属组
	Writer       gin.ResponseWriter // Response Writer
	Request      *http.Request      // HTTP Request
	send         chan *Event        // 消息发送通道
	closeChan    chan struct{}      // 关闭信号
	lastEventID  string             // 最后接收的事件ID（用于重连恢复）
	connectedAt  time.Time          // 连接建立时间
	lastActiveAt time.Time          // 最后活跃时间
	metadata     map[string]string  // 客户端元数据（如IP、UA）
	closeOnce    sync.Once          // 确保只关闭一次
	isClosed     atomic.Bool        // 是否已关闭
}

// Event SSE 事件（优化版）
type Event struct {
	ID      string      // 事件ID（可选，用于重连恢复）
	Event   string      // 事件类型（可选）
	Data    interface{} // 事件数据
	Retry   int         // 重连时间（毫秒，可选）
	Comment string      // 注释（用于心跳等，可选）
}

// BroadcastMessage 广播消息
type BroadcastMessage struct {
	Event     *Event
	ExcludeID string // 排除的客户端ID（可选）
}

// GroupMessage 组播消息
type GroupMessage struct {
	Group     string
	Event     *Event
	ExcludeID string // 排除的客户端ID（可选）
}

// UnicastMessage 单播消息
type UnicastMessage struct {
	ClientID string
	Event    *Event
}

// NewManager 创建 SSE 管理器
func NewManager() *Manager {
	return &Manager{
		register:          make(chan *Client, 128),
		unregister:        make(chan *Client, 128),
		broadcast:         make(chan *BroadcastMessage, 256),
		groupcast:         make(chan *GroupMessage, 256),
		unicast:           make(chan *UnicastMessage, 256),
		HeartbeatInterval: 30 * time.Second,
		ClientTimeout:     90 * time.Second,
		SendTimeout:       5 * time.Second,
	}
}

// NewEvent 创建简单事件
func NewEvent(data interface{}) *Event {
	return &Event{Data: data}
}

// NewTypedEvent 创建带类型的事件
func NewTypedEvent(eventType string, data interface{}) *Event {
	return &Event{Event: eventType, Data: data}
}

// NewEventWithID 创建带ID的事件（用于重连恢复）
func NewEventWithID(id string, eventType string, data interface{}) *Event {
	return &Event{ID: id, Event: eventType, Data: data}
}

// NewCommentEvent 创建注释事件（用于心跳）
func NewCommentEvent(comment string) *Event {
	return &Event{Comment: comment}
}

// Format 格式化为 SSE 协议格式
func (e *Event) Format() string {
	var result string

	// 注释（心跳）
	if e.Comment != "" {
		return fmt.Sprintf(": %s\n\n", e.Comment)
	}

	// 事件ID
	if e.ID != "" {
		result += fmt.Sprintf("id: %s\n", e.ID)
	}

	// 事件类型
	if e.Event != "" {
		result += fmt.Sprintf("event: %s\n", e.Event)
	}

	// 重连时间
	if e.Retry > 0 {
		result += fmt.Sprintf("retry: %d\n", e.Retry)
	}

	// 事件数据
	var dataStr string
	switch v := e.Data.(type) {
	case string:
		dataStr = v
	case []byte:
		dataStr = string(v)
	case nil:
		dataStr = ""
	default:
		// JSON 序列化
		jsonData, err := json.Marshal(v)
		if err != nil {
			log.Errorf("Failed to marshal event data: %v", err)
			dataStr = fmt.Sprintf("%v", v)
		} else {
			dataStr = string(jsonData)
		}
	}

	// SSE 数据格式（支持多行）
	if dataStr != "" {
		result += fmt.Sprintf("data: %s\n", dataStr)
	}

	result += "\n"
	return result
}

// Start 启动 SSE 管理器
func (m *Manager) Start() {
	log.Info("SSE Manager started")

	// 启动清理协程
	go m.cleanupInactiveClients()

	for {
		select {
		case client := <-m.register:
			m.registerClient(client)

		case client := <-m.unregister:
			m.unregisterClient(client)

		case msg := <-m.unicast:
			m.sendUnicast(msg)

		case msg := <-m.groupcast:
			m.sendGroupcast(msg)

		case msg := <-m.broadcast:
			m.sendBroadcast(msg)
		}
	}
}

// registerClient 注册客户端
func (m *Manager) registerClient(client *Client) {
	// 存储到全局客户端列表
	m.clients.Store(client.ID, client)
	atomic.AddInt64(&m.clientCount, 1)

	// 添加到分组
	if client.Group != "" {
		groupClientsInterface, _ := m.groups.LoadOrStore(client.Group, &sync.Map{})
		groupClients := groupClientsInterface.(*sync.Map)

		// 检查组是否是新创建的
		var isNewGroup bool
		groupClients.Range(func(key, value interface{}) bool {
			isNewGroup = false
			return false // 找到任意一个就停止，说明组已存在
		})
		if isNewGroup {
			atomic.AddInt64(&m.groupCount, 1)
		}

		groupClients.Store(client.ID, client)
	}

	log.Infof("SSE client registered: id=%s, userID=%s, group=%s, totalClients=%d",
		client.ID, client.UserID, client.Group, atomic.LoadInt64(&m.clientCount))
}

// unregisterClient 注销客户端
func (m *Manager) unregisterClient(client *Client) {
	// 从全局列表删除
	if _, ok := m.clients.LoadAndDelete(client.ID); ok {
		atomic.AddInt64(&m.clientCount, -1)

		// 从分组删除
		if client.Group != "" {
			if groupClientsInterface, ok := m.groups.Load(client.Group); ok {
				groupClients := groupClientsInterface.(*sync.Map)
				groupClients.Delete(client.ID)

				// 如果组为空，删除组
				var isEmpty = true
				groupClients.Range(func(key, value interface{}) bool {
					isEmpty = false
					return false
				})
				if isEmpty {
					m.groups.Delete(client.Group)
					atomic.AddInt64(&m.groupCount, -1)
				}
			}
		}

		// 关闭客户端
		client.Close()

		log.Infof("SSE client unregistered: id=%s, group=%s, remainingClients=%d",
			client.ID, client.Group, atomic.LoadInt64(&m.clientCount))
	}
}

// sendUnicast 发送单播消息
func (m *Manager) sendUnicast(msg *UnicastMessage) {
	if clientInterface, ok := m.clients.Load(msg.ClientID); ok {
		client := clientInterface.(*Client)
		client.Send(msg.Event)
	} else {
		log.Warnf("SSE unicast failed: client not found, clientID=%s", msg.ClientID)
	}
}

// sendGroupcast 发送组播消息
func (m *Manager) sendGroupcast(msg *GroupMessage) {
	if groupClientsInterface, ok := m.groups.Load(msg.Group); ok {
		groupClients := groupClientsInterface.(*sync.Map)
		count := 0

		groupClients.Range(func(key, value interface{}) bool {
			client := value.(*Client)
			if msg.ExcludeID == "" || client.ID != msg.ExcludeID {
				client.Send(msg.Event)
				count++
			}
			return true
		})

		log.Debugf("SSE groupcast sent: group=%s, recipients=%d", msg.Group, count)
	} else {
		log.Warnf("SSE groupcast failed: group not found, group=%s", msg.Group)
	}
}

// sendBroadcast 发送广播消息
func (m *Manager) sendBroadcast(msg *BroadcastMessage) {
	count := 0

	m.clients.Range(func(key, value interface{}) bool {
		client := value.(*Client)
		if msg.ExcludeID == "" || client.ID != msg.ExcludeID {
			client.Send(msg.Event)
			count++
		}
		return true
	})

	log.Debugf("SSE broadcast sent: recipients=%d", count)
}

// cleanupInactiveClients 清理不活跃的客户端
func (m *Manager) cleanupInactiveClients() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now()
		var toRemove []*Client

		m.clients.Range(func(key, value interface{}) bool {
			client := value.(*Client)
			if now.Sub(client.lastActiveAt) > m.ClientTimeout {
				toRemove = append(toRemove, client)
			}
			return true
		})

		for _, client := range toRemove {
			log.Warnf("SSE client timeout: id=%s, lastActive=%v", client.ID, client.lastActiveAt)
			m.unregister <- client
		}
	}
}

// Send 发送消息给指定客户端
func (m *Manager) Send(clientID string, event *Event) {
	m.unicast <- &UnicastMessage{
		ClientID: clientID,
		Event:    event,
	}
}

// SendToUser 发送消息给指定用户的所有连接
func (m *Manager) SendToUser(userID string, event *Event) {
	m.clients.Range(func(key, value interface{}) bool {
		client := value.(*Client)
		if client.UserID == userID {
			client.Send(event)
		}
		return true
	})
}

// SendToGroup 发送消息到指定分组
func (m *Manager) SendToGroup(group string, event *Event) {
	m.groupcast <- &GroupMessage{
		Group: group,
		Event: event,
	}
}

// SendToGroupExcept 发送消息到指定分组（排除某个客户端）
func (m *Manager) SendToGroupExcept(group string, event *Event, excludeClientID string) {
	m.groupcast <- &GroupMessage{
		Group:     group,
		Event:     event,
		ExcludeID: excludeClientID,
	}
}

// Broadcast 广播消息到所有客户端
func (m *Manager) Broadcast(event *Event) {
	m.broadcast <- &BroadcastMessage{
		Event: event,
	}
}

// BroadcastExcept 广播消息到所有客户端（排除某个客户端）
func (m *Manager) BroadcastExcept(event *Event, excludeClientID string) {
	m.broadcast <- &BroadcastMessage{
		Event:     event,
		ExcludeID: excludeClientID,
	}
}

// GetClient 获取指定客户端
func (m *Manager) GetClient(clientID string) (*Client, bool) {
	clientInterface, ok := m.clients.Load(clientID)
	if !ok {
		return nil, false
	}
	return clientInterface.(*Client), true
}

// GetClientCount 获取客户端总数
func (m *Manager) GetClientCount() int64 {
	return atomic.LoadInt64(&m.clientCount)
}

// GetGroupCount 获取分组总数
func (m *Manager) GetGroupCount() int64 {
	return atomic.LoadInt64(&m.groupCount)
}

// GetGroupClientCount 获取指定分组的客户端数量
func (m *Manager) GetGroupClientCount(group string) int {
	if groupClientsInterface, ok := m.groups.Load(group); ok {
		groupClients := groupClientsInterface.(*sync.Map)
		count := 0
		groupClients.Range(func(key, value interface{}) bool {
			count++
			return true
		})
		return count
	}
	return 0
}

// GetClientsByGroup 获取指定分组的所有客户端ID
func (m *Manager) GetClientsByGroup(group string) []string {
	var clientIDs []string

	if groupClientsInterface, ok := m.groups.Load(group); ok {
		groupClients := groupClientsInterface.(*sync.Map)
		groupClients.Range(func(key, value interface{}) bool {
			client := value.(*Client)
			clientIDs = append(clientIDs, client.ID)
			return true
		})
	}

	return clientIDs
}

// Info 获取管理器信息
func (m *Manager) Info() map[string]interface{} {
	return map[string]interface{}{
		"clientCount":       atomic.LoadInt64(&m.clientCount),
		"groupCount":        atomic.LoadInt64(&m.groupCount),
		"registerChanLen":   len(m.register),
		"unregisterChanLen": len(m.unregister),
		"unicastChanLen":    len(m.unicast),
		"groupcastChanLen":  len(m.groupcast),
		"broadcastChanLen":  len(m.broadcast),
		"heartbeatInterval": m.HeartbeatInterval.String(),
		"clientTimeout":     m.ClientTimeout.String(),
	}
}

// ========== Client 方法 ==========

// Send 发送事件到客户端
func (c *Client) Send(event *Event) {
	if c.isClosed.Load() {
		return
	}

	select {
	case c.send <- event:
		// 发送成功
	default:
		log.Warnf("SSE client send channel full: clientID=%s", c.ID)
	}
}

// Close 关闭客户端连接
func (c *Client) Close() {
	c.closeOnce.Do(func() {
		c.isClosed.Store(true)
		close(c.closeChan)
		close(c.send)
	})
}

// IsClosed 检查客户端是否已关闭
func (c *Client) IsClosed() bool {
	return c.isClosed.Load()
}

// UpdateLastActive 更新最后活跃时间
func (c *Client) UpdateLastActive() {
	c.lastActiveAt = time.Now()
}

// GetMetadata 获取元数据
func (c *Client) GetMetadata(key string) string {
	return c.metadata[key]
}

// SetMetadata 设置元数据
func (c *Client) SetMetadata(key, value string) {
	if c.metadata == nil {
		c.metadata = make(map[string]string)
	}
	c.metadata[key] = value
}

// writePump 写入泵（发送消息到客户端）
func (c *Client) writePump(manager *Manager) {
	defer func() {
		log.Infof("SSE client writePump stopped: clientID=%s", c.ID)
		manager.unregister <- c
	}()

	// 心跳 ticker
	heartbeatTicker := time.NewTicker(manager.HeartbeatInterval)
	defer heartbeatTicker.Stop()

	// 发送连接成功事件
	welcomeEvent := NewTypedEvent("connected", map[string]interface{}{
		"clientId": c.ID,
		"userId":   c.UserID,
		"group":    c.Group,
		"time":     time.Now().Format(time.RFC3339),
	})
	if err := c.writeEvent(welcomeEvent); err != nil {
		log.Errorf("SSE failed to send welcome event: clientID=%s, error=%v", c.ID, err)
		return
	}

	for {
		select {
		case <-c.closeChan:
			return

		case event, ok := <-c.send:
			if !ok {
				return
			}

			c.UpdateLastActive()
			if err := c.writeEvent(event); err != nil {
				log.Errorf("SSE write event failed: clientID=%s, error=%v", c.ID, err)
				return
			}

			// 记录最后的事件ID
			if event.ID != "" {
				c.lastEventID = event.ID
			}

		case <-heartbeatTicker.C:
			// 发送心跳
			if err := c.writeHeartbeat(); err != nil {
				log.Errorf("SSE heartbeat failed: clientID=%s, error=%v", c.ID, err)
				return
			}
			c.UpdateLastActive()
		}
	}
}

// writeEvent 写入事件
func (c *Client) writeEvent(event *Event) error {
	_, err := fmt.Fprint(c.Writer, event.Format())
	if err != nil {
		return err
	}

	if flusher, ok := c.Writer.(http.Flusher); ok {
		flusher.Flush()
	}

	return nil
}

// writeHeartbeat 写入心跳
func (c *Client) writeHeartbeat() error {
	comment := NewCommentEvent(fmt.Sprintf("heartbeat %d", time.Now().Unix()))
	return c.writeEvent(comment)
}

// ========== 全局管理器实例 ==========

var SSEManager = NewManager()

// InitSSE 初始化 SSE 服务
func InitSSE() {
	log.Info("Initializing SSE service...")
	go SSEManager.Start()
	log.Info("SSE service initialized successfully")
}
