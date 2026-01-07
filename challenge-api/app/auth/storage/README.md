# Auth Storage - Redis 缓存操作层

## 📝 概述

`app/auth/storage` 模块集中管理 auth 模块的所有 Redis 缓存操作，包括：
- 登录Token管理
- 登录失败计数
- 登录锁定
- 注册限流
- 验证码管理
- 会话管理

## 🗂️ 文件结构

```
app/auth/storage/
├── auth_cache.go    # 所有Redis操作函数（250+ 行）
├── prefix.go        # 缓存key前缀（向后兼容）
└── README.md        # 本文档
```

## 🔑 缓存Key前缀

| 前缀 | 说明 | 示例 |
|------|------|------|
| `auth:login:token` | 登录token | `auth:login:token:123` |
| `auth:login:fail` | 登录失败计数 | `auth:login:fail:test001` |
| `auth:login:lock` | 登录锁定 | `auth:login:lock:test001` |
| `auth:login:session` | 登录会话 | `auth:login:session:abc123` |
| `auth:register:ip` | IP注册限流 | `auth:register:ip:192.168.1.1` |
| `auth:register:device` | 设备注册限流 | `auth:register:device:xxxxx` |
| `auth:captcha` | 图形验证码 | `auth:captcha:captcha123` |
| `auth:sms:code` | 短信验证码 | `auth:sms:code:13800138000` |
| `auth:email:code` | 邮箱验证码 | `auth:email:code:user@example.com` |

## 🚀 核心功能

### 1. 登录Token管理

用于JWT Token或Session Token的缓存。

```go
import (
    authStorage "challenge/app/auth/storage"
    "challenge/core/utils/storage"
    "context"
)

// 设置登录token（7天过期）
err := authStorage.SetLoginToken(ctx, cache, userID, token, authStorage.DefaultTokenExpire)

// 获取登录token
token, err := authStorage.GetLoginToken(ctx, cache, userID)

// 删除token（登出）
err := authStorage.DelLoginToken(ctx, cache, userID)
```

### 2. 登录失败计数（防暴力破解）

用于统计用户登录失败次数，配合锁定策略。

```go
// 获取失败次数
failCount, err := authStorage.GetLoginFailCount(ctx, cache, username)

// 增加失败次数（15分钟窗口）
err := authStorage.IncrLoginFailCount(ctx, cache, username, 900)

// 清除失败次数（登录成功后）
err := authStorage.ClearLoginFailCount(ctx, cache, username)
```

**业务规则：**
- 失败3次 → 锁定5分钟
- 失败4次 → 锁定30分钟
- 失败5次 → 永久封禁

### 3. 登录锁定

临时锁定用户登录。

```go
// 锁定5分钟
err := authStorage.LockLogin(ctx, cache, username, 300)

// 检查是否被锁定
locked, err := authStorage.IsLoginLocked(ctx, cache, username)
if locked {
    return errors.New("账号已锁定")
}

// 解除锁定（管理员操作）
err := authStorage.UnlockLogin(ctx, cache, username)
```

### 4. 注册限流

防止批量注册、刷号。

```go
// IP限流（1分钟3次）
hit, err := authStorage.CheckRegisterIPLimit(ctx, cache, ip, 3)
if hit {
    return errors.New("注册过于频繁，请稍后再试")
}

// 增加IP计数
err := authStorage.IncrRegisterIPCount(ctx, cache, ip, 60) // 1分钟窗口

// 设备限流（24小时2次）
hit, err := authStorage.CheckRegisterDeviceLimit(ctx, cache, deviceFP, 2)
if hit {
    return errors.New("该设备注册次数已达上限")
}

// 增加设备计数
err := authStorage.IncrRegisterDeviceCount(ctx, cache, deviceFP, 86400) // 24小时
```

### 5. 会话管理

基于Hash结构的会话存储，支持多字段。

