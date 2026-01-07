# SSE 模块 - 实时推送服务

## 📖 模块说明

SSE (Server-Sent Events) 实时推送模块，提供服务器到客户端的单向实时数据推送能力。

### ✨ 核心特性

- ✅ **实时推送**：服务器主动推送数据到客户端
- ✅ **分组管理**：支持按业务场景分组推送
- ✅ **用户维度**：支持向指定用户的所有连接推送
- ✅ **消息持久化**：支持消息持久化和离线推送
- ✅ **重连恢复**：基于 Last-Event-ID 的重连恢复机制
- ✅ **优先级队列**：支持消息优先级
- ✅ **订阅机制**：用户可订阅/取消订阅分组
- ✅ **未读计数**：自动维护未读消息计数
- ✅ **自动清理**：定时清理过期和旧消息

---

## 🏗️ 模块结构

```
app/sse/
├── models/                    # 数据模型
│   ├── sse_message.go         # 消息记录表
│   └── sse_subscription.go    # 订阅关系表
├── repo/                      # 数据库操作
│   ├── sse_message.go         # 消息CRUD
│   └── sse_subscription.go    # 订阅CRUD
├── storage/                   # Redis缓存
│   └── sse_cache.go           # 在线状态、未读计数等
├── service/                   # 业务服务
│   └── sse.go                 # SSE核心业务逻辑
├── apis/                      # HTTP接口
│   └── sse.go                 # RESTful API
├── router/                    # 路由配置
│   └── sse.go                 # 路由注册
├── init.go                    # 模块初始化
└── README.md                  # 本文档
```

---

## 🚀 快速开始

### 1. 初始化

在应用启动时调用模块初始化（已在 `app/init.go` 中集成）：

```go
import "challenge/app"

func main() {
    // 初始化所有模块
    app.InitModules()
    
    // ... 其他初始化代码
}
```

### 2. 客户端连接

#### 方式一：JavaScript 原生

```javascript
// 连接到 SSE
const userId = '123';
const group = 'notifications';
const eventSource = new EventSource(`/api/v1/sse/stream/${group}/${userId}`);

// 连接成功
eventSource.addEventListener('connected', (e) => {
    const data = JSON.parse(e.data);
    console.log('Connected:', data);
});

// 监听自定义事件
eventSource.addEventListener('notification', (e) => {
    const data = JSON.parse(e.data);
    console.log('Notification:', data);
    showNotification(data);
});

// 错误处理
eventSource.onerror = (err) => {
    console.error('SSE Error:', err);
};
```

#### 方式二：React Hook

```jsx
import { useEffect, useState } from 'react';

function useSSE(userId, group) {
    const [data, setData] = useState(null);
    const [isConnected, setIsConnected] = useState(false);

    useEffect(() => {
        const url = `/api/v1/sse/stream/${group}/${userId}`;
        const eventSource = new EventSource(url);

        eventSource.addEventListener('connected', () => {
            setIsConnected(true);
        });

        eventSource.addEventListener('notification', (e) => {
            setData(JSON.parse(e.data));
        });

        eventSource.onerror = () => {
            setIsConnected(false);
        };

        return () => eventSource.close();
    }, [userId, group]);

    return { data, isConnected };
}
```

### 3. 服务端推送

```go
import (
    "challenge/app/sse/service"
    "challenge/core/runtime"
)

// 创建服务
orm := runtime.RuntimeConfig.GetDbByKey("*")
run := runtime.RuntimeConfig
sseService := service.NewSSEService(orm, run, "zh-CN")

// 发送消息给指定用户
sseService.SendToUser("user123", "notification", map[string]interface{}{
    "title":   "新消息",
    "content": "您有一条新消息",
    "time":    time.Now().Unix(),
})

// 发送消息到分组
sseService.SendToGroup("admins", "alert", map[string]interface{}{
    "message": "系统告警",
    "level":   "warning",
})

// 广播给所有在线用户
sseService.Broadcast("system", map[string]interface{}{
    "message": "系统维护通知",
})
```

