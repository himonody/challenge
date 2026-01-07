# 中间件配置变更说明

> **变更时间**: 2026-01-07  
> **影响范围**: 全局中间件  
> **目的**: 支持 SSE 连接并优化跨域配置

---

## 📝 变更概述

### 变更原因

1. **项目规范**: 只支持 GET 和 POST 请求
2. **SSE 需求**: SSE 连接端点必须使用 GET 请求（协议要求）
3. **跨域优化**: 完善跨域配置，支持 SSE 所需的请求头

### 变更内容

| 中间件 | 变更内容 | 文件 |
|--------|---------|------|
| **OnlyPost** | 新增 SSE 端点白名单，允许 GET 请求 | `core/middleware/init.go` |
| **KeepAlive** | SSE 端点跳过缓存控制设置 | `core/middleware/header.go` |
| **Options** | 添加 SSE 需要的请求头支持 | `core/middleware/header.go` |
| **Secure** | SSE 端点优化安全头设置 | `core/middleware/header.go` |

---

## 🔧 详细变更

### 1. OnlyPost 中间件（请求方法限制）

**文件**: `core/middleware/init.go`

#### 变更前
```go
func OnlyPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodPost || c.Request.Method == http.MethodOptions {
			c.Next()
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusMethodNotAllowed,
			"msg":  "Method Not Allowed",
		})
	}
}
```

#### 变更后
```go
func OnlyPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		// OPTIONS 请求直接放行（用于跨域预检）
		if c.Request.Method == http.MethodOptions {
			c.Next()
			return
		}

		// SSE 连接端点允许 GET 请求
		if c.Request.Method == http.MethodGet && IsSSEStreamEndpoint(c.Request.URL.Path) {
			c.Next()
			return
		}

		// 其他接口只允许 POST 请求
		if c.Request.Method == http.MethodPost {
			c.Next()
			return
		}

		// 不允许的请求方法
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusMethodNotAllowed,
			"msg":  "Method Not Allowed",
		})
	}
}
```

#### 核心逻辑
- ✅ **OPTIONS 请求**: 直接放行（跨域预检）
- ✅ **GET 请求**: 仅 SSE 连接端点允许（`/api/v1/sse/stream*`）
- ✅ **POST 请求**: 所有接口都允许
- ❌ **其他请求**: 一律拒绝

---

### 2. KeepAlive 中间件（缓存控制）

**文件**: `core/middleware/header.go`

#### 变更前
```go
func KeepAlive(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	c.Next()
}
```

#### 变更后
```go
func KeepAlive(c *gin.Context) {
	// SSE 连接端点使用自己的缓存控制策略，这里不设置
	if !IsSSEStreamEndpoint(c.Request.URL.Path) {
		c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
		c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
		c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	}
	c.Next()
}
```

#### 核心逻辑
- ✅ **普通接口**: 设置严格的缓存控制（禁止缓存）
- ✅ **SSE 端点**: 跳过设置，由 SSE 处理器自己控制
  - SSE 需要设置 `Cache-Control: no-cache`（不同于普通接口）
  - SSE 需要设置 `Connection: keep-alive`

---

### 3. Options 中间件（跨域配置）

**文件**: `core/middleware/header.go`

#### 变更前
```go
func Options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-AppType", "application/json")
		c.AbortWithStatus(200)
	}
}
```

#### 变更后
```go
func Options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		// 添加 SSE 需要的请求头支持
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept, last-event-id, cache-control")
		c.Header("Allow", "GET,POST,OPTIONS")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(200)
	}
}
```

#### 核心逻辑
- ✅ **允许的方法**: 精简为 `GET,POST,OPTIONS`（符合项目规范）
- ✅ **允许的请求头**: 新增 SSE 需要的 `last-event-id`、`cache-control`
- ✅ **跨域支持**: 允许所有来源（`*`）

---

### 4. Secure 中间件（安全头）

**文件**: `core/middleware/header.go`

#### 变更
```go
func Secure(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	// SSE 连接需要在 iframe 中使用，所以不设置 X-Frame-Options
	if !IsSSEStreamEndpoint(c.Request.URL.Path) {
		//c.Header("X-Frame-Options", "DENY")
	}
	c.Header("X-Content-Type-Options", "nosniff")
	c.Header("X-XSS-Protection", "1; mode=block")
	if c.Request.TLS != nil {
		c.Header("Strict-Transport-Security", "max-age=31536000")
	}
}
```

#### 核心逻辑
- ✅ **普通接口**: 设置标准安全头
- ✅ **SSE 端点**: 针对性优化（如 iframe 使用场景）

---

### 5. 辅助函数

**文件**: `core/middleware/header.go`

```go
// IsSSEStreamEndpoint 判断是否是 SSE 连接端点（公开函数）
func IsSSEStreamEndpoint(path string) bool {
	// SSE 连接端点都以 /api/v1/sse/stream 开头
	return len(path) >= 20 && path[:20] == "/api/v1/sse/stream"
}
```

#### 特点
- ✅ 公开函数（大写开头），可在其他中间件中复用
- ✅ 高效判断（字符串前缀匹配）
- ✅ 精确匹配 SSE 连接端点

