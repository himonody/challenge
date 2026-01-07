# Challenge API - App 模块完善总结

## 📦 模块结构

```
app/
├── auth/          # 认证模块（注册/登录）
├── risk/          # 风控模块
├── user/          # 用户管理模块
├── challenge/     # 挑战打卡模块
├── withdraw/      # 提现模块
├── message/       # 消息模块
├── common/        # 公共模块
└── init.go        # 路由汇总
```

## ✅ 已完善模块

### 1. Auth 模块（认证）

**文件结构：**
- `apis/auth.go` - 注册/登录/登出接口
- `service/auth_register.go` - 注册服务（含完整风控）
- `service/auth_login.go` - 登录服务（含完整风控）
- `service/risk_check.go` - 风控检查服务
- `service/utils.go` - 工具函数（用户名密码正则）
- `storage/prefix.go` - 缓存key前缀定义
- `router/auth.go` - 路由配置

**核心功能：**
- ✅ 注册流程（13步完整风控）
- ✅ 登录流程（三维联动风控）
- ✅ 登出接口
- ✅ IP限流（1分钟3次）
- ✅ 设备限流（24小时2次）
- ✅ 黑名单检查
- ✅ 登录失败锁定（3次锁5分钟，5次永久封禁）

### 2. Risk 模块（风控）

**数据层（Repo）：**
- `repo/risk_device.go` - 设备管理
- `repo/risk_event.go` - 风控事件
- `repo/risk_user.go` - 风控用户
- `repo/risk_strategy.go` - 策略管理
- `repo/risk_action.go` - 动作管理
- `repo/risk_rate_limit.go` - 限流管理
- `repo/risk_blacklist.go` - 黑名单管理

**缓存层（Storage）：**
- `storage/rate_limit.go` - Redis限流操作
- `storage/strategy_cache.go` - 策略缓存

**服务层：**
- `service/risk.go` - 风控核心服务（策略加载、黑名单检查）
- `service/dto/risk_context.go` - 风控上下文

**模型层（Models）：**
- 9个完整的GORM模型，包含所有索引定义

### 3. User 模块优化

**优化内容：**
- ✅ 添加 `UserLoginLog` 模型（用户登录日志表）
- ✅ 完善 `AppUserOperLog` 模型（添加 OperateIP 字段）
- ✅ 完善 `UserInviteCode` 模型（邀请码管理）

## 🎯 核心风控实现

### 注册风控流程（13步）

```
1. 采集风控上下文（IP/UA/DeviceFP）
2. 黑名单检查
3. IP限流检查（1分钟3次）
4. 设备限流检查（24小时2次）
5. 参数校验（用户名/密码/验证码）
6. 用户名唯一性检查
7. 邀请码验证（带 FOR UPDATE 锁）
8. 密码Hash
9. 生成推荐码
10. 创建用户（事务）
11. 记录操作日志
12. 绑定设备
13. 初始化风控用户
```

### 登录风控流程（三维联动）

```
1. 采集风控上下文
2. 黑名单检查
3. 用户锁定状态检查
4. 失败次数检查
5. 查询用户
6. 校验密码
   - 成功：清除失败计数 + 记录登录日志 + 更新最后登录信息
   - 失败：增加失败计数 + 执行锁定策略
      * 3次失败 → 锁定5分钟
      * 4次失败 → 锁定30分钟
      * 5次失败 → 永久封禁
7. 三维度累计（User/IP/Device）
```

### 三维风控联动

| 维度 | 作用 | 策略 |
|------|------|------|
| **User** | 账号级封禁 | 连续失败5次永久封禁 |
| **IP** | 识别扫描/撞库 | IP维度失败累计（15分钟窗口） |
| **Device** | 工作室打击 | 设备指纹关联多账号检测 |

## 🔧 关键技术点

### 1. 缓存设计
- 使用 `Run.GetCacheAdapter()` 动态获取缓存
- 优雅降级：缓存不可用时直接查库
- Redis key 规范：`risk:场景:维度:标识`

### 2. 数据库设计
- 所有表已建立合适索引
- GORM模型与SQL保持一致
- 支持树形用户关系（parent_id/parent_ids）