---

## 📡 API 接口

### SSE 连接端点

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/api/v1/sse/stream` | 连接（查询参数：id, user_id, group） |
| GET | `/api/v1/sse/stream/:group/:id` | 连接（路径参数） |
| GET | `/api/v1/sse/stream/:id` | 连接（简化方式） |

**查询参数：**
- `id`: 客户端ID（可选，不提供则自动生成）
- `user_id`: 用户ID（可选，业务层关联）
- `group`: 分组名称（默认：default）

**请求头：**
- `Last-Event-ID`: 最后接收的事件ID（用于重连恢复）

### 消息发送接口

#### 发送给指定用户

```http
POST /api/v1/sse/send
Content-Type: application/json

{
  "user_id": "123",
  "event_type": "notification",
  "data": {
    "title": "标题",
    "content": "内容"
  },
  "priority": 1,
  "persist": true,
  "ttl": 3600
}
```

#### 发送到分组

```http
POST /api/v1/sse/send/group
Content-Type: application/json

{
  "group": "admins",
  "event_type": "alert",
  "data": {
    "message": "系统告警"
  },
  "priority": 2,
  "persist": false
}
```

#### 广播给所有用户

```http
POST /api/v1/sse/broadcast
Content-Type: application/json

{
  "event_type": "system",
  "data": {
    "message": "系统维护通知"
  }
}
```

### 消息管理接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/api/v1/sse/messages/pending` | 获取待发送消息（重连恢复） |
| POST | `/api/v1/sse/messages/read` | 标记消息已读 |
| GET | `/api/v1/sse/messages/unread` | 获取未读消息数量 |

### 订阅管理接口

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/api/v1/sse/subscribe` | 订阅分组 |
| POST | `/api/v1/sse/unsubscribe` | 取消订阅 |
| GET | `/api/v1/sse/subscriptions` | 获取订阅列表 |

### 管理接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/api/v1/sse/info` | 获取管理器信息 |
| GET | `/api/v1/sse/group/:group` | 获取分组信息 |
| POST | `/api/v1/sse/disconnect/:id` | 断开客户端连接 |

---

## 💾 数据库表

### app_sse_message - 消息记录表

用于消息持久化、离线推送和重连恢复。

| 字段 | 类型 | 说明 |
|------|------|------|
| id | bigint | 主键ID |
| event_id | varchar(64) | 事件ID（唯一） |
| event_type | varchar(64) | 事件类型 |
| receiver_id | varchar(64) | 接收者ID |
| receiver_type | varchar(20) | 接收者类型（user/client/group） |
| group_name | varchar(64) | 分组名称 |
| priority | tinyint | 优先级（0-普通 1-高 2-紧急） |
| data | text | 消息数据（JSON） |
| status | tinyint | 状态（0-待发送 1-已发送 2-已读） |
| expire_at | datetime | 过期时间 |
| created_at | datetime | 创建时间 |
| updated_at | datetime | 更新时间 |

**索引：**
- `idx_event_id` (event_id) - 唯一索引
- `idx_event_type` (event_type)
- `idx_receiver_created` (receiver_id, created_at)
- `idx_group_created` (group_name, created_at)
- `idx_expire_at` (expire_at)

### app_sse_subscription - 订阅关系表

用于管理用户对分组的订阅关系。

| 字段 | 类型 | 说明 |
|------|------|------|
| id | bigint | 主键ID |
| user_id | varchar(64) | 用户ID |
| group_name | varchar(64) | 订阅组名 |
| event_types | varchar(255) | 订阅的事件类型（逗号分隔） |
| status | tinyint | 状态（0-禁用 1-启用） |
| created_at | datetime | 创建时间 |
| updated_at | datetime | 更新时间 |

**索引：**
- `idx_user_group` (user_id, group_name) - 唯一索引
- `idx_group_status` (group_name, status)