```go
// 设置会话（30分钟）
userInfo := map[string]interface{}{
    "user_id": 123,
    "username": "test001",
    "role": "user",
}
err := authStorage.SetLoginSession(ctx, cache, sessionID, userInfo, 1800)

// 获取单个字段
username, err := authStorage.GetLoginSession(ctx, cache, sessionID, "username")

// 获取完整会话
session, err := authStorage.GetLoginSessionAll(ctx, cache, sessionID)

// 删除会话
err := authStorage.DelLoginSession(ctx, cache, sessionID)
```

### 6. 验证码管理

支持图形验证码、短信验证码、邮箱验证码。

```go
// 图形验证码（5分钟）
err := authStorage.SetCaptcha(ctx, cache, captchaID, "1234", 300)
code, err := authStorage.GetCaptcha(ctx, cache, captchaID)
err := authStorage.DelCaptcha(ctx, cache, captchaID) // 验证后删除

// 短信验证码（10分钟）
err := authStorage.SetSMSCode(ctx, cache, "13800138000", "123456", 600)
code, err := authStorage.GetSMSCode(ctx, cache, "13800138000")
err := authStorage.DelSMSCode(ctx, cache, "13800138000")

// 邮箱验证码（10分钟）
err := authStorage.SetEmailCode(ctx, cache, "user@example.com", "123456", 600)
code, err := authStorage.GetEmailCode(ctx, cache, "user@example.com")
err := authStorage.DelEmailCode(ctx, cache, "user@example.com")
```

## 💡 使用示例

### 完整的登录流程（含风控）

```go
func (a *AuthLogin) Login(req *dto.LoginReq) (*models.User, error) {
    ctx := context.Background()
    cache := a.Run.GetCacheAdapter()
    username := req.UserName
    
    // 1. 检查是否被锁定
    locked, _ := authStorage.IsLoginLocked(ctx, cache, username)
    if locked {
        return nil, errors.New("账号已锁定，请稍后再试")
    }
    
    // 2. 检查失败次数
    failCount, _ := authStorage.GetLoginFailCount(ctx, cache, username)
    if failCount >= 5 {
        return nil, errors.New("登录失败次数过多，账号已封禁")
    }
    
    // 3. 验证密码
    user, err := getUserByUsername(username)
    if err != nil || !checkPassword(user.Password, req.Password) {
        // 登录失败
        authStorage.IncrLoginFailCount(ctx, cache, username, 900)
        
        // 获取最新失败次数
        newFailCount, _ := authStorage.GetLoginFailCount(ctx, cache, username)
        
        // 执行锁定策略
        if newFailCount == 3 {
            authStorage.LockLogin(ctx, cache, username, 300) // 5分钟
        } else if newFailCount == 4 {
            authStorage.LockLogin(ctx, cache, username, 1800) // 30分钟
        } else if newFailCount >= 5 {
            // 永久封禁（写入数据库）
            banUser(user.ID)
        }
        
        return nil, errors.New("用户名或密码错误")
    }
    
    // 4. 登录成功
    authStorage.ClearLoginFailCount(ctx, cache, username)
    
    // 5. 生成并缓存token
    token := generateToken(user)
    authStorage.SetLoginToken(ctx, cache, user.ID, token, authStorage.DefaultTokenExpire) // 7天
    
    return user, nil
}
```

### 注册限流检查

```go
func (a *AuthRegister) Register(req *dto.RegisterReq, ip, deviceFP string) error {
    ctx := context.Background()
    cache := a.Run.GetCacheAdapter()
    
    // 1. IP限流（1分钟3次）
    hit, _ := authStorage.CheckRegisterIPLimit(ctx, cache, ip, 3)
    if hit {
        return errors.New("注册过于频繁，请1分钟后再试")
    }
    
    // 2. 设备限流（24小时2次）
    if deviceFP != "" {
        hit, _ := authStorage.CheckRegisterDeviceLimit(ctx, cache, deviceFP, 2)
        if hit {
            return errors.New("该设备今日注册次数已达上限")
        }
    }
    
    // 3. 创建用户
    user := createUser(req)
    
    // 4. 注册成功，增加计数
    authStorage.IncrRegisterIPCount(ctx, cache, ip, 60) // 1分钟窗口
    if deviceFP != "" {
        authStorage.IncrRegisterDeviceCount(ctx, cache, deviceFP, 86400) // 24小时窗口
    }
    
    return nil
}
```