### 3. 策略配置化
- 风控规则可通过 `app_risk_strategy` 表配置
- 支持多场景（register/login/withdraw）
- 支持多维度（user/ip/device）
- 支持多动作（LOCK/BAN/SCORE）

### 4. 审计追溯
- 所有风控事件记录到 `app_risk_event`
- 用户操作记录到 `app_user_oper_log`
- 登录记录到 `app_user_login_log`

## 📝 配置文件

### 常量定义
- `config/base/constant/risk.go` - 风控常量（事件类型）
- `config/base/constant/user.go` - 用户常量（操作类型、状态）

### 多语言
- `config/base/lang/risk.go` - 风控错误码定义
- `config/lang/zh.go` - 中文翻译
- `config/lang/en.go` - 英文翻译（待补充）

## 🚀 API 接口

### 注册
```
POST /api/v1/app/auth/register
Content-Type: application/json

{
  "username": "test001",
  "password": "Pass123!",
  "refCode": "xxxx",
  "captchaId": "xxx",
  "captchaCode": "1234"
}

Headers:
X-Device-FP: xxxxxxxx  // 设备指纹（可选）
```

### 登录
```
POST /api/v1/app/auth/login
Content-Type: application/json

{
  "username": "test001",
  "password": "Pass123!",
  "captchaId": "xxx",
  "captchaCode": "1234"
}

Headers:
X-Device-FP: xxxxxxxx  // 设备指纹（可选）
```

### 登出
```
POST /api/v1/app/auth/logout
Authorization: Bearer {token}
```

## 📊 风控规则示例

### 注册限流规则
```sql
INSERT INTO app_risk_strategy (scene, rule_code, identity_type, window_seconds, threshold, action, status, priority)
VALUES 
('register', 'REG_IP_LIMIT', 'ip', 60, 3, 'BLOCK', 1, 100),
('register', 'REG_DEVICE_LIMIT', 'device', 86400, 2, 'BLOCK', 1, 90);
```

### 登录失败规则
```sql
INSERT INTO app_risk_strategy (scene, rule_code, identity_type, window_seconds, threshold, action, status, priority)
VALUES 
('login', 'LOGIN_FAIL_3', 'user', 900, 3, 'LOCK_5M', 1, 100),
('login', 'LOGIN_FAIL_5', 'user', 900, 5, 'BAN', 1, 90);
```

### 动作定义
```sql
-- 已预置以下动作
INSERT INTO app_risk_action VALUES
('LOCK_5M','LOCK',300,'锁定5分钟'),
('LOCK_30M','LOCK',1800,'锁定30分钟'),
('BAN','BAN',0,'永久封禁'),
('SCORE_50','SCORE',50,'增加50风险分'),
('SCORE_80','SCORE',80,'增加80风险分');
```

## ⚠️ 注意事项

### User 模块已存在问题
user 模块部分 service 文件存在导入错误，这些是原有代码问题：
- 缺少 `middleware.DataPermission` 定义
- 缺少 `constant.AccountMobileType` 等常量
- 导入路径错误（`challenge/app/admin/sys/service`）

**建议：** 需要补充完整的 middleware 和 constant 定义。

### 前端需要实现
1. 设备指纹采集（X-Device-FP header）
2. 验证码组件
3. 登录/注册表单
4. 错误提示优化

### 运维监控建议
- 注册成功率监控
- 登录失败率监控
- 风控拦截日志分析
- IP/设备封禁统计
- 申诉处理时效

## 🎨 代码质量

- ✅ 所有代码已 gofmt 格式化
- ✅ Auth 和 Risk 模块无 lint 错误
- ✅ 完整的错误处理
- ✅ 详细的注释说明
- ✅ 统一的命名规范

## 📚 相关文档

- `app/auth/README_RISK.md` - 风控系统详细说明
- `app_mysql.sql` - 完整的数据库schema
- `config/settings.yml` - 系统配置文件

---

**优化完成时间：** 2026-01-07  
**Auth & Risk 模块状态：** ✅ 可直接运行
**User 模块状态：** ⚠️ 需补充middleware定义