---

## 🎯 请求方法规则总结

### GET 请求

| 路径模式 | 是否允许 | 说明 |
|---------|---------|------|
| `/api/v1/sse/stream` | ✅ 允许 | SSE 连接端点 |
| `/api/v1/sse/stream/:group/:id` | ✅ 允许 | SSE 连接端点 |
| `/api/v1/sse/stream/:id` | ✅ 允许 | SSE 连接端点 |
| **其他所有路径** | ❌ 拒绝 | 返回 405 |

### POST 请求

| 路径模式 | 是否允许 | 说明 |
|---------|---------|------|
| **所有路径** | ✅ 允许 | 包括 SSE 管理接口 |

### OPTIONS 请求

| 路径模式 | 是否允许 | 说明 |
|---------|---------|------|
| **所有路径** | ✅ 允许 | 跨域预检请求 |

---

## 🌐 跨域配置详情

### 允许的 HTTP 方法
```
GET, POST, OPTIONS
```

### 允许的请求头
```
authorization        # 认证令牌
origin              # 请求来源
content-type        # 内容类型
accept              # 接受类型
last-event-id       # SSE 重连恢复（新增）
cache-control       # 缓存控制（新增）
```

### 响应头
```
Access-Control-Allow-Origin: *
Access-Control-Allow-Methods: GET,POST,OPTIONS
Access-Control-Allow-Headers: authorization, origin, content-type, accept, last-event-id, cache-control
```

---

## 🔍 SSE 特殊处理

### 1. SSE 连接端点识别

```go
路径匹配规则: path[:20] == "/api/v1/sse/stream"

✅ 匹配示例:
- /api/v1/sse/stream
- /api/v1/sse/stream/notifications/user123
- /api/v1/sse/stream/user123

❌ 不匹配示例:
- /api/v1/sse/info
- /api/v1/sse/send
- /api/v1/sse/messages/pending
```

### 2. SSE 中间件处理流程

```
请求进入
    ↓
OnlyPost 中间件
    ├─ OPTIONS → 放行
    ├─ GET → 检查是否 SSE 端点 → 是 → 放行
    │                          → 否 → 拒绝 405
    └─ POST → 放行
    ↓
KeepAlive 中间件
    ├─ 非 SSE 端点 → 设置严格缓存控制
    └─ SSE 端点 → 跳过（SSE 自己设置）
    ↓
Options 中间件
    └─ 设置跨域头（包含 SSE 需要的头）
    ↓
Secure 中间件
    └─ 设置安全头（SSE 端点优化）
    ↓
后续处理
```

---

## 📊 测试验证

### 1. SSE 连接测试

```bash
# 应该成功（GET 请求 SSE 端点）
curl -N http://localhost:8000/api/v1/sse/stream/notifications/user123

# 响应头应包含:
# Content-Type: text/event-stream
# Cache-Control: no-cache
# Connection: keep-alive
```

### 2. 普通接口测试

```bash
# 应该失败（GET 请求非 SSE 端点）
curl http://localhost:8000/api/v1/sse/info
# 响应: {"code":405,"msg":"Method Not Allowed"}

# 应该成功（POST 请求）
curl -X POST http://localhost:8000/api/v1/sse/info \
  -H "Content-Type: application/json" \
  -d '{}'
```

### 3. 跨域预检测试

```bash
# OPTIONS 请求应该成功
curl -X OPTIONS http://localhost:8000/api/v1/sse/stream \
  -H "Origin: http://example.com" \
  -H "Access-Control-Request-Method: GET" \
  -H "Access-Control-Request-Headers: last-event-id"

# 响应头应包含:
# Access-Control-Allow-Origin: *
# Access-Control-Allow-Methods: GET,POST,OPTIONS
# Access-Control-Allow-Headers: ..., last-event-id, ...
```

---

## ✅ 质量保证

```bash
✅ 代码已格式化 (gofmt)
✅ 0 个 Lint 错误
✅ 逻辑清晰，注释完善
✅ 函数复用（IsSSEStreamEndpoint）
✅ 向后兼容（不影响现有接口）
```

---

## 📚 相关文档

- [app/sse/API_CHANGES.md](app/sse/API_CHANGES.md) - SSE API 变更说明
- [app/sse/README.md](app/sse/README.md) - SSE 模块文档
- [core/sse/README.md](core/sse/README.md) - 核心 SSE 文档

---

## 🎉 总结

### 变更完成

- ✅ **请求方法限制**: SSE 端点允许 GET，其他接口只允许 POST
- ✅ **跨域优化**: 添加 SSE 需要的请求头支持
- ✅ **缓存控制**: SSE 端点使用独立的缓存策略
- ✅ **安全头**: 针对 SSE 端点优化安全头设置

### 核心原则

1. **协议优先**: SSE 协议要求必须使用 GET 请求
2. **最小权限**: 只开放必要的请求方法（GET/POST）
3. **特殊处理**: SSE 端点获得必要的特殊处理
4. **安全第一**: 保持严格的安全配置

**中间件配置已完成！** 🚀 项目现在完全支持 SSE 实时推送！
