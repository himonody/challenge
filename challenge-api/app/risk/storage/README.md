# Risk Storage - Redis 缓存操作层

## 📝 概述

Risk Storage 模块负责风控系统的所有 Redis 缓存操作，包括限流、锁定、策略缓存、黑名单缓存等。

## 🗂️ 文件结构

```
app/risk/storage/
├── rate_limit.go         # 限流和锁定操作（243行）
├── strategy_cache.go     # 策略和黑名单缓存（119行）
└── README.md            # 本文档
```

## 🔑 缓存Key设计

### Key 命名规范

遵循统一的命名规范：`模块:功能:维度:标识`

### 过期时间常量

```go
// rate_limit.go
DefaultRegisterIPWindow     = 60     // IP注册限流：1分钟
DefaultRegisterDeviceWindow = 86400  // 设备注册限流：24小时
DefaultLoginFailWindow      = 900    // 登录失败窗口：15分钟
DefaultLockDuration5M       = 300    // 锁定：5分钟
DefaultLockDuration30M      = 1800   // 锁定：30分钟

// strategy_cache.go
DefaultStrategyTTL  = 300   // 策略缓存：5分钟
DefaultBlacklistTTL = 600   // 黑名单缓存：10分钟
DefaultRiskScoreTTL = 3600  // 风险分数缓存：1小时
```

### Key 前缀定义

| 前缀 | 说明 | 过期时间 | 示例 |
|------|------|---------|------|
| `risk:register:ip` | IP注册限流 | 1分钟 | `risk:register:ip:192.168.1.1` |
| `risk:register:device` | 设备注册限流 | 24小时 | `risk:register:device:fp123` |
| `risk:login:fail:user` | 用户登录失败计数 | 15分钟 | `risk:login:fail:user:test001` |
| `risk:login:fail:ip` | IP登录失败计数 | 15分钟 | `risk:login:fail:ip:192.168.1.1` |
| `risk:login:fail:device` | 设备登录失败计数 | 15分钟 | `risk:login:fail:device:fp123` |
| `risk:login:lock:user` | 用户登录锁定 | 5分钟/30分钟 | `risk:login:lock:user:test001` |
| `risk:login:lock:ip` | IP锁定 | 动态 | `risk:login:lock:ip:192.168.1.1` |
| `risk:login:lock:device` | 设备锁定 | 动态 | `risk:login:lock:device:fp123` |
| `risk:strategy:{scene}` | 策略缓存 | 5分钟 | `risk:strategy:register` |
| `risk:blacklist:{type}:{value}` | 黑名单缓存 | 10分钟 | `risk:blacklist:ip:192.168.1.1` |
| `risk:score:{userID}` | 风险分数缓存 | 1小时 | `risk:score:123` |

## 📋 rate_limit.go - 限流和锁定

### 注册限流

```go
// 检查IP注册限流（1分钟3次）
hit, err := CheckRegisterIPLimit(ctx, cache, ip, window, limit)
if hit {
    return errors.New("注册过于频繁")
}

// 增加IP注册计数
err := IncrRegisterIPLimit(ctx, cache, ip, DefaultRegisterIPWindow)

// 检查设备注册限流（24小时2次）
hit, err := CheckRegisterDeviceLimit(ctx, cache, deviceFP, window, limit)

// 增加设备注册计数
err := IncrRegisterDeviceLimit(ctx, cache, deviceFP, DefaultRegisterDeviceWindow)
```

### 登录失败计数（三维度）

```go
// 获取失败次数
failCount, err := GetLoginFailCount(ctx, cache, "user", username)
failCount, err := GetLoginFailCount(ctx, cache, "ip", ip)
failCount, err := GetLoginFailCount(ctx, cache, "device", deviceFP)

// 增加失败计数
err := IncrLoginFailCount(ctx, cache, "user", username, DefaultLoginFailWindow)
err := IncrLoginFailCount(ctx, cache, "ip", ip, DefaultLoginFailWindow)
err := IncrLoginFailCount(ctx, cache, "device", deviceFP, DefaultLoginFailWindow)

// 清除失败计数
err := ClearLoginFailCount(ctx, cache, "user", username)

// 清除所有维度的失败计数
err := ClearAllLoginFails(ctx, cache, username, ip, deviceFP)
```

### 锁定管理

