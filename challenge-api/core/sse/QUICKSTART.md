# SSE 快速开始指南

## 🚀 5分钟快速集成

### 步骤 1: 在应用启动时初始化 SSE 服务

编辑 `core/cmd/api/server.go` 或你的主入口文件：

```go
import (
    "challenge/core/sse"
    // ... 其他导入
)

func setup() {
    // ... 现有的初始化代码 ...
    
    // 初始化 SSE 服务
    sse.InitSSEServices()
    
    log.Info("SSE service initialized")
}
```

### 步骤 2: 注册 SSE 路由

在 `app/router.go` 或路由配置文件中：

```go
import (
    "challenge/core/sse"
    "github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
    api := r.Group("/api/v1")
    {
        // 注册 SSE 路由
        sse.RegisterSSERoutes(api)
        
        // ... 其他路由 ...
    }
}
```

### 步骤 3: 在业务代码中发送消息

```go
import "challenge/core/sse"

// 示例：用户注册成功后发送欢迎通知
func (s *AuthService) Register(username string) error {
    // ... 注册逻辑 ...
    
    // 发送 SSE 通知
    event := sse.NewTypedEvent("notification", map[string]interface{}{
        "title":   "欢迎加入",
        "content": fmt.Sprintf("欢迎 %s 加入我们！", username),
        "type":    "welcome",
    })
    sse.SSEManager.Send(context.Background(), userId, "notifications", event)
    
    return nil
}
```

### 步骤 4: 前端连接

```javascript
// 连接到 SSE
const userId = 'user123';
const eventSource = new EventSource(`/api/v1/sse/stream/notifications/${userId}`);

// 监听通知
eventSource.addEventListener('notification', (e) => {
    const data = JSON.parse(e.data);
    console.log('收到通知:', data);
    // 显示通知
    showToast(data.title, data.content);
});

// 错误处理
eventSource.onerror = (err) => {
    console.error('SSE 连接错误:', err);
};
```

## 🎯 完整示例

### 后端：发送实时通知

```go
package notification

import (
    "challenge/core/sse"
    "context"
    "time"
)

// 发送用户通知
func SendUserNotification(userId string, title, content string) {
    event := sse.NewTypedEvent("notification", map[string]interface{}{
        "title":     title,
        "content":   content,
        "timestamp": time.Now().Unix(),
    })
    sse.SSEManager.Send(context.Background(), userId, "notifications", event)
}

// 发送系统广播
func BroadcastSystemMessage(message string) {
    event := sse.NewTypedEvent("system", map[string]interface{}{
        "message":   message,
        "timestamp": time.Now().Unix(),
    })
    sse.SSEManager.SendAll(event)
}

// 发送组消息
func SendGroupMessage(group string, message string) {
    event := sse.NewTypedEvent("group_message", map[string]interface{}{
        "message":   message,
        "timestamp": time.Now().Unix(),
    })
    sse.SSEManager.SendGroup(group, event)
}
```

### 前端：React Hook 示例

```jsx
import { useEffect, useState } from 'react';

function useSSE(userId) {
    const [notifications, setNotifications] = useState([]);
    const [isConnected, setIsConnected] = useState(false);

    useEffect(() => {
        const url = `/api/v1/sse/stream/notifications/${userId}`;
        const eventSource = new EventSource(url);

        eventSource.addEventListener('connected', () => {
            setIsConnected(true);
            console.log('SSE 已连接');
        });

        eventSource.addEventListener('notification', (e) => {
            const data = JSON.parse(e.data);
            setNotifications(prev => [data, ...prev]);
        });

        eventSource.onerror = () => {
            setIsConnected(false);
        };

        return () => {
            eventSource.close();
        };
    }, [userId]);

    return { notifications, isConnected };
}

// 使用
function NotificationPanel() {
    const { notifications, isConnected } = useSSE('user123');

    return (
        <div>
            <div>状态: {isConnected ? '🟢 已连接' : '🔴 未连接'}</div>
            {notifications.map((notif, index) => (
                <div key={index}>
                    <h3>{notif.title}</h3>
                    <p>{notif.content}</p>
                </div>
            ))}
        </div>
    );
}
```

## 📡 常用 API

### 发送消息

```go
// 1. 发送给单个用户
sse.SSEManager.Send(ctx, "user123", "notifications", event)

// 2. 发送给整个组
sse.SSEManager.SendGroup("admins", event)

// 3. 广播给所有用户
sse.SSEManager.SendAll(event)
```

### 创建事件

```go
// 简单事件
event := sse.NewEvent("Hello World")

// 带类型的事件
event := sse.NewTypedEvent("notification", map[string]interface{}{
    "title": "标题",
    "content": "内容",
})

// 带ID的事件（支持重连恢复）
event := sse.NewEventWithID("msg_001", data)
```

