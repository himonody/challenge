# SSE（Server-Sent Events）服务端推送模块

## 📖 简介

SSE 是一种服务器向客户端推送数据的轻量级技术，基于 HTTP 协议，适用于单向实时数据推送场景。

### ✨ 特性

- ✅ **单向推送**：服务器到客户端的实时数据推送
- ✅ **自动重连**：客户端断开后自动重连
- ✅ **轻量级**：基于 HTTP，无需额外协议
- ✅ **分组管理**：支持客户端分组
- ✅ **多种发送模式**：单播、组播、广播
- ✅ **心跳机制**：自动保持连接活跃
- ✅ **连接管理**：自动清理超时连接
- ✅ **事件类型**：支持自定义事件类型
- ✅ **JSON 支持**：自动序列化 JSON 数据

### 🆚 SSE vs WebSocket

| 特性 | SSE | WebSocket |
|------|-----|-----------|
| 通信方向 | 单向（服务器→客户端） | 双向 |
| 协议 | HTTP | WebSocket |
| 浏览器支持 | 广泛支持（IE不支持） | 广泛支持 |
| 重连机制 | 浏览器自动重连 | 需手动实现 |
| 数据格式 | 文本 | 文本/二进制 |
| 复杂度 | 简单 | 复杂 |
| 适用场景 | 实时通知、进度推送 | 聊天、游戏 |

---

## 🚀 快速开始

### 1. 初始化 SSE 管理器

```go
package main

import (
	"challenge/core/sse"
)

func main() {
	// 启动 SSE 管理器（在应用启动时调用一次）
	go sse.SSEManager.Start()           // 管理器主循环
	go sse.SSEManager.SendService()     // 单播服务
	go sse.SSEManager.SendGroupService() // 组播服务
	go sse.SSEManager.SendAllService()  // 广播服务
	
	// ... 启动 HTTP 服务器
}
```

### 2. 注册 SSE 路由

```go
package router

import (
	"challenge/core/sse"
	"github.com/gin-gonic/gin"
)

func SetupSSERoutes(r *gin.Engine) {
	sseGroup := r.Group("/api/v1/sse")
	{
		// SSE 连接端点
		// 路径参数方式：/api/v1/sse/stream/:group/:id
		sseGroup.GET("/stream/:group/:id", sse.SSEManager.SSEHandler)
		
		// 查询参数方式：/api/v1/sse/stream?group=notifications&id=user123
		sseGroup.GET("/stream", sse.SSEManager.SSEHandler)
		
		// 断开连接接口
		sseGroup.POST("/disconnect/:group/:id", sse.SSEManager.DisconnectHandler)
		
		// 管理器信息
		sseGroup.GET("/info", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"code": 200,
				"data": sse.SSEManager.Info(),
			})
		})
	}
}
```

### 3. 客户端连接

#### JavaScript 客户端

```javascript
// 基本连接
const eventSource = new EventSource('/api/v1/sse/stream/notifications/user123');

// 监听连接事件
eventSource.addEventListener('connected', (e) => {
  const data = JSON.parse(e.data);
  console.log('Connected:', data);
});

// 监听默认消息
eventSource.onmessage = (e) => {
  const data = JSON.parse(e.data);
  console.log('Message:', data);
};

// 监听自定义事件
eventSource.addEventListener('notification', (e) => {
  const data = JSON.parse(e.data);
  console.log('Notification:', data);
  // 显示通知
  showNotification(data);
});

// 错误处理
eventSource.onerror = (err) => {
  console.error('SSE Error:', err);
  if (eventSource.readyState === EventSource.CLOSED) {
    console.log('Connection closed');
  }
};

// 关闭连接
// eventSource.close();
```

#### 带查询参数的连接

```javascript
const url = new URL('/api/v1/sse/stream', window.location.origin);
url.searchParams.set('group', 'notifications');
url.searchParams.set('id', 'user123');

const eventSource = new EventSource(url);
```

#### 带认证的连接（需要自定义）

SSE 原生不支持自定义请求头，但可以通过 URL 参数传递 token：

```javascript
const token = localStorage.getItem('token');
const url = `/api/v1/sse/stream/notifications/user123?token=${token}`;
const eventSource = new EventSource(url);
```

---

## 📡 服务端使用

### 1. 发送消息到单个客户端

