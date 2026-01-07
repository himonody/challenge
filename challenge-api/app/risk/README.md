# Risk 模块 - 风险控制系统

## 📦 模块概览

完整的风险控制系统，实现四层风控模型，提供黑名单、限流、评分、锁定等全方位风控能力。

## 📁 目录结构

```
app/risk/
├── models/                    # 数据模型层
│   ├── risk_user.go          # 风控用户表
│   ├── risk_device.go        # 设备管理表
│   ├── risk_event.go         # 风控事件表
│   ├── risk_strategy.go      # 策略配置表
│   ├── risk_strategy_cache.go # 策略缓存表
│   ├── risk_action.go        # 动作定义表
│   ├── risk_blacklist.go     # 黑名单表
│   ├── risk_rate_limit.go    # 限流记录表
│   └── risk_appeal.go        # 申诉表
├── repo/                      # 数据访问层
│   ├── risk_user.go          # 用户风控操作
│   ├── risk_device.go        # 设备操作
│   ├── risk_event.go         # 事件操作
│   ├── risk_strategy.go      # 策略操作
│   ├── risk_action.go        # 动作操作
│   ├── risk_blacklist.go     # 黑名单操作
│   └── risk_rate_limit.go    # 限流操作
├── service/                   # 服务层
│   ├── risk.go               # 风控核心服务
│   └── dto/
│       └── risk_context.go   # 风控上下文
├── storage/                   # Redis缓存层
│   ├── rate_limit.go         # 限流缓存操作（240+ 行）✨
│   └── strategy_cache.go     # 策略缓存操作（100+ 行）✨
└── README.md                  # 本文档
```

## 🎯 核心功能

### 1. 四层风控模型

```
┌─────────────────────────────────────────┐
│  第1层：信号采集层（RiskContext）       │
│  - IP地址、设备指纹、UA、用户ID         │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│  第2层：实时拦截层（Redis）             │
│  - 注册限流（IP/设备）                  │
│  - 登录锁定（User/IP/Device）           │
│  - 黑名单检查（三维）                   │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│  第3层：风险评估层（评分模型）          │
│  - 行为评分、累计分数                   │
│  - 分数映射风险等级（0-3）              │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│  第4层：人工兜底层（申诉）              │
│  - 申诉机制、人工审核                   │
└─────────────────────────────────────────┘
```

### 2. 风险等级定义

| 等级 | 值 | 分数范围 | 说明 | 行为 |
|------|---|---------|------|------|
| **正常** | 0 | 0-19 | 正常用户 | 无限制 |
| **观察** | 1 | 20-49 | 轻微异常 | 增加验证频率 |
| **限制** | 2 | 50-79 | 中度风险 | 限制部分功能 |
| **封禁** | 3 | 80+ | 高度风险 | 禁止所有操作 |

### 3. 三维风控联动

| 维度 | 作用 | 限流策略 | 锁定策略 |
|------|------|---------|---------|
| **User** | 账号级风控 | 登录失败计数 | 3次→5分钟，5次→永久 |
| **IP** | 识别撞库/扫描 | 注册1分钟3次 | 异常行为锁定 |
| **Device** | 识别工作室 | 注册24小时2次 | 设备封禁 |

## 💾 数据模型

### app_risk_user（用户风控表）

| 字段 | 类型 | 说明 |
|------|------|------|
| user_id | bigint | 用户ID（主键） |
| risk_level | tinyint | 风险等级：0正常 1观察 2限制 3封禁 |
| risk_score | int | 风险评分（0-100+） |
| reason | varchar(255) | 风险原因 |
| updated_at | datetime | 更新时间 |

### app_risk_event（风控事件表）

| 字段 | 类型 | 说明 |
|------|------|------|
| id | bigint | 事件ID |
| user_id | bigint | 用户ID |
| event_type | int | 事件类型 |
| detail | text | 事件详情 |
| score | int | 分数变化 |
| ip | varchar(45) | IP地址 |
| device_fp | varchar(64) | 设备指纹 |
| created_at | datetime | 创建时间 |