---

## 🎯 业务场景示例

### 1. 实时通知系统

```go
// 用户收到新消息
func NotifyNewMessage(userID string, messageID uint64) {
    sseService.SendToUser(userID, "new_message", map[string]interface{}{
        "message_id": messageID,
        "timestamp":  time.Now().Unix(),
    }, service.WithPriority(models.PriorityHigh))
}

// 系统通知
func SendSystemNotification(userID string, title, content string) {
    sseService.SendToUser(userID, "notification", map[string]interface{}{
        "title":     title,
        "content":   content,
        "type":      "system",
        "timestamp": time.Now().Unix(),
    }, service.WithPersist(true), service.WithTTL(86400))
}
```

### 2. 进度推送

```go
// 文件上传进度
func NotifyUploadProgress(userID, fileID string, progress int) {
    sseService.SendToUser(userID, "upload_progress", map[string]interface{}{
        "file_id":  fileID,
        "progress": progress,
        "status":   "uploading",
    })
}

// 任务处理进度
func NotifyTaskProgress(userID, taskID string, current, total int) {
    progress := (current * 100) / total
    sseService.SendToUser(userID, "task_progress", map[string]interface{}{
        "task_id":  taskID,
        "current":  current,
        "total":    total,
        "progress": progress,
    })
}
```

### 3. 实时数据更新

```go
// 挑战排名更新
func NotifyChallengeUpdate(challengeID uint64, data map[string]interface{}) {
    group := fmt.Sprintf("challenge_%d", challengeID)
    sseService.SendToGroup(group, "challenge_update", map[string]interface{}{
        "challenge_id": challengeID,
        "data":         data,
        "timestamp":    time.Now().Unix(),
    })
}

// 在线状态变更
func NotifyUserStatusChange(userID string, status string, friendIDs []string) {
    for _, friendID := range friendIDs {
        sseService.SendToUser(friendID, "user_status", map[string]interface{}{
            "user_id": userID,
            "status":  status,
        })
    }
}
```

### 4. 管理后台实时监控

```go
// 推送实时统计数据
func PushDashboardStats(adminGroup string, stats map[string]interface{}) {
    sseService.SendToGroup(adminGroup, "dashboard_stats", stats)
}

// 系统告警
func SendSystemAlert(level, message string) {
    sseService.SendToGroup("admins", "alert", map[string]interface{}{
        "level":     level,
        "message":   message,
        "timestamp": time.Now().Unix(),
    }, service.WithPriority(models.PriorityUrgent))
}
```

---

## 🔧 高级特性

### 1. 消息持久化和重连恢复

```go
// 发送持久化消息（用于重连恢复）
sseService.SendToUser(userID, "important", data,
    service.WithPersist(true),         // 持久化
    service.WithTTL(3600),              // 1小时过期
    service.WithPriority(models.PriorityHigh),
)

// 客户端重连时恢复消息
eventSource.addEventListener('open', () => {
    fetch(`/api/v1/sse/messages/pending?user_id=${userId}&last_event_id=${lastEventID}`)
        .then(res => res.json())
        .then(data => {
            // 处理待发送的消息
            data.data.forEach(msg => {
                handleMessage(msg);
            });
        });
});
```

### 2. 订阅机制

```go
// 用户订阅挑战更新
sseService.Subscribe("user123", "challenge_456", []string{
    "challenge_update",
    "rank_change",
})

// 取消订阅
sseService.Unsubscribe("user123", "challenge_456")

// 获取用户的所有订阅
subs, _ := sseService.GetUserSubscriptions("user123")
```

### 3. 未读消息管理

```go
// 获取未读消息数量
count, _ := sseService.GetUnreadCount("user123")

// 标记消息已读
sseService.MarkMessageRead("evt_123", "user123")
```

### 4. 消息优先级