```go
import (
	"challenge/core/sse"
	"context"
)

func SendToUser(userId string, message interface{}) {
	// 简单消息
	event := sse.NewEvent(message)
	sse.SSEManager.Send(context.Background(), userId, "notifications", event)
}

// 带事件类型的消息
func SendTypedNotification(userId string, notification map[string]interface{}) {
	event := sse.NewTypedEvent("notification", notification)
	sse.SSEManager.Send(context.Background(), userId, "notifications", event)
}

// 带ID的消息（支持重连恢复）
func SendWithID(userId string, eventId string, data interface{}) {
	event := sse.NewEventWithID(eventId, data)
	sse.SSEManager.Send(context.Background(), userId, "notifications", event)
}
```

### 2. 发送消息到组（组播）

```go
func NotifyGroup(group string, message interface{}) {
	event := sse.NewTypedEvent("group_message", message)
	sse.SSEManager.SendGroup(group, event)
}

// 示例：通知所有在线管理员
func NotifyAdmins(message string) {
	event := sse.NewTypedEvent("admin_alert", map[string]interface{}{
		"message":   message,
		"timestamp": time.Now().Unix(),
		"level":     "warning",
	})
	sse.SSEManager.SendGroup("admins", event)
}
```

### 3. 发送消息到所有客户端（广播）

```go
func BroadcastSystemMessage(message string) {
	event := sse.NewTypedEvent("system", map[string]interface{}{
		"type":    "maintenance",
		"message": message,
		"time":    time.Now().Format(time.RFC3339),
	})
	sse.SSEManager.SendAll(event)
}
```

### 4. 在 HTTP 接口中使用

```go
package apis

import (
	"challenge/core/sse"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 发送通知接口
func SendNotification(c *gin.Context) {
	var req struct {
		UserId  string      `json:"user_id" binding:"required"`
		Type    string      `json:"type"`
		Message interface{} `json:"message" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 发送 SSE 消息
	event := sse.NewTypedEvent(req.Type, req.Message)
	sse.SSEManager.Send(c.Request.Context(), req.UserId, "notifications", event)

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Notification sent",
	})
}

// 获取 SSE 管理器状态
func GetSSEStatus(c *gin.Context) {
	info := sse.SSEManager.Info()
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": info,
	})
}

// 断开指定客户端连接
func DisconnectClient(c *gin.Context) {
	clientId := c.Param("id")
	group := c.DefaultQuery("group", "default")

	// 这会触发客户端的 onerror 和自动重连
	sse.SSEManager.UnRegisterClient(&sse.Client{
		Id:    clientId,
		Group: group,
	})

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Client disconnected",
	})
}
```

---

## 🎯 实际应用场景

### 1. 实时通知系统

```go
// 用户收到新消息通知
func NotifyNewMessage(userId string, messageId uint64) {
	event := sse.NewTypedEvent("new_message", map[string]interface{}{
		"message_id": messageId,
		"timestamp":  time.Now().Unix(),
		"unread":     true,
	})
	sse.SSEManager.Send(context.Background(), userId, "notifications", event)
}

// 系统通知
func SendSystemNotification(userId string, title, content string) {
	event := sse.NewTypedEvent("notification", map[string]interface{}{
		"title":     title,
		"content":   content,
		"type":      "system",
		"timestamp": time.Now().Unix(),
	})
	sse.SSEManager.Send(context.Background(), userId, "notifications", event)
}
```

### 2. 实时进度推送

```go
// 文件上传进度
func NotifyUploadProgress(userId string, fileId string, progress int) {
	event := sse.NewTypedEvent("upload_progress", map[string]interface{}{
		"file_id":  fileId,
		"progress": progress,
		"status":   "uploading",
	})
	sse.SSEManager.Send(context.Background(), userId, "uploads", event)
}

// 任务处理进度
func NotifyTaskProgress(userId string, taskId string, progress int, message string) {
	event := sse.NewTypedEvent("task_progress", map[string]interface{}{
		"task_id":  taskId,
		"progress": progress,
		"message":  message,
		"time":     time.Now().Format(time.RFC3339),
	})
	sse.SSEManager.Send(context.Background(), userId, "tasks", event)
}
```

### 3. 实时数据更新

```go
// 挑战数据更新
func NotifyChallengeUpdate(challengeId uint64, data map[string]interface{}) {
	event := sse.NewTypedEvent("challenge_update", map[string]interface{}{
		"challenge_id": challengeId,
		"data":         data,
		"timestamp":    time.Now().Unix(),
	})
	
	// 向所有关注这个挑战的用户推送
	group := fmt.Sprintf("challenge_%d", challengeId)
	sse.SSEManager.SendGroup(group, event)
}