```go
// 用户锁定
err := LockLoginUser(ctx, cache, username, DefaultLockDuration5M)
locked, err := IsLoginUserLocked(ctx, cache, username)
ttl, err := GetLockTTL(ctx, cache, username)

// IP锁定
err := LockIP(ctx, cache, ip, DefaultLockDuration30M)
locked, err := IsIPLocked(ctx, cache, ip)
err := UnlockIP(ctx, cache, ip)

// 设备锁定
err := LockDevice(ctx, cache, deviceFP, DefaultLockDuration30M)
locked, err := IsDeviceLocked(ctx, cache, deviceFP)
err := UnlockDevice(ctx, cache, deviceFP)
```

## 📋 strategy_cache.go - 策略和缓存

### 策略缓存

```go
// 缓存策略集合
items := []models.RiskStrategyCache{ ... }
err := CacheStrategies(ctx, cache, "register", items)

// 获取策略集合
strategies, err := GetStrategies(ctx, cache, "register")

// 清除策略缓存
err := ClearStrategyCache(ctx, cache, "register")
```

### 黑名单缓存

```go
// 缓存黑名单命中结果
err := CacheBlacklistFlag(ctx, cache, "ip", "192.168.1.1", true, 10*time.Minute)

// 获取黑名单命中结果
hit, ok := GetBlacklistFlag(ctx, cache, "ip", "192.168.1.1")
if ok && hit {
    return errors.New("IP已被封禁")
}

// 清除黑名单缓存
err := ClearBlacklistCache(ctx, cache, "ip", "192.168.1.1")
```

### 风险分数缓存

```go
// 缓存风险分数
err := CacheRiskScore(ctx, cache, userID, 50, time.Hour)

// 获取风险分数
score, ok := GetRiskScore(ctx, cache, userID)
if ok {
    fmt.Printf("用户风险分数: %d\n", score)
}
```

## 💡 使用示例

### 完整的注册限流流程

```go
import riskStorage "challenge/app/risk/storage"

func CheckRegisterRisk(ctx context.Context, cache storage.AdapterCache, ip, deviceFP string) error {
    // 1. IP限流检查（1分钟3次）
    hit, _ := riskStorage.CheckRegisterIPLimit(
        ctx, cache, ip, 
        riskStorage.DefaultRegisterIPWindow, 3,
    )
    if hit {
        return errors.New("IP注册过于频繁")
    }
    
    // 2. 设备限流检查（24小时2次）
    if deviceFP != "" {
        hit, _ := riskStorage.CheckRegisterDeviceLimit(
            ctx, cache, deviceFP,
            riskStorage.DefaultRegisterDeviceWindow, 2,
        )
        if hit {
            return errors.New("设备注册次数已达上限")
        }
    }
    
    return nil
}

func RecordRegisterSuccess(ctx context.Context, cache storage.AdapterCache, ip, deviceFP string) {
    // 注册成功，更新计数
    _ = riskStorage.IncrRegisterIPLimit(ctx, cache, ip, riskStorage.DefaultRegisterIPWindow)
    if deviceFP != "" {
        _ = riskStorage.IncrRegisterDeviceLimit(ctx, cache, deviceFP, riskStorage.DefaultRegisterDeviceWindow)
    }
}
```

### 完整的登录失败处理流程

```go
func HandleLoginFailure(ctx context.Context, cache storage.AdapterCache, username, ip, deviceFP string) error {
    // 1. 增加三维度失败计数
    _ = riskStorage.IncrLoginFailCount(ctx, cache, "user", username, riskStorage.DefaultLoginFailWindow)
    _ = riskStorage.IncrLoginFailCount(ctx, cache, "ip", ip, riskStorage.DefaultLoginFailWindow)
    if deviceFP != "" {
        _ = riskStorage.IncrLoginFailCount(ctx, cache, "device", deviceFP, riskStorage.DefaultLoginFailWindow)
    }
    
    // 2. 获取失败次数
    failCount, _ := riskStorage.GetLoginFailCount(ctx, cache, "user", username)
    
    // 3. 执行锁定策略
    switch failCount {
    case 3:
        // 3次失败，锁定5分钟
        _ = riskStorage.LockLoginUser(ctx, cache, username, riskStorage.DefaultLockDuration5M)
        return errors.New("登录失败3次，账号已锁定5分钟")
        
    case 4:
        // 4次失败，锁定30分钟
        _ = riskStorage.LockLoginUser(ctx, cache, username, riskStorage.DefaultLockDuration30M)
        return errors.New("登录失败4次，账号已锁定30分钟")
        
    case 5:
        // 5次失败，永久封禁（需要调用 service 层添加黑名单）
        return errors.New("登录失败5次，账号已封禁")
    }
    
    return nil
}
```