### app_risk_blacklist（黑名单表）

| 字段 | 类型 | 说明 |
|------|------|------|
| id | bigint | 主键 |
| type | varchar(16) | 类型：ip/device/mobile/email |
| value | varchar(128) | 值 |
| risk_level | tinyint | 风险等级 |
| reason | varchar(255) | 原因 |
| status | char(1) | 状态：1生效 2失效 |
| created_at | datetime | 创建时间 |

## 🔧 Storage 层（Redis操作）

### rate_limit.go - 限流和锁定

```go
// 注册限流
CheckRegisterIPLimit(ctx, cache, ip, window, limit) (bool, error)
IncrRegisterIPLimit(ctx, cache, ip, window) error

CheckRegisterDeviceLimit(ctx, cache, deviceFP, window, limit) (bool, error)
IncrRegisterDeviceLimit(ctx, cache, deviceFP, window) error

// 登录失败计数（三维度）
GetLoginFailCount(ctx, cache, dimension, key) (int, error)  // dimension: user/ip/device
IncrLoginFailCount(ctx, cache, dimension, key, window) error
ClearLoginFailCount(ctx, cache, dimension, key) error
ClearAllLoginFails(ctx, cache, username, ip, deviceFP) error

// 用户锁定
LockLoginUser(ctx, cache, username, seconds) error
IsLoginUserLocked(ctx, cache, username) (bool, error)
GetLockTTL(ctx, cache, username) (int, error)

// IP/设备锁定
LockIP(ctx, cache, ip, seconds) error
IsIPLocked(ctx, cache, ip) (bool, error)
UnlockIP(ctx, cache, ip) error

LockDevice(ctx, cache, deviceFP, seconds) error
IsDeviceLocked(ctx, cache, deviceFP) (bool, error)
UnlockDevice(ctx, cache, deviceFP) error
```

### strategy_cache.go - 策略和黑名单缓存

```go
// 策略缓存
CacheStrategies(ctx, cache, scene, items) error
GetStrategies(ctx, cache, scene) ([]models.RiskStrategyCache, error)
ClearStrategyCache(ctx, cache, scene) error

// 黑名单缓存
CacheBlacklistFlag(ctx, cache, typ, value, blocked, ttl) error
GetBlacklistFlag(ctx, cache, typ, value) (bool, bool)
ClearBlacklistCache(ctx, cache, typ, value) error

// 风险分数缓存
CacheRiskScore(ctx, cache, userID, score, ttl) error
GetRiskScore(ctx, cache, userID) (int64, bool)
```

## 🚀 Service 层

### risk.go - 核心服务

```go
type Risk struct {
    service.Service
}

// 策略管理
LoadStrategies(ctx, scene) ([]models.RiskStrategyCache, int, error)
ListActions() (map[string]models.RiskAction, int, error)
RefreshStrategyCache(ctx, scene) error

// 黑名单管理
CheckBlacklist(ctx, typ, value) (bool, int, error)
AddToBlacklist(ctx, typ, value, reason) error
RemoveFromBlacklist(ctx, typ, value) error

// 用户风控
GetUserRiskLevel(ctx, userID) (int64, int64, error)  // 返回: level, score, error
UpdateUserRiskScore(ctx, userID, deltaScore, reason) error
```

## 📊 Redis Key 设计

