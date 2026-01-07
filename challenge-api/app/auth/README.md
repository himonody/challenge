# Auth 模块 - 用户认证系统

## 📦 模块概览

完整的用户认证系统，包含注册、登录、登出功能，集成完善的风控体系。

## 📁 目录结构

```
app/auth/
├── apis/
│   └── auth.go              # API 接口层（Register/Login/Logout）
├── service/
│   ├── auth_register.go     # 注册服务（13步完整流程）
│   ├── auth_login.go        # 登录服务（多场景处理）
│   ├── auth_logout.go       # 登出服务（日志记录）
│   ├── risk_check.go        # 风控检查服务（统一入口）
│   └── dto/
│       └── auth.go          # 数据传输对象
├── storage/
│   ├── auth_cache.go        # Redis 缓存操作（250+ 行）
│   ├── prefix.go            # 缓存前缀定义
│   └── README.md            # 存储层使用文档
├── router/
│   ├── auth.go              # 路由配置
│   └── router.go            # 路由注册器
├── init.go                  # 模块初始化
├── README_RISK.md           # 风控系统设计文档
└── README.md                # 本文档
```

## 🚀 核心功能

### 1. 用户注册

**接口：** `POST /api/v1/app/auth/register`

**流程：** 13步完整注册流程

```go
1. 采集风控上下文（IP/UA/DeviceFP）
2. 风控检查（黑名单+IP限流+设备限流）
3. 参数校验（用户名/密码/邀请码/验证码）
4. 验证码校验
5. 用户名唯一性检查
6. 邀请码验证（带锁）
7. 密码加密（bcrypt）
8. 生成推荐码
9. 创建用户（事务）
10. 记录操作日志
11. 记录登录日志（注册即登录）
12. 绑定设备
13. 初始化风控用户
```

**请求示例：**

```json
POST /api/v1/app/auth/register
Content-Type: application/json

{
  "username": "test001",
  "password": "Pass123!",
  "refCode": "ABC123",
  "captchaId": "cap_xxx",
  "captchaCode": "1234"
}

Headers:
X-Device-FP: device_fingerprint_xxx
```

**风控策略：**
- ✅ IP限流：1分钟3次
- ✅ 设备限流：24小时2次
- ✅ 黑名单检查（IP/设备）
- ✅ 邀请码验证

### 2. 用户登录

**接口：** `POST /api/v1/app/auth/login`

**流程：** 多场景智能处理

```go
1. 采集风控上下文
2. 风控检查（黑名单+锁定+失败次数）
3. 参数校验（用户名/密码/验证码）
4. 验证码校验
5. 查询用户
   - 不存在 → 记录失败日志
6. 检查用户状态
   - 已禁用 → 拒绝登录
7. 校验密码
   - 错误 → 记录失败+风控计数+执行锁定策略
   - 成功 → 清除失败计数+记录成功日志+更新登录信息
```

**请求示例：**

```json
POST /api/v1/app/auth/login
Content-Type: application/json

{
  "username": "test001",
  "password": "Pass123!",
  "captchaId": "cap_xxx",
  "captchaCode": "1234"
}

Headers:
X-Device-FP: device_fingerprint_xxx
```

**登录失败策略（三维联动）：**

| 失败次数 | 动作 | 持续时间 |
|---------|------|---------|
| 3次 | 锁定账号 | 5分钟 |
| 4次 | 锁定账号 | 30分钟 |
| 5次+ | 永久封禁 | 永久 |

**多维度累计：**
- **User维度**：账号级封禁
- **IP维度**：识别撞库攻击
- **Device维度**：识别设备异常

### 3. 用户登出

**接口：** `POST /api/v1/app/auth/logout`

**流程：**

```go
1. 获取用户信息（从Token）
2. 清除Redis中的Token
3. 记录登出日志（app_user_login_log）
4. 记录操作日志（app_user_oper_log）
5. 调用认证中间件登出
```

**请求示例：**

```json
POST /api/v1/app/auth/logout
Authorization: Bearer {token}
```

## 📊 数据库表

### app_user_login_log（登录日志）

| 字段 | 类型 | 说明 |
|------|------|------|
| id | bigint | 主键 |
| user_id | bigint | 用户ID |
| login_at | datetime | 登录时间 |
| login_ip | varchar(45) | 登录IP |
| device_fp | varchar(64) | 设备指纹 |
| user_agent | varchar(500) | UA信息 |
| status | tinyint | 状态：1-成功 2-失败 3-风控拦截 4-登出 |
| fail_reason | varchar(255) | 失败原因 |
| created_at | datetime | 创建时间 |

**索引：**
- `idx_user_time` (user_id, login_at)
- `idx_status` (status)
- `idx_ip` (login_ip)