### 完整的登录成功处理流程

```go
func HandleLoginSuccess(ctx context.Context, cache storage.AdapterCache, username, ip, deviceFP string) {
    // 清除所有维度的失败计数
    _ = riskStorage.ClearAllLoginFails(ctx, cache, username, ip, deviceFP)
}
```

## 🔧 辅助工具

### BuildRiskKey

```go
// 构建风控相关的缓存key
key := BuildRiskKey("risk:custom", "identifier")
// 返回: "risk:custom:identifier"
```

### GetDimensionPrefix

```go
// 根据维度获取前缀
prefix, err := GetDimensionPrefix("user")  // "risk:login:fail:user"
prefix, err := GetDimensionPrefix("ip")    // "risk:login:fail:ip"
prefix, err := GetDimensionPrefix("device") // "risk:login:fail:device"
```

## 📊 Redis 数据示例

### 查看注册限流

```bash
redis-cli

# IP限流
> GET "risk:register:ip:192.168.1.1"
"2"  # 已注册2次

> TTL "risk:register:ip:192.168.1.1"
(integer) 45  # 还有45秒过期

# 设备限流
> GET "risk:register:device:fp123"
"1"

> TTL "risk:register:device:fp123"
(integer) 86395  # 24小时
```

### 查看登录失败计数

```bash
# 用户维度
> GET "risk:login:fail:user:test001"
"3"

# IP维度
> GET "risk:login:fail:ip:192.168.1.1"
"5"

# 设备维度
> GET "risk:login:fail:device:fp123"
"2"
```

### 查看锁定状态

```bash
# 用户锁定
> GET "risk:login:lock:user:test001"
"1"

> TTL "risk:login:lock:user:test001"
(integer) 280  # 还有280秒（约4分40秒）
```

### 查看策略缓存

```bash
# 策略缓存（JSON格式）
> GET "risk:strategy:register"
"[{\"scene\":\"register\",\"rule_code\":\"REG_IP_LIMIT\",\"threshold\":3,...}]"
```

### 查看黑名单缓存

```bash
# 黑名单缓存
> GET "risk:blacklist:ip:192.168.1.100"
"1"  # 1表示已被封禁，0表示正常
```

### 查看风险分数

```bash
# 风险分数
> GET "risk:score:123"
"50"  # 用户123的风险分数为50
```

## ⚠️ 注意事项

1. **Context 传递**
   - 所有函数都需要 `context.Context`
   - 用于超时控制和链路追踪

2. **缓存降级**
   - 当缓存不可用时，函数应优雅降级
   - 不阻塞主流程

3. **Key 命名**
   - 使用常量定义，避免硬编码
   - 遵循统一的命名规范

4. **过期时间**
   - 使用预定义的常量
   - 单位统一为秒

5. **维度参数**
   - dimension 支持: "user", "ip", "device"
   - 传入无效维度会返回错误

## 🔄 与 auth/storage 的对比

| 模块 | 职责 | Key 前缀 |
|------|------|----------|
| **auth/storage** | 认证相关缓存 | `auth:*` |
| **risk/storage** | 风控相关缓存 | `risk:*` |

**协作示例：**

```go
// auth/storage - 用户维度（单维度）
authStorage.GetLoginFailCount(ctx, cache, username)
authStorage.LockLogin(ctx, cache, username, 300)

// risk/storage - 多维度（三维度）
riskStorage.GetLoginFailCount(ctx, cache, "user", username)
riskStorage.GetLoginFailCount(ctx, cache, "ip", ip)
riskStorage.GetLoginFailCount(ctx, cache, "device", deviceFP)
```

## 🚀 性能优化建议

1. **批量操作**
   - 使用 `ClearAllLoginFails` 批量清除
   - 考虑使用 Redis Pipeline

2. **合理的过期时间**
   - 避免过长导致内存占用
   - 避免过短导致频繁查询

3. **监控缓存命中率**
   - 策略缓存命中率应 > 90%
   - 黑名单缓存命中率应 > 85%

4. **Key 数量控制**
   - 定期清理过期Key
   - 监控Redis内存使用

---

**文档版本：** v1.0  
**创建时间：** 2026-01-07  
**维护者：** Challenge Team