### 管理连接

```go
// 获取管理器信息
info := sse.SSEManager.Info()

// 获取组内客户端数量
count := sse.SSEManager.GetGroupClients("notifications")

// 获取总客户端数
total := sse.SSEManager.LenClient()
```

## 🔧 测试

### 1. 启动服务器

```bash
cd /Users/mac/workspace/challenge/challenge-api
go run main.go server
```

### 2. 打开测试页面

在浏览器中打开：
```
file:///Users/mac/workspace/challenge/challenge-api/core/sse/client_example.html
```

### 3. 测试连接

1. 填写服务器地址（如 `http://localhost:8000`）
2. 填写客户端 ID（如 `user123`）
3. 填写分组（如 `notifications`）
4. 点击「连接」按钮
5. 点击「测试发送」按钮发送测试消息

### 4. 使用 curl 测试

```bash
# 测试发送单个消息
curl -X POST http://localhost:8000/api/v1/sse/test/send \
  -H "Content-Type: application/json" \
  -d '{
    "client_id": "user123",
    "group": "notifications",
    "type": "notification",
    "data": {
      "title": "测试通知",
      "content": "这是一条测试消息"
    }
  }'

# 测试广播
curl -X POST http://localhost:8000/api/v1/sse/test/broadcast \
  -H "Content-Type: application/json" \
  -d '{
    "type": "system",
    "data": {
      "message": "系统维护通知"
    }
  }'

# 查看管理器状态
curl http://localhost:8000/api/v1/sse/info
```

## 📝 实际应用场景

### 1. 用户通知系统

```go
// 新消息通知
func NotifyNewMessage(userId string, messageId uint64) {
    event := sse.NewTypedEvent("new_message", map[string]interface{}{
        "message_id": messageId,
        "timestamp":  time.Now().Unix(),
    })
    sse.SSEManager.Send(context.Background(), userId, "notifications", event)
}
```

### 2. 进度推送

```go
// 文件上传进度
func UpdateUploadProgress(userId, fileId string, progress int) {
    event := sse.NewTypedEvent("upload_progress", map[string]interface{}{
        "file_id":  fileId,
        "progress": progress,
    })
    sse.SSEManager.Send(context.Background(), userId, "uploads", event)
}
```

### 3. 实时数据更新

```go
// 挑战排名更新
func UpdateChallengeRank(challengeId uint64, rankings []interface{}) {
    event := sse.NewTypedEvent("rank_update", map[string]interface{}{
        "challenge_id": challengeId,
        "rankings":     rankings,
    })
    group := fmt.Sprintf("challenge_%d", challengeId)
    sse.SSEManager.SendGroup(group, event)
}
```

## 🎨 最佳实践

### 1. 合理使用分组

```go
// ✅ 好的做法：按业务场景分组
"notifications"  // 通用通知
"chat_room_123"  // 聊天室
"challenge_456"  // 挑战
"admin_panel"    // 管理面板

// ❌ 不好的做法：所有用户都在一个组
"all_users"      // 导致不必要的消息推送
```

### 2. 控制消息频率

```go
// ✅ 好的做法：合并频繁更新
func UpdateProgress(userId string) {
    // 每秒最多更新一次
    throttle(1*time.Second, func() {
        event := sse.NewTypedEvent("progress", data)
        sse.SSEManager.Send(ctx, userId, "tasks", event)
    })
}

// ❌ 不好的做法：频繁推送
for i := 0; i < 1000; i++ {
    sse.SSEManager.Send(ctx, userId, "tasks", event) // 太频繁
}
```

### 3. 错误处理

```go
// 检查客户端是否在线
if sse.SSEManager.GetGroupClients(group) > 0 {
    sse.SSEManager.SendGroup(group, event)
} else {
    // 记录日志或存储消息供后续拉取
    log.Warnf("No clients in group %s", group)
}
```

## 🔍 故障排查

### 问题 1: 连接不上

**检查项：**
- 服务器是否已启动 SSE 服务（`sse.InitSSEServices()`）
- 路由是否正确注册
- 端口是否正确
- 防火墙是否阻止连接

### 问题 2: 收不到消息

**检查项：**
- 客户端 ID 和分组是否匹配
- 服务端是否调用了发送方法
- 查看服务端日志确认消息已发送
- 客户端事件监听器名称是否正确

### 问题 3: 频繁断线

**检查项：**
- 网络是否稳定
- Nginx 是否配置正确（禁用缓冲）
- 调整心跳和超时参数

## 📚 更多文档

- [完整文档](./README.md) - 详细的功能说明和 API 文档
- [示例代码](./example.go) - 更多业务场景示例
- [测试页面](./client_example.html) - 可视化测试工具

---

**恭喜！** 🎉 您已经成功集成了 SSE 服务端推送功能！