| Key前缀 | 说明 | 过期时间 | 示例 |
|---------|------|---------|------|
| `risk:register:ip` | IP注册限流 | 1分钟 | `risk:register:ip:192.168.1.1` |
| `risk:register:device` | 设备注册限流 | 24小时 | `risk:register:device:xxx` |
| `risk:login:fail:user` | 用户登录失败计数 | 15分钟 | `risk:login:fail:user:test001` |
| `risk:login:fail:ip` | IP登录失败计数 | 15分钟 | `risk:login:fail:ip:192.168.1.1` |
| `risk:login:fail:device` | 设备登录失败计数 | 15分钟 | `risk:login:fail:device:xxx` |
| `risk:login:lock:user` | 用户锁定 | 动态 | `risk:login:lock:user:test001` |
| `risk:login:lock:ip` | IP锁定 | 动态 | `risk:login:lock:ip:192.168.1.1` |
| `risk:login:lock:device` | 设备锁定 | 动态 | `risk:login:lock:device:xxx` |
| `risk:strategy` | 策略缓存 | 5分钟 | `risk:strategy:register` |
| `risk:blacklist` | 黑名单缓存 | 10分钟 | `risk:blacklist:ip:xxx` |
| `risk:score` | 风险分数缓存 | 1小时 | `risk:score:123` |

## 💡 使用示例

### 1. 检查黑名单

```go
riskSvc := service.NewRiskService(&baseService)

// 检查IP黑名单
hit, code, err := riskSvc.CheckBlacklist(ctx, "ip", "192.168.1.1")
if hit {
    return errors.New("IP已被封禁")
}

// 检查设备黑名单
hit, code, err := riskSvc.CheckBlacklist(ctx, "device", deviceFP)
if hit {
    return errors.New("设备已被封禁")
}
```

### 2. 注册限流检查

```go
import riskStorage "challenge/app/risk/storage"

cache := runtime.GetCacheAdapter()

// IP限流：1分钟3次
hit, _ := riskStorage.CheckRegisterIPLimit(ctx, cache, ip, 60, 3)
if hit {
    return errors.New("注册过于频繁")
}

// 设备限流：24小时2次
hit, _ := riskStorage.CheckRegisterDeviceLimit(ctx, cache, deviceFP, 86400, 2)
if hit {
    return errors.New("该设备注册次数已达上限")
}

// 注册成功，更新计数
_ = riskStorage.IncrRegisterIPLimit(ctx, cache, ip, 60)
_ = riskStorage.IncrRegisterDeviceLimit(ctx, cache, deviceFP, 86400)
```

### 3. 登录失败处理

```go
// 增加失败计数（三维度）
_ = riskStorage.IncrLoginFailCount(ctx, cache, "user", username, 900)
_ = riskStorage.IncrLoginFailCount(ctx, cache, "ip", ip, 900)
_ = riskStorage.IncrLoginFailCount(ctx, cache, "device", deviceFP, 900)

// 获取失败次数
failCount, _ := riskStorage.GetLoginFailCount(ctx, cache, "user", username)

// 执行锁定策略
switch failCount {
case 3:
    _ = riskStorage.LockLoginUser(ctx, cache, username, 300) // 5分钟
case 4:
    _ = riskStorage.LockLoginUser(ctx, cache, username, 1800) // 30分钟
case 5:
    // 永久封禁
    _ = riskSvc.AddToBlacklist(ctx, "user", username, "登录失败5次")
}
```

### 4. 登录成功清理

```go
// 清除所有维度的失败计数
_ = riskStorage.ClearAllLoginFails(ctx, cache, username, ip, deviceFP)
```

### 5. 用户风险评分

```go
// 获取用户风险等级和分数
level, score, err := riskSvc.GetUserRiskLevel(ctx, userID)
fmt.Printf("等级: %d, 分数: %d\n", level, score)

// 更新风险分数（增加50分）
err = riskSvc.UpdateUserRiskScore(ctx, userID, 50, "登录失败3次")

// 更新风险分数（减少10分）
err = riskSvc.UpdateUserRiskScore(ctx, userID, -10, "正常行为")
```

### 6. 策略管理

```go
// 加载场景策略
strategies, code, err := riskSvc.LoadStrategies(ctx, "register")
for _, s := range strategies {
    fmt.Printf("规则: %s, 阈值: %d\n", s.RuleCode, s.Threshold)
}

// 刷新策略缓存
err = riskSvc.RefreshStrategyCache(ctx, "register")
```

### 7. 黑名单管理