## 🔧 技术要点

### 1. 缓存穿透保护

所有函数都包含错误处理，缓存不可用时自动降级。

```go
val, err := cache.Get(prefix, key)
if err != nil || val == "" {
    // 缓存miss，返回默认值
    return 0, nil
}
```

### 2. 原子操作

使用 Redis 的 `INCR` 保证计数的原子性。

```go
// 第一次计数
cache.Set(prefix, key, "1", expire)

// 后续递增（原子操作）
cache.Increase(prefix, key)
```

### 3. 过期时间

| 场景 | 过期时间 | 说明 |
|------|---------|------|
| 登录Token | 7天 | 保持登录状态 |
| 登录失败计数 | 15分钟 | 滑动窗口 |
| 临时锁定 | 5分钟/30分钟 | 分级锁定 |
| IP注册限流 | 1分钟 | 防刷 |
| 设备注册限流 | 24小时 | 防批量注册 |
| 验证码 | 5-10分钟 | 安全性 |
| 会话 | 30分钟 | 活跃保持 |

### 4. 多维度联动

登录失败计数支持三个维度（在 `risk/storage` 中）：
- **User维度**：`auth:login:fail:username`
- **IP维度**：`risk:login:fail:ip:192.168.1.1`
- **Device维度**：`risk:login:fail:device:xxxxx`

## ⚠️ 注意事项

1. **Context传递**：所有函数都需要 `context.Context`，用于超时控制和链路追踪。

2. **缓存降级**：
   ```go
   cache := a.Run.GetCacheAdapter()
   if cache == nil {
       // 缓存不可用，直接通过或查数据库
       return false, nil
   }
   ```

3. **Key命名规范**：
   - 使用冒号 `:` 分隔层级
   - 格式：`模块:功能:维度:标识`
   - 示例：`auth:login:fail:test001`

4. **过期时间单位**：所有过期时间都是**秒**。

5. **空设备指纹处理**：
   ```go
   if deviceFP == "" {
       return false, nil // 设备指纹为空时跳过检查
   }
   ```

## 🎯 与 risk/storage 的分工

| 模块 | 职责 | 示例 |
|------|------|------|
| **auth/storage** | 认证相关的Redis操作 | Token、登录锁定、注册限流、验证码 |
| **risk/storage** | 风控相关的Redis操作 | 策略缓存、黑名单、多维度计数 |

**协作示例：**
```go
// auth/storage - 用户维度
authStorage.IncrLoginFailCount(ctx, cache, username, 900)

// risk/storage - IP/设备维度
riskStorage.IncrLoginFailCount(ctx, cache, "ip", ip, 900)
riskStorage.IncrLoginFailCount(ctx, cache, "device", deviceFP, 900)
```

## 📊 Redis数据示例

```bash
# 登录失败计数
redis> GET "auth:login:fail:test001"
"3"

# 登录锁定
redis> GET "auth:login:lock:test001"
"1"
redis> TTL "auth:login:lock:test001"
(integer) 280

# 注册IP限流
redis> GET "auth:register:ip:192.168.1.1"
"2"

# 会话Hash
redis> HGETALL "auth:login:session:abc123"
1) "user_id"
2) "123"
3) "username"
4) "test001"
5) "role"
6) "user"
```

## 🚀 性能优化建议

1. **批量操作**：对于多个维度的计数，考虑使用 Pipeline。
2. **监控告警**：监控Redis慢查询和缓存命中率。
3. **热key处理**：对于高频访问的key（如IP限流），可考虑本地缓存。
4. **容量规划**：注册限流key在高峰期可能非常多，注意内存规划。

---

**文档维护：** 2026-01-07  
**相关模块：** `app/auth/service`, `app/risk/storage`
