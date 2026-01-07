# SSE API 变更说明

## 📝 变更原因

项目只支持 GET 和 POST 请求，且除了 SSE 连接端点外，其他接口都不支持 GET 请求。因此将所有查询类接口从 GET 改为 POST。

## 🔄 变更内容

### ✅ 保持不变（SSE 连接必须用 GET）

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/api/v1/sse/stream` | SSE 连接（查询参数） |
| GET | `/api/v1/sse/stream/:group/:id` | SSE 连接（路径参数） |
| GET | `/api/v1/sse/stream/:id` | SSE 连接（简化） |

### 🔄 已变更（GET → POST）

| 原方法 | 原路径 | 新方法 | 新路径 | 说明 |
|--------|--------|--------|--------|------|
| ~~GET~~ | ~~/info~~ | **POST** | `/info` | 获取管理器信息 |
| ~~GET~~ | ~~/group/:group~~ | **POST** | `/group/info` | 获取分组信息 |
| ~~POST~~ | ~~/disconnect/:id~~ | **POST** | `/disconnect` | 断开客户端 |
| ~~GET~~ | ~~/messages/pending~~ | **POST** | `/messages/pending` | 获取待发送消息 |
| ~~GET~~ | ~~/messages/unread~~ | **POST** | `/messages/unread` | 获取未读计数 |
| ~~GET~~ | ~~/subscriptions~~ | **POST** | `/subscriptions` | 获取订阅列表 |

---

## 📡 新的 API 调用方式

### 1. 获取管理器信息

**旧方式（已废弃）：**
```http
GET /api/v1/sse/info
```

**新方式：**
```http
POST /api/v1/sse/info
Content-Type: application/json

{}
```

**响应示例：**
```json
{
  "code": 200,
  "msg": "success",
  "data": {
    "clientCount": 128,
    "groupCount": 5,
    "registerChanLen": 0,
    "unregisterChanLen": 0,
    "unicastChanLen": 0,
    "groupcastChanLen": 0,
    "broadcastChanLen": 0,
    "heartbeatInterval": "30s",
    "clientTimeout": "90s"
  }
}
```

---

### 2. 获取分组信息

**旧方式（已废弃）：**
```http
GET /api/v1/sse/group/notifications
```

**新方式：**
```http
POST /api/v1/sse/group/info
Content-Type: application/json

{
  "group": "notifications"
}
```

**响应示例：**
```json
{
  "code": 200,
  "msg": "success",
  "data": {
    "group": "notifications",
    "clientCount": 15,
    "clients": ["client_001", "client_002", "..."]
  }
}
```

---

### 3. 断开客户端连接

**旧方式（已废弃）：**
```http
POST /api/v1/sse/disconnect/client_123
```

**新方式：**
```http
POST /api/v1/sse/disconnect
Content-Type: application/json

{
  "client_id": "client_123"
}
```

**响应示例：**
```json
{
  "code": 200,
  "msg": "Client disconnected successfully",
  "data": {
    "client_id": "client_123"
  }
}
```

---

### 4. 获取待发送消息（重连恢复）

**旧方式（已废弃）：**
```http
GET /api/v1/sse/messages/pending?user_id=user123&last_event_id=evt_001&limit=50
```

**新方式：**
```http
POST /api/v1/sse/messages/pending
Content-Type: application/json

{
  "user_id": "user123",
  "last_event_id": "evt_001",
  "limit": 50
}
```

**响应示例：**
```json
{
  "code": 200,
  "msg": "success",
  "data": [
    {
      "id": 1,
      "event_id": "evt_002",
      "event_type": "notification",
      "receiver_id": "user123",
      "receiver_type": "user",
      "data": "{\"title\":\"新消息\",\"content\":\"您有一条新消息\"}",
      "status": 0,
      "created_at": "2026-01-07T10:00:00Z"
    }
  ]
}
```

---

### 5. 获取未读消息数量

**旧方式（已废弃）：**
```http
GET /api/v1/sse/messages/unread?user_id=user123
```

**新方式：**
```http
POST /api/v1/sse/messages/unread
Content-Type: application/json