// 排行榜更新
func NotifyRankUpdate(userId string, rank int, score int64) {
	event := sse.NewTypedEvent("rank_update", map[string]interface{}{
		"rank":      rank,
		"score":     score,
		"timestamp": time.Now().Unix(),
	})
	sse.SSEManager.Send(context.Background(), userId, "rankings", event)
}
```

### 4. 在线状态监控

```go
// 通知用户上线
func NotifyUserOnline(userId string, friendIds []string) {
	event := sse.NewTypedEvent("user_online", map[string]interface{}{
		"user_id":   userId,
		"timestamp": time.Now().Unix(),
	})

	// 通知所有好友
	for _, friendId := range friendIds {
		sse.SSEManager.Send(context.Background(), friendId, "presence", event)
	}
}
```

---

## 🔧 高级用法

### 1. 自定义事件类型

```go
// 创建不同类型的事件
func SendMultiTypeEvents(userId string) {
	// 普通消息
	event1 := sse.NewEvent("Hello World")
	sse.SSEManager.Send(context.Background(), userId, "notifications", event1)

	// 带类型的事件
	event2 := sse.NewTypedEvent("alert", map[string]interface{}{
		"level":   "warning",
		"message": "System maintenance in 10 minutes",
	})
	sse.SSEManager.Send(context.Background(), userId, "notifications", event2)

	// 带ID的事件（用于重连恢复）
	event3 := sse.NewEventWithID("msg_001", map[string]interface{}{
		"content": "Important message",
	})
	sse.SSEManager.Send(context.Background(), userId, "notifications", event3)
}
```

### 2. 批量通知

```go
func NotifyMultipleUsers(userIds []string, message interface{}) {
	event := sse.NewTypedEvent("notification", message)
	
	for _, userId := range userIds {
		sse.SSEManager.Send(context.Background(), userId, "notifications", event)
	}
}
```

### 3. 条件推送

```go
func NotifyByCondition(userIds []string, condition func(string) bool, message interface{}) {
	event := sse.NewTypedEvent("conditional", message)
	
	for _, userId := range userIds {
		if condition(userId) {
			sse.SSEManager.Send(context.Background(), userId, "notifications", event)
		}
	}
}
```

### 4. 获取连接信息

```go
func GetConnectionStats(c *gin.Context) {
	info := sse.SSEManager.Info()
	
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": info,
		"details": map[string]interface{}{
			"total_groups":  info["groupCount"],
			"total_clients": info["clientCount"],
			"channels": map[string]interface{}{
				"register":    info["registerChannelLen"],
				"unregister":  info["unregisterChannelLen"],
				"message":     info["messageChannelLen"],
				"group":       info["groupMessageChannelLen"],
				"broadcast":   info["broadcastMessageChannelLen"],
			},
		},
	})
}

// 获取特定组的连接数
func GetGroupStats(group string) int {
	return sse.SSEManager.GetGroupClients(group)
}
```

---

## 🌐 客户端完整示例

### React 示例

```jsx
import { useEffect, useState } from 'react';

function useSSE(url) {
  const [data, setData] = useState(null);
  const [error, setError] = useState(null);
  const [isConnected, setIsConnected] = useState(false);

  useEffect(() => {
    const eventSource = new EventSource(url);

    eventSource.addEventListener('connected', (e) => {
      console.log('SSE Connected:', JSON.parse(e.data));
      setIsConnected(true);
    });

    eventSource.addEventListener('notification', (e) => {
      const data = JSON.parse(e.data);
      setData(data);
      // 显示通知
      if (Notification.permission === 'granted') {
        new Notification(data.title, {
          body: data.content,
          icon: '/icon.png'
        });
      }
    });

    eventSource.onerror = (err) => {
      console.error('SSE Error:', err);
      setError(err);
      setIsConnected(false);
    };

    return () => {
      eventSource.close();
      setIsConnected(false);
    };
  }, [url]);

  return { data, error, isConnected };
}