```go
// 普通消息
sseService.SendToUser(userID, "info", data,
    service.WithPriority(models.PriorityNormal),
)

// 高优先级消息
sseService.SendToUser(userID, "warning", data,
    service.WithPriority(models.PriorityHigh),
)

// 紧急消息
sseService.SendToUser(userID, "critical", data,
    service.WithPriority(models.PriorityUrgent),
)
```

---

## 🎨 架构设计

### 核心层 (core/sse)

```
core/sse/
├── sse.go        # 核心管理器（连接管理、消息分发）
└── handler.go    # HTTP 处理器（SSE 协议实现）
```

**职责：**
- ✅ SSE 协议实现
- ✅ 连接生命周期管理
- ✅ 消息分发（单播、组播、广播）
- ✅ 心跳和超时处理
- ✅ 线程安全的并发管理

**特点：**
- 使用 `sync.Map` 实现高性能并发
- 原子操作管理计数器
- Channel 实现异步消息分发
- 自动清理超时连接

### 业务层 (app/sse)

```
app/sse/
├── models/       # 数据模型（消息、订阅）
├── repo/         # 数据库操作
├── storage/      # Redis 缓存
├── service/      # 业务逻辑
├── apis/         # HTTP 接口
└── router/       # 路由配置
```

**职责：**
- ✅ 消息持久化和重连恢复
- ✅ 订阅关系管理
- ✅ 未读消息计数
- ✅ 在线状态管理
- ✅ 业务接口封装

---

## 📊 性能优化

### 1. 并发性能

- 使用 `sync.Map` 替代 `map + sync.RWMutex`，提升并发读写性能
- 原子操作管理计数器，避免锁竞争
- Channel 异步处理消息分发

### 2. 内存优化

- 自动清理超时连接
- 定时清理过期消息
- 可配置的消息缓冲区大小

### 3. 缓存策略

```
在线状态 → Redis（5分钟）
未读计数 → Redis（24小时）
消息历史 → MySQL（可配置保留天数）
```

---

## ⚙️ 配置项

### 管理器配置

```go
SSEManager.HeartbeatInterval = 30 * time.Second // 心跳间隔
SSEManager.ClientTimeout = 90 * time.Second     // 客户端超时
SSEManager.SendTimeout = 5 * time.Second        // 发送超时
```

### 缓存配置

```go
DefaultOnlineExpire = 300      // 在线状态：5分钟
DefaultClientInfoExpire = 600  // 客户端信息：10分钟
DefaultUnreadExpire = 86400    // 未读计数：24小时
```

---

## 🐛 故障排查

### 问题 1: 连接无法建立

**检查项：**
- [ ] SSE 服务是否已初始化（`app.InitModules()`）
- [ ] 路由是否正确注册
- [ ] Nginx 是否禁用了缓冲
- [ ] 防火墙是否允许连接

### 问题 2: 收不到消息

**检查项：**
- [ ] 客户端ID和用户ID是否正确
- [ ] 事件类型是否匹配
- [ ] 用户是否在线
- [ ] 查看服务端日志

### 问题 3: 频繁断线

**检查项：**
- [ ] 网络是否稳定
- [ ] 心跳间隔和超时时间是否合理
- [ ] Nginx 是否配置正确

---

## 📚 相关文档

- [core/sse/README.md](../../core/sse/README.md) - 核心 SSE 实现文档
- [core/sse/QUICKSTART.md](../../core/sse/QUICKSTART.md) - 快速开始指南
- [core/sse/client_example.html](../../core/sse/client_example.html) - 客户端测试页面

---

## ✅ 总结

SSE 模块提供了完整的实时推送解决方案，包括：

- ✅ 高性能的核心引擎（`core/sse`）
- ✅ 完善的业务封装（`app/sse`）
- ✅ 消息持久化和重连恢复
- ✅ 丰富的API接口
- ✅ 详细的使用文档

适用于实时通知、进度推送、数据监控等多种场景！🎉