```go
// 添加到黑名单
err := riskSvc.AddToBlacklist(ctx, "ip", "192.168.1.100", "恶意攻击")

// 从黑名单移除
err := riskSvc.RemoveFromBlacklist(ctx, "ip", "192.168.1.100")
```

## 🔍 常量定义

### 风控事件类型

```go
RiskEventRegister      = 1 // 注册
RiskEventLoginSuccess  = 2 // 登录成功
RiskEventLoginFail     = 3 // 登录失败
RiskEventDeviceBinding = 4 // 设备绑定
RiskEventScoreChange   = 5 // 分数变化
RiskEventBlacklist     = 6 // 加入黑名单
RiskEventUnlock        = 7 // 解除锁定
```

### 风险等级

```go
RiskLevelNormal   = "normal"   // 正常（0）
RiskLevelObserve  = "observe"  // 观察（1）
RiskLevelRestrict = "restrict" // 限制（2）
RiskLevelBan      = "ban"      // 封禁（3）
```

## 📈 监控建议

### 关键指标

1. **风控拦截率**
   - 注册拦截率（IP/设备）
   - 登录拦截率（黑名单/锁定）
   - 目标：< 5%

2. **用户风险分布**
   - 正常用户占比（目标 > 95%）
   - 观察用户占比（目标 < 3%）
   - 限制/封禁用户占比（目标 < 2%）

3. **缓存性能**
   - 策略缓存命中率（目标 > 90%）
   - 黑名单缓存命中率（目标 > 85%）
   - Redis响应时间（目标 < 10ms）

4. **事件统计**
   - 每日风控事件数
   - 分数变化趋势
   - 黑名单增长趋势

## 🐛 故障排查

### 1. 用户被误封

```bash
# 检查风险分数
redis-cli
> GET "risk:score:123"

# 检查黑名单
SELECT * FROM app_risk_blacklist WHERE type='user' AND value='123';

# 解除封禁
UPDATE app_risk_user SET risk_level=0, risk_score=0 WHERE user_id=123;
```

### 2. 策略不生效

```bash
# 清除策略缓存
> DEL "risk:strategy:register"

# 重新加载
curl -X POST /api/v1/admin/risk/refresh-strategy?scene=register
```

### 3. IP被误封

```bash
# 检查IP黑名单
SELECT * FROM app_risk_blacklist WHERE type='ip' AND value='192.168.1.1';

# 移除IP黑名单
UPDATE app_risk_blacklist SET status='2' WHERE type='ip' AND value='192.168.1.1';

# 清除缓存
> DEL "risk:blacklist:ip:192.168.1.1"
```

## ✅ 已完成功能

- ✅ 三维限流（User/IP/Device）
- ✅ 三维锁定（User/IP/Device）
- ✅ 黑名单检查（含缓存）
- ✅ 风险评分系统
- ✅ 策略配置化
- ✅ 多级缓存（Redis + 数据库缓存表）
- ✅ 事件记录
- ✅ 完整的Redis操作封装
- ✅ 统一常量管理
- ✅ 0个Lint错误

## 🚀 扩展功能（建议）

### 优先级 P0

1. **实时风控规则引擎**
   - Lua脚本实现复杂规则
   - 规则热更新
   - A/B测试支持

2. **风控仪表板**
   - 实时监控面板
   - 风控告警
   - 数据可视化

### 优先级 P1

3. **机器学习模型**
   - 异常行为检测
   - 设备指纹分析
   - 用户画像构建

4. **申诉流程完善**
   - 申诉工单系统
   - 人工审核流程
   - 自动解封机制

### 优先级 P2

5. **地理位置风控**
   - IP地理位置识别
   - 异地登录检测
   - VPN/代理识别

6. **行为分析**
   - 用户行为轨迹
   - 操作频率分析
   - 异常模式识别

## 📝 相关文档

- **设计文档**：`app/auth/README_RISK.md`
- **常量定义**：`config/base/constant/risk.go`
- **数据库Schema**：`app_mysql.sql`

---

**版本：** v1.0  
**最后更新：** 2026-01-07  
**维护者：** Challenge Team