{
  "user_id": "user123"
}
```

**响应示例：**
```json
{
  "code": 200,
  "msg": "success",
  "data": {
    "user_id": "user123",
    "unread_count": 5
  }
}
```

---

### 6. 获取订阅列表

**旧方式（已废弃）：**
```http
GET /api/v1/sse/subscriptions?user_id=user123
```

**新方式：**
```http
POST /api/v1/sse/subscriptions
Content-Type: application/json

{
  "user_id": "user123"
}
```

**响应示例：**
```json
{
  "code": 200,
  "msg": "success",
  "data": [
    {
      "id": 1,
      "user_id": "user123",
      "group_name": "notifications",
      "event_types": ",notification,new_message,",
      "status": 1,
      "created_at": "2026-01-07T10:00:00Z"
    }
  ]
}
```

---

## 🔧 前端调用示例

### JavaScript Fetch

```javascript
// 获取未读消息数量
async function getUnreadCount(userId) {
  const response = await fetch('/api/v1/sse/messages/unread', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      user_id: userId
    })
  });
  
  const data = await response.json();
  return data.data.unread_count;
}

// 获取待发送消息
async function getPendingMessages(userId, lastEventId) {
  const response = await fetch('/api/v1/sse/messages/pending', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      user_id: userId,
      last_event_id: lastEventId,
      limit: 50
    })
  });
  
  return await response.json();
}

// 获取订阅列表
async function getSubscriptions(userId) {
  const response = await fetch('/api/v1/sse/subscriptions', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      user_id: userId
    })
  });
  
  return await response.json();
}
```

### Axios

```javascript
import axios from 'axios';

// 获取未读消息数量
const getUnreadCount = async (userId) => {
  const { data } = await axios.post('/api/v1/sse/messages/unread', {
    user_id: userId
  });
  return data.data.unread_count;
};

// 获取分组信息
const getGroupInfo = async (group) => {
  const { data } = await axios.post('/api/v1/sse/group/info', {
    group: group
  });
  return data.data;
};

// 断开客户端
const disconnectClient = async (clientId) => {
  await axios.post('/api/v1/sse/disconnect', {
    client_id: clientId
  });
};
```

---

## 📋 完整 API 列表

### SSE 连接端点（GET - 不变）

```
GET  /api/v1/sse/stream
GET  /api/v1/sse/stream/:group/:id
GET  /api/v1/sse/stream/:id
```

### 管理接口（POST）

```
POST /api/v1/sse/info           # 获取管理器信息
POST /api/v1/sse/group/info     # 获取分组信息
POST /api/v1/sse/disconnect     # 断开客户端连接
```

### 消息发送接口（POST - 不变）

```
POST /api/v1/sse/send           # 发送给指定用户
POST /api/v1/sse/send/group     # 发送到分组
POST /api/v1/sse/broadcast      # 广播
```

### 消息管理接口（POST）

```
POST /api/v1/sse/messages/pending  # 获取待发送消息
POST /api/v1/sse/messages/read     # 标记消息已读
POST /api/v1/sse/messages/unread   # 获取未读消息数量
```

### 订阅管理接口（POST - 不变）

```
POST /api/v1/sse/subscribe      # 订阅分组
POST /api/v1/sse/unsubscribe    # 取消订阅
POST /api/v1/sse/subscriptions  # 获取订阅列表
```

---

## ⚠️ 迁移注意事项

### 1. 前端需要更新的地方

- ✅ 所有查询类接口改为 POST 请求
- ✅ 参数从 URL 查询参数改为 JSON body
- ✅ 部分路径参数改为 body 参数

### 2. 测试工具

```bash
# 获取管理器信息
curl -X POST http://localhost:8000/api/v1/sse/info \
  -H "Content-Type: application/json" \
  -d '{}'

# 获取未读消息数量
curl -X POST http://localhost:8000/api/v1/sse/messages/unread \
  -H "Content-Type: application/json" \
  -d '{"user_id":"user123"}'

# 获取分组信息
curl -X POST http://localhost:8000/api/v1/sse/group/info \
  -H "Content-Type: application/json" \
  -d '{"group":"notifications"}'
```

---

## ✅ 变更完成

- ✅ 路由已更新
- ✅ API 处理函数已更新
- ✅ 参数获取方式已改为 JSON body
- ✅ 代码已格式化
- ✅ 0 个 Lint 错误
- ✅ 文档已更新

所有变更已完成，可以正常使用！🎉
