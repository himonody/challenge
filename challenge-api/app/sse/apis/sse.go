package apis

import (
	"challenge/app/sse/service"
	coreSSE "challenge/core/sse"
	"challenge/core/utils/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SSEApis SSE API 接口
type SSEApis struct {
	Service *service.SSEService
}

// Connect SSE 连接端点
func (a *SSEApis) Connect(c *gin.Context) {
	// 使用核心 SSE 管理器处理连接
	coreSSE.SSEManager.ServeHTTP(c)
}

// SendMessage 发送消息接口（管理后台使用）
func (a *SSEApis) SendMessage(c *gin.Context) {
	var req struct {
		UserID    string                 `json:"user_id" binding:"required"`
		EventType string                 `json:"event_type" binding:"required"`
		Data      map[string]interface{} `json:"data" binding:"required"`
		Priority  int                    `json:"priority"`
		Persist   bool                   `json:"persist"`
		TTL       int                    `json:"ttl"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}

	// 构建选项
	var options []service.MessageOption
	if req.Priority > 0 {
		options = append(options, service.WithPriority(req.Priority))
	}
	if req.Persist {
		options = append(options, service.WithPersist(true))
	}
	if req.TTL > 0 {
		options = append(options, service.WithTTL(req.TTL))
	}

	// 发送消息
	if err := a.Service.SendToUser(req.UserID, req.EventType, req.Data, options...); err != nil {
		log.Errorf("Failed to send message: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "Failed to send message",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Message sent successfully",
	})
}

// SendGroupMessage 发送组消息接口
func (a *SSEApis) SendGroupMessage(c *gin.Context) {
	var req struct {
		Group     string                 `json:"group" binding:"required"`
		EventType string                 `json:"event_type" binding:"required"`
		Data      map[string]interface{} `json:"data" binding:"required"`
		Priority  int                    `json:"priority"`
		Persist   bool                   `json:"persist"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}

	// 构建选项
	var options []service.MessageOption
	if req.Priority > 0 {
		options = append(options, service.WithPriority(req.Priority))
	}
	if req.Persist {
		options = append(options, service.WithPersist(true))
	}

	// 发送消息
	if err := a.Service.SendToGroup(req.Group, req.EventType, req.Data, options...); err != nil {
		log.Errorf("Failed to send group message: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "Failed to send group message",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Group message sent successfully",
	})
}

// Broadcast 广播消息接口
func (a *SSEApis) Broadcast(c *gin.Context) {
	var req struct {
		EventType string                 `json:"event_type" binding:"required"`
		Data      map[string]interface{} `json:"data" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}

	// 广播消息
	if err := a.Service.Broadcast(req.EventType, req.Data); err != nil {
		log.Errorf("Failed to broadcast message: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "Failed to broadcast message",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Broadcast sent successfully",
		"data": map[string]interface{}{
			"clientCount": coreSSE.SSEManager.GetClientCount(),
		},
	})
}

// GetPendingMessages 获取待发送消息（重连恢复）
func (a *SSEApis) GetPendingMessages(c *gin.Context) {
	var req struct {
		UserID      string `json:"user_id" binding:"required"`
		LastEventID string `json:"last_event_id"`
		Limit       int    `json:"limit"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}

	limit := req.Limit
	if limit <= 0 || limit > 100 {
		limit = 50
	}

	messages, err := a.Service.GetPendingMessages(req.UserID, req.LastEventID, limit)
	if err != nil {
		log.Errorf("Failed to get pending messages: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "Failed to get pending messages",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": messages,
	})
}

// MarkRead 标记消息已读
func (a *SSEApis) MarkRead(c *gin.Context) {
	var req struct {
		EventID string `json:"event_id" binding:"required"`
		UserID  string `json:"user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}

	if err := a.Service.MarkMessageRead(req.EventID, req.UserID); err != nil {
		log.Errorf("Failed to mark message read: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "Failed to mark message read",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Message marked as read",
	})
}

// GetUnreadCount 获取未读消息数量
func (a *SSEApis) GetUnreadCount(c *gin.Context) {
	var req struct {
		UserID string `json:"user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}

	count, err := a.Service.GetUnreadCount(req.UserID)
	if err != nil {
		log.Errorf("Failed to get unread count: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "Failed to get unread count",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": map[string]interface{}{
			"user_id":      req.UserID,
			"unread_count": count,
		},
	})
}

// Subscribe 订阅分组
func (a *SSEApis) Subscribe(c *gin.Context) {
	var req struct {
		UserID     string   `json:"user_id" binding:"required"`
		Group      string   `json:"group" binding:"required"`
		EventTypes []string `json:"event_types"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}

	if err := a.Service.Subscribe(req.UserID, req.Group, req.EventTypes); err != nil {
		log.Errorf("Failed to subscribe: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "Failed to subscribe",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Subscribed successfully",
	})
}

// Unsubscribe 取消订阅
func (a *SSEApis) Unsubscribe(c *gin.Context) {
	var req struct {
		UserID string `json:"user_id" binding:"required"`
		Group  string `json:"group" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}

	if err := a.Service.Unsubscribe(req.UserID, req.Group); err != nil {
		log.Errorf("Failed to unsubscribe: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "Failed to unsubscribe",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Unsubscribed successfully",
	})
}

// GetSubscriptions 获取订阅列表
func (a *SSEApis) GetSubscriptions(c *gin.Context) {
	var req struct {
		UserID string `json:"user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}

	subs, err := a.Service.GetUserSubscriptions(req.UserID)
	if err != nil {
		log.Errorf("Failed to get subscriptions: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "Failed to get subscriptions",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": subs,
	})
}

// GetInfo 获取 SSE 管理器信息
func (a *SSEApis) GetInfo(c *gin.Context) {
	coreSSE.SSEManager.GetInfo(c)
}

// GetGroupInfo 获取分组信息
func (a *SSEApis) GetGroupInfo(c *gin.Context) {
	coreSSE.SSEManager.GetGroupInfo(c)
}

// Disconnect 断开客户端连接
func (a *SSEApis) Disconnect(c *gin.Context) {
	var req struct {
		ClientID string `json:"client_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}

	if client, ok := coreSSE.SSEManager.GetClient(req.ClientID); ok {
		client.Close()
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "Client disconnected successfully",
			"data": map[string]string{
				"client_id": req.ClientID,
			},
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"msg":  "Client not found",
		})
	}
}
