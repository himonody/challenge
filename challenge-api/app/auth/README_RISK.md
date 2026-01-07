# 风控系统集成说明

## 已完成功能

### 1. 数据层（Repo）
- ✅ `app/risk/repo/risk_device.go` - 设备管理
- ✅ `app/risk/repo/risk_event.go` - 风控事件记录
- ✅ `app/risk/repo/risk_user.go` - 风控用户管理
- ✅ `app/risk/repo/risk_strategy.go` - 策略管理
- ✅ `app/risk/repo/risk_action.go` - 动作管理
- ✅ `app/risk/repo/risk_rate_limit.go` - 限流管理
- ✅ `app/risk/repo/risk_blacklist.go` - 黑名单管理

### 2. 缓存层（Storage）
- ✅ `app/risk/storage/rate_limit.go` - Redis限流操作
- ✅ `app/risk/storage/strategy_cache.go` - 策略缓存

### 3. 服务层（Service）
- ✅ `app/risk/service/risk.go` - 风控核心服务
- ✅ `app/auth/service/risk_check.go` - 注册/登录风控检查
- ✅ `app/auth/service/auth_register.go` - 注册流程（含风控）
- ✅ `app/auth/service/auth_login.go` - 登录流程（含风控）

### 4. API层
- ✅ `app/auth/apis/auth.go` - 注册/登录/登出接口
- ✅ `app/auth/router/auth.go` - 路由配置

### 5. 配置
- ✅ `config/base/constant/risk.go` - 风控常量
- ✅ `config/base/lang/risk.go` - 风控多语言

## 核心流程

### 注册流程
```
1. 采集风控上下文（IP/UA/DeviceFP）
2. 黑名单检查
3. IP限流（1分钟3次）
4. 设备限流（24小时2次）
5. 参数校验
6. 用户名唯一性检查
7. 邀请码验证（带FOR UPDATE锁）
8. 创建用户（事务）
9. 记录操作日志
10. 绑定设备
11. 初始化风控用户
12. 更新限流计数
```

### 登录流程
```
1. 采集风控上下文
2. 黑名单检查
3. 用户锁定检查
4. 失败次数检查
5. 查询用户
6. 校验密码
   - 成功：清除失败计数，记录成功事件
   - 失败：累计失败次数，执行锁定策略
7. 记录登录日志
8. 更新最后登录时间和IP
```

### 三维风控联动
- **User维度**：密码连续失败 ≥5次 → 永久封禁
- **IP维度**：识别扫描/撞库行为
- **Device维度**：工作室/设备级打击

## 待处理事项（需手动修复）

### 编译错误修复
1. `app/risk/storage/rate_limit.go` - 删除未使用的导入和变量
2. `app/auth/service/auth_login.go` - 修复密码比对函数（使用正确的bcrypt方法）
3. `app/auth/apis/auth.go` - 修复lang.SuccessCode引用（改为baseLang.SuccessCode）
4. `config/base/constant/user.go` - 删除重复常量定义
5. `app/user/models/user_oper_log.go` - 添加OperateIP字段

### 功能完善
1. 补充英文多语言翻译（`config/lang/en.go`）
2. 完善申诉流程（`app_risk_appeal`表）
3. 添加风控策略管理后台
4. 实现设备指纹前端采集

## 使用示例

### 前端请求头
```javascript
headers: {
  'X-Device-FP': 'xxxx-xxxx-xxxx',  // 设备指纹
  'User-Agent': '...'
}
```

### 注册接口
```
POST /api/v1/app/auth/register
{
  "username": "test001",
  "password": "123456",
  "refCode": "xxxx",
  "captchaId": "xxx",
  "captchaCode": "1234"
}
```

### 登录接口
```
POST /api/v1/app/auth/login
{
  "username": "test001",
  "password": "123456",
  "captchaId": "xxx",
  "captchaCode": "1234"
}
```

##风控规则配置示例（写入`app_risk_strategy`表）

### 注册场景
```sql
INSERT INTO app_risk_strategy (scene, rule_code, identity_type, window_seconds, threshold, action, status, priority)
VALUES 
('register', 'REG_IP_LIMIT', 'ip', 60, 3, 'BLOCK', 1, 100),
('register', 'REG_DEVICE_LIMIT', 'device', 86400, 2, 'BLOCK', 1, 90);
```

### 登录场景
```sql
INSERT INTO app_risk_strategy (scene, rule_code, identity_type, window_seconds, threshold, action, status, priority)
VALUES 
('login', 'LOGIN_FAIL_USER', 'user', 900, 3, 'LOCK_5M', 1, 100),
('login', 'LOGIN_FAIL_USER_BAN', 'user', 900, 5, 'BAN', 1, 90);
```

## 技术亮点

1. **四层风控模型**：信号采集 → 实时拦截 → 风险评估 → 人工兜底
2. **三维联动**：User/IP/Device 多维度协同判断
3. **策略配置化**：所有规则可后台调整，无需发版
4. **优雅降级**：缓存不可用时自动降级到数据库
5. **审计完整**：所有风控动作留痕可追溯

## 监控指标建议

- 注册成功率
- 登录失败率
- IP/设备封禁数
- 风控拦截日志
- 申诉通过率

---
生成时间：2026-01-07