### app_user_oper_log（操作日志）

| 字段 | 类型 | 说明 |
|------|------|------|
| id | int | 主键 |
| user_id | int | 用户ID |
| action_type | char(2) | 操作类型 |
| operate_ip | varchar(45) | 操作IP |
| by_type | char(2) | 操作来源：1-APP 2-后台 |
| status | char(1) | 状态 |
| remark | varchar(255) | 备注 |
| created_at | datetime | 创建时间 |

## 🔐 安全特性

### 1. 密码安全

- ✅ **bcrypt 加密**：单向不可逆
- ✅ **格式限制**：4-12位，支持字母数字特殊字符
- ✅ **强度校验**：可选8-20位强密码（大小写+数字）

### 2. 防暴力破解

- ✅ **失败计数**：15分钟滑动窗口
- ✅ **自动锁定**：3次失败锁5分钟，5次永久封禁
- ✅ **三维联动**：User + IP + Device

### 3. 防刷注册

- ✅ **IP限流**：1分钟3次
- ✅ **设备限流**：24小时2次
- ✅ **邀请码验证**：必须有效邀请码

### 4. 黑名单机制

- ✅ **IP黑名单**：拦截恶意IP
- ✅ **设备黑名单**：拦截恶意设备
- ✅ **用户黑名单**：封禁恶意账号

### 5. 审计追踪

- ✅ **完整日志**：所有操作都有记录
- ✅ **失败原因**：详细记录失败原因
- ✅ **设备追踪**：记录设备指纹和UA

## 🎯 风控体系

### 四层风控模型

```
┌─────────────────────────────────────────┐
│   1. 信号采集层（RiskContext）          │
│   - IP地址、设备指纹、UA、用户ID        │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│   2. 实时拦截层（Redis）                │
│   - 注册限流（IP/设备）                 │
│   - 登录锁定（滑动窗口）                │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│   3. 风险评估层（评分模型）             │
│   - 行为评分、累计分数                  │
│   - 分数映射风险等级                    │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│   4. 人工兜底层（申诉）                 │
│   - 申诉机制、人工审核                  │
└─────────────────────────────────────────┘
```

### 风控检查服务

**文件：** `service/risk_check.go`

**核心方法：**

```go
// 注册风控检查
func (r *RiskCheck) CheckRegisterRisk(ctx, rc) error

// 登录风控检查
func (r *RiskCheck) CheckLoginRisk(ctx, username, rc) error

// 记录注册成功（更新限流计数）
func (r *RiskCheck) RecordRegisterSuccess(ctx, rc) error

// 记录登录成功（清除失败计数）
func (r *RiskCheck) RecordLoginSuccess(ctx, userID, username, rc) error

// 记录登录失败（增加计数+执行策略）
func (r *RiskCheck) RecordLoginFail(ctx, userID, username, rc) error
```

## 💾 Redis 缓存

### 缓存 Key 设计

| Key前缀 | 说明 | 过期时间 | 示例 |
|---------|------|---------|------|
| `auth:login:token` | 登录Token | 7天 | `auth:login:token:123` |
| `auth:login:fail` | 登录失败计数 | 15分钟 | `auth:login:fail:test001` |
| `auth:login:lock` | 登录锁定 | 5分钟/30分钟 | `auth:login:lock:test001` |
| `auth:register:ip` | IP注册限流 | 1分钟 | `auth:register:ip:192.168.1.1` |
| `auth:register:device` | 设备注册限流 | 24小时 | `auth:register:device:xxxxx` |
| `auth:captcha` | 验证码 | 5分钟 | `auth:captcha:cap_123` |

**详细文档：** 参见 `storage/README.md`

## 📝 常量定义

### 操作类型（action_type）

```go
UserActionRegister     = "1"  // 注册
UserActionLogin        = "2"  // 登录成功
UserActionLoginFail    = "22" // 登录失败
UserActionLogout       = "3"  // 登出
```

### 登录状态（login_log.status）

```go
1 = 登录成功
2 = 登录失败
3 = 风控拦截
4 = 用户登出
```

### 失败原因常量

```go
LoginFailReasonPasswordError = "密码错误"
LoginFailReasonUserNotFound  = "用户不存在"
LoginFailReasonRiskBlock     = "风控拦截"
LoginFailReasonLocked        = "登录已锁定"
```

**完整常量：** 参见 `config/base/constant/messages.go`

## 🔧 工具函数

### 字符串验证（core/utils/strutils）

```go
// 验证用户名格式（4-12位）
strutils.IsValidUsername(username string) bool

// 验证密码格式（4-12位）
strutils.IsValidPassword(password string) bool

// 验证强密码（8-20位，大小写+数字）
strutils.IsStrongPassword(password string) bool

// 验证邮箱格式
strutils.VerifyEmailFormat(email string) bool

// 验证手机号
strutils.IsMobile(mobile string) bool
```