// 使用
function NotificationPanel() {
  const userId = 'user123';
  const { data, isConnected } = useSSE(`/api/v1/sse/stream/notifications/${userId}`);

  return (
    <div>
      <div>Status: {isConnected ? '🟢 Connected' : '🔴 Disconnected'}</div>
      {data && (
        <div className="notification">
          <h3>{data.title}</h3>
          <p>{data.content}</p>
        </div>
      )}
    </div>
  );
}
```

### Vue 示例

```vue
<template>
  <div>
    <div>状态: {{ isConnected ? '🟢 已连接' : '🔴 未连接' }}</div>
    <div v-for="notification in notifications" :key="notification.id">
      <h3>{{ notification.title }}</h3>
      <p>{{ notification.content }}</p>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      eventSource: null,
      isConnected: false,
      notifications: []
    };
  },
  mounted() {
    this.connectSSE();
  },
  beforeUnmount() {
    if (this.eventSource) {
      this.eventSource.close();
    }
  },
  methods: {
    connectSSE() {
      const userId = this.$store.state.user.id;
      this.eventSource = new EventSource(`/api/v1/sse/stream/notifications/${userId}`);

      this.eventSource.addEventListener('connected', (e) => {
        console.log('SSE 已连接:', JSON.parse(e.data));
        this.isConnected = true;
      });

      this.eventSource.addEventListener('notification', (e) => {
        const data = JSON.parse(e.data);
        this.notifications.unshift(data);
      });

      this.eventSource.onerror = (err) => {
        console.error('SSE 错误:', err);
        this.isConnected = false;
      };
    }
  }
};
</script>
```

---

## ⚙️ 配置和优化

### 1. 心跳和超时配置

在 `sse.go` 中可以调整以下参数：

```go
// 心跳间隔（默认30秒）
heartbeatTicker := time.NewTicker(30 * time.Second)

// 超时检查间隔（默认5秒检查一次）
timeoutTicker := time.NewTicker(5 * time.Second)

// 超时时间（默认90秒无活动则断开）
if time.Since(c.LastActive) > 90*time.Second {
	// 断开连接
}
```

### 2. Channel 缓冲区大小

```go
var SSEManager = Manager{
	Register:         make(chan *Client, 128),      // 注册通道
	UnRegister:       make(chan *Client, 128),      // 注销通道
	Message:          make(chan *MessageData, 256), // 消息通道
	GroupMessage:     make(chan *GroupMessageData, 256),
	BroadCastMessage: make(chan *BroadCastMessageData, 256),
}

// 客户端消息通道
Message: make(chan *Event, 100), // 每个客户端的消息缓冲
```

### 3. Nginx 配置

如果使用 Nginx 反向代理，需要禁用缓冲：

```nginx
location /api/v1/sse/ {
    proxy_pass http://backend;
    proxy_http_version 1.1;
    proxy_set_header Connection '';
    proxy_buffering off;
    proxy_cache off;
    proxy_read_timeout 86400s;
    chunked_transfer_encoding off;
}
```

---

## 🐛 故障排查

### 1. 连接无法建立

- 检查响应头是否正确设置了 `Content-Type: text/event-stream`
- 确认 Nginx 已禁用缓冲（`X-Accel-Buffering: no`）
- 检查是否支持 `http.Flusher`

### 2. 消息收不到

- 检查客户端 ID 和分组是否正确
- 查看服务端日志确认消息已发送
- 确认客户端事件监听器名称与服务端事件类型匹配

### 3. 频繁断线重连

- 调整心跳间隔和超时时间
- 检查网络稳定性
- 查看服务器日志确认断开原因

### 4. 内存占用过高

- 减少每个客户端的消息缓冲区大小
- 及时清理不活跃的连接
- 使用分组管理避免广播消息

---

## 📊 性能建议

1. **合理使用分组**：避免不必要的全局广播
2. **控制消息频率**：避免短时间内发送大量消息
3. **设置合理的超时**：及时清理僵尸连接
4. **使用连接池**：对于高并发场景，考虑使用多个管理器实例
5. **监控连接数**：定期检查 `SSEManager.Info()` 避免连接泄漏

---

## 📝 总结

SSE 是一个轻量级、易用的服务端推送方案，特别适合：

✅ 实时通知
✅ 进度推送
✅ 数据更新
✅ 系统广播

相比 WebSocket，SSE 更简单，但只支持单向通信。选择合适的技术取决于具体的业务场景！
