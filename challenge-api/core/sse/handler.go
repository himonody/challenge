package sse

import (
	"challenge/core/utils/log"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ServeHTTP SSE HTTP 处理器
func (m *Manager) ServeHTTP(c *gin.Context) {
	// 设置 SSE 响应头
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("X-Accel-Buffering", "no") // 禁用 Nginx 缓冲

	// 检查是否支持 Flusher
	if _, ok := c.Writer.(http.Flusher); !ok {
		log.Error("SSE not supported: ResponseWriter does not implement http.Flusher")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "SSE not supported",
		})
		return
	}

	// 获取客户端参数
	clientID := getClientID(c)
	userID := getUserID(c)
	group := getGroup(c)
	lastEventID := c.Request.Header.Get("Last-Event-ID")

	// 创建客户端
	client := &Client{
		ID:           clientID,
		UserID:       userID,
		Group:        group,
		Writer:       c.Writer,
		Request:      c.Request,
		send:         make(chan *Event, 100),
		closeChan:    make(chan struct{}),
		lastEventID:  lastEventID,
		connectedAt:  time.Now(),
		lastActiveAt: time.Now(),
		metadata:     make(map[string]string),
	}

	// 收集客户端元数据
	client.SetMetadata("ip", c.ClientIP())
	client.SetMetadata("user_agent", c.Request.UserAgent())
	client.SetMetadata("referer", c.Request.Referer())

	// 注册客户端
	m.register <- client

	// 启动写入泵
	client.writePump(m)

	log.Infof("SSE connection closed: clientID=%s, duration=%v",
		clientID, time.Since(client.connectedAt))
}

// DisconnectClient 断开指定客户端
func (m *Manager) DisconnectClient(c *gin.Context) {
	clientID := c.Param("id")
	if clientID == "" {
		clientID = c.Query("id")
	}

	if clientID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "clientID is required",
		})
		return
	}

	if client, ok := m.GetClient(clientID); ok {
		m.unregister <- client
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "Client disconnected successfully",
			"data": map[string]string{
				"clientId": clientID,
			},
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"msg":  "Client not found",
		})
	}
}

// GetInfo 获取管理器信息
func (m *Manager) GetInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": m.Info(),
	})
}

// GetGroupInfo 获取分组信息
func (m *Manager) GetGroupInfo(c *gin.Context) {
	group := c.Param("group")
	if group == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "group is required",
		})
		return
	}

	clientIDs := m.GetClientsByGroup(group)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": map[string]interface{}{
			"group":       group,
			"clientCount": len(clientIDs),
			"clients":     clientIDs,
		},
	})
}

// ========== 辅助函数 ==========

// getClientID 获取客户端ID
func getClientID(c *gin.Context) string {
	// 优先级：路径参数 > 查询参数 > 自动生成
	clientID := c.Param("id")
	if clientID == "" {
		clientID = c.Query("id")
	}
	if clientID == "" {
		clientID = c.Query("client_id")
	}
	if clientID == "" {
		// 自动生成唯一ID
		clientID = fmt.Sprintf("client_%d", time.Now().UnixNano())
	}
	return clientID
}

// getUserID 获取用户ID
func getUserID(c *gin.Context) string {
	// 优先级：JWT token > 查询参数
	if userID, exists := c.Get("user_id"); exists {
		if uid, ok := userID.(string); ok {
			return uid
		}
		if uid, ok := userID.(int64); ok {
			return fmt.Sprintf("%d", uid)
		}
	}

	userID := c.Query("user_id")
	if userID == "" {
		userID = c.Query("userId")
	}

	return userID
}

// getGroup 获取分组
func getGroup(c *gin.Context) string {
	// 优先级：路径参数 > 查询参数 > 默认值
	group := c.Param("group")
	if group == "" {
		group = c.Query("group")
	}
	if group == "" {
		group = "default"
	}
	return group
}