## 🎨 使用示例

### 完整的注册流程

```go
// 1. 初始化服务
authApi := &apis.Auth{}
authApi.MakeContext(c).MakeOrm().MakeService().MakeRuntime()

// 2. 调用注册
req := &dto.RegisterReq{
    UserName:     "test001",
    Password:     "Pass123!",
    RefCode:      "ABC123",
    CaptchaId:    "cap_xxx",
    CaptchaCode:  "1234",
}

user, code := authService.Register(req)
if code != 200 {
    // 注册失败处理
    return
}

// 3. 自动登录（设置Token）
c.Set(authdto.UserId, user.ID)
c.Set(authdto.Username, user.Username)
auth.Auth.Login(c)
```

### 完整的登录流程

```go
req := &dto.LoginReq{
    UserName:     "test001",
    Password:     "Pass123!",
    CaptchaId:    "cap_xxx",
    CaptchaCode:  "1234",
}

user, code := authService.Login(req)
if code != 200 {
    // 登录失败处理
    // 已自动记录失败日志和风控计数
    return
}

// 登录成功，设置Token
auth.Auth.Login(c)
```

## 📈 监控指标

### 建议监控的指标

1. **注册相关**
   - 注册成功率
   - 注册风控拦截率
   - IP限流触发次数
   - 设备限流触发次数

2. **登录相关**
   - 登录成功率
   - 登录失败率（按原因分类）
   - 账号锁定次数
   - 风控拦截次数

3. **安全相关**
   - 暴力破解尝试次数
   - 黑名单命中次数
   - 异常IP登录次数
   - 多设备登录次数

## 🐛 故障排查

### 常见问题

**Q1: 用户无法注册，提示"注册过于频繁"**

```bash
# 检查IP限流
redis-cli
> GET "auth:register:ip:192.168.1.1"
> TTL "auth:register:ip:192.168.1.1"

# 手动清除（谨慎使用）
> DEL "auth:register:ip:192.168.1.1"
```

**Q2: 用户登录失败，账号被锁定**

```bash
# 检查锁定状态
> GET "auth:login:lock:test001"
> TTL "auth:login:lock:test001"

# 检查失败次数
> GET "auth:login:fail:test001"

# 手动解锁（谨慎使用）
> DEL "auth:login:lock:test001"
> DEL "auth:login:fail:test001"
```

**Q3: 查看用户登录历史**

```sql
-- 最近20次登录记录
SELECT 
    login_at,
    login_ip,
    device_fp,
    CASE status
        WHEN 1 THEN '成功'
        WHEN 2 THEN '失败'
        WHEN 3 THEN '风控拦截'
        WHEN 4 THEN '登出'
    END as status_desc,
    fail_reason
FROM app_user_login_log
WHERE user_id = 123
ORDER BY login_at DESC
LIMIT 20;
```

## 🚀 性能优化

### 已实现的优化

1. **Redis 缓存**
   - 减少数据库查询
   - 快速风控判断
   - 支持高并发

2. **数据库索引**
   - 所有查询字段已建立索引
   - 复合索引优化查询
   - 避免全表扫描

3. **事务控制**
   - 注册流程使用事务
   - 保证数据一致性
   - 失败自动回滚

4. **异步处理**
   - 风控事件异步记录
   - 不影响主流程性能

## 📚 相关文档

- **风控系统设计**：`README_RISK.md`
- **存储层文档**：`storage/README.md`
- **常量定义**：`config/base/constant/messages.go`
- **数据库Schema**：`app_mysql.sql`

## ✅ 功能清单

### 已完成

- ✅ 用户注册（含风控）
- ✅ 用户登录（多场景处理）
- ✅ 用户登出（含日志）
- ✅ 密码加密（bcrypt）
- ✅ 验证码校验
- ✅ 邀请码系统
- ✅ 注册限流（IP/设备）
- ✅ 登录限制（失败锁定）
- ✅ 黑名单机制
- ✅ 三维风控联动
- ✅ 完整日志记录
- ✅ Redis 缓存
- ✅ 数据库索引
- ✅ 常量统一管理
- ✅ 中文描述规范

### 待扩展（可选）

- ⏳ 手机号注册
- ⏳ 邮箱注册
- ⏳ 第三方登录（微信/Google）
- ⏳ 双因素认证（2FA）
- ⏳ 找回密码
- ⏳ 修改密码
- ⏳ Token 刷新机制
- ⏳ 多设备管理
- ⏳ 异地登录提醒
- ⏳ 登录历史查询API

---

**版本：** v1.0  
**最后更新：** 2026-01-07  
**维护者：** Challenge Team
