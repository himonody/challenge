# 限流参数调优指南 - 单机部署

## 📊 你的项目配置分析

### 当前资源配置

```yaml
Redis 连接池: 100
队列并发数: 10
业务类型: 数据库密集型（认证、挑战、订单、消息、风控）
```

### 数据库连接池

当前 `settings.yml` 没有配置 `maxOpenConns`，使用 GORM 默认值：

```go
// GORM 默认配置
MaxOpenConns: 不限制（但实际受 MySQL max_connections 限制）
MaxIdleConns: 2
MySQL 默认 max_connections: 151
```

---

## 🎯 限流大小推荐

### 计算公式

```
QPS = (数据库最大连接数 × 利用率) / 平均请求处理时间

其中:
- 利用率通常取 0.6-0.7（留安全余量）
- 平均请求处理时间需要实测（通常 50-200ms）
```

### 根据服务器配置的推荐值

| 服务器配置 | CPU/内存 | 推荐 QPS | Burst | 说明 |
|-----------|---------|---------|-------|------|
| 🔴 **低配** | 2核4G | **200-300** | 50 | 适合开发/测试环境 |
| 🟡 **中配** | 4核8G | **500-800** | 100 | ✅ **推荐生产环境** |
| 🟢 **高配** | 8核16G+ | **1000-1500** | 200 | 大流量场景 |

### 📌 具体场景建议

#### 1️⃣ 小型应用（日活 < 1万）

```go
var DefaultRateLimitConfig = RateLimitConfig{
    QPS:   200,  // 每秒 200 个请求
    Burst: 50,   // 允许 50 个突发
}
```

**适用场景**:
- 刚上线的产品
- 企业内部系统
- 开发/测试环境

---

#### 2️⃣ 中型应用（日活 1-10万）- ✅ **推荐**

```go
var DefaultRateLimitConfig = RateLimitConfig{
    QPS:   500,  // 每秒 500 个请求
    Burst: 100,  // 允许 100 个突发
}
```

**适用场景**:
- 常规的 C 端应用
- 4核8G 服务器
- 数据库连接数: 50-100

**为什么推荐 500 QPS？**
- ✅ 4核8G 服务器可以稳定支撑
- ✅ 数据库连接数不会打满
- ✅ 有足够的安全余量
- ✅ 可以应对 2-3 倍的突发流量

---

#### 3️⃣ 大型应用（日活 > 10万）

```go
var DefaultRateLimitConfig = RateLimitConfig{
    QPS:   1000,  // 每秒 1000 个请求
    Burst: 200,   // 允许 200 个突发
}
```

**适用场景**:
- 高流量应用
- 8核16G+ 服务器
- 需要配合数据库连接池优化

**⚠️ 需要同步优化**:
```yaml
# settings.yml
database:
  maxOpenConns: 100      # 最大连接数
  maxIdleConns: 20       # 最大空闲连接数
  connMaxLifeTime: 3600  # 连接最大生命周期（秒）
  connMaxIdleTime: 1800  # 空闲连接最大生命周期（秒）
```

---

## 🔬 如何测算你的实际 QPS 上限？

### 方法 1: 压力测试（推荐）

```bash
# 安装 Apache Bench
brew install httpd  # macOS
sudo apt install apache2-utils  # Ubuntu

# 测试一个典型接口（如用户信息查询）
ab -n 10000 -c 50 -p body.json -T application/json \
   http://localhost:9000/api/v1/auth/info

# 关注这些指标:
# - Requests per second: XXX [#/sec] ← 这就是你的 QPS 上限
# - Time per request: XX [ms]        ← 平均响应时间
# - Failed requests: 0               ← 确保为 0
```

**示例输出分析**:

```
Requests per second:    856.32 [#/sec] (mean)
Time per request:       58.390 [ms] (mean)
Failed requests:        0

结论: 
- 当前服务器可以支撑约 850 QPS
- 建议限流设置为 850 × 0.7 = 595，取整为 600 QPS
```

---

### 方法 2: 根据数据库连接数估算

#### 步骤 1: 确定数据库最大连接数

```bash
# 登录 MySQL
mysql -u root -p

# 查看当前连接数
SHOW STATUS LIKE 'Threads_connected';

# 查看最大连接数
SHOW VARIABLES LIKE 'max_connections';
```

**假设结果**:
```
max_connections: 151
```

#### 步骤 2: 分配连接数

```
总连接数: 151
预留给队列/后台任务: 20
预留给 SSE 长连接: 30
可用于 API 请求: 101
```

#### 步骤 3: 测算平均响应时间

```bash
# 在日志中查看平均响应时间
# 或使用 curl 测试
time curl -X POST http://localhost:9000/api/v1/auth/info \
  -H "Content-Type: application/json" \
  -d '{}'

# 多次测试取平均值，假设结果: 80ms
```

#### 步骤 4: 计算 QPS

```
QPS = 可用连接数 / (平均响应时间 / 1000) × 安全系数

假设:
- 可用连接数: 100
- 平均响应时间: 80ms = 0.08s
- 安全系数: 0.7

QPS = 100 / 0.08 × 0.7 = 875

建议设置: 800 QPS（取整）
```

---

## 📊 不同业务类型的 QPS 参考

### 你的项目业务分析

| 业务类型 | 响应时间估算 | QPS 占比 | 说明 |
|---------|------------|---------|------|
| **用户登录/注册** | 100-200ms | 10% | 密码校验、风控检查、数据库写入 |
| **用户信息查询** | 30-50ms | 30% | 简单查询，Redis 缓存 |
| **挑战签到** | 80-150ms | 25% | 数据库写入、统计更新 |
| **消息推送** | 20-50ms | 10% | Redis 操作为主 |
| **订单查询** | 50-100ms | 15% | 多表关联查询 |
| **风控检查** | 10-30ms | 10% | Redis 操作为主 |

**加权平均响应时间**:
```
平均响应时间 = 0.1×150 + 0.3×40 + 0.25×115 + 0.1×35 + 0.15×75 + 0.1×20
             = 15 + 12 + 28.75 + 3.5 + 11.25 + 2
             = 72.5ms ≈ 75ms
```

**估算 QPS**（假设 100 个数据库连接）:
```
QPS = 100 / 0.075 × 0.7 = 933

建议设置: 800-900 QPS
```

---

## 🎯 最终建议

### 方案 1: 保守稳健（推荐新手）

```go
var DefaultRateLimitConfig = RateLimitConfig{
    QPS:   300,  // 保守估计
    Burst: 60,   // 20% 的 QPS
}
```

**优点**:
- ✅ 绝对安全，不会压垮服务器
- ✅ 适合初期上线
- ✅ 出问题概率极低

**缺点**:
- ❌ 可能浪费服务器资源
- ❌ 可能过度限流正常用户

---

### 方案 2: 均衡配置（✅ 强烈推荐）

```go
var DefaultRateLimitConfig = RateLimitConfig{
    QPS:   500,  // 中等配置
    Burst: 100,  // 20% 的 QPS
}
```

**优点**:
- ✅ 平衡性能和安全
- ✅ 适合大多数场景
- ✅ 4核8G 服务器可以稳定支撑
- ✅ 有足够的安全余量

**适用场景**:
- 4核8G 或以上服务器
- 日活 1-10万
- MySQL 连接数: 50-100

**这就是当前代码的默认配置！** ✅

---

### 方案 3: 激进配置（需要压测验证）

```go
var DefaultRateLimitConfig = RateLimitConfig{
    QPS:   1000,  // 激进配置
    Burst: 200,   // 20% 的 QPS
}
```

**前提条件**:
- ✅ 8核16G+ 服务器
- ✅ 数据库连接数优化（100+）
- ✅ 通过压力测试验证
- ✅ 做好监控告警

**⚠️ 必须配置**:
```yaml
# settings.yml - 必须添加
database:
  maxOpenConns: 100
  maxIdleConns: 20
  connMaxLifeTime: 3600
  connMaxIdleTime: 1800
```

---

## 🔧 数据库连接池优化建议

### 当前问题

你的 `settings.yml` 没有配置数据库连接池参数，使用默认值（不安全）。

### 优化配置

```yaml
# config/settings.yml
settings:
  database:
    driver: mysql
    source: root:root@tcp(127.0.0.1:3306)/app_db?charset=utf8&parseTime=True&loc=Local&timeout=1000ms
    # ⬇️ 新增以下配置
    maxOpenConns: 100      # 最大连接数（根据 MySQL max_connections 设置）
    maxIdleConns: 20       # 最大空闲连接数（建议 maxOpenConns 的 20%）
    connMaxLifeTime: 3600  # 连接最大生命周期 1小时（防止连接泄漏）
    connMaxIdleTime: 1800  # 空闲连接最大生命周期 30分钟
```

### 配置说明

| 参数 | 推荐值 | 说明 |
|-----|-------|------|
| `maxOpenConns` | 50-100 | 最大连接数，不要超过 MySQL max_connections 的 70% |
| `maxIdleConns` | 10-20 | 空闲连接数，通常是 maxOpenConns 的 20% |
| `connMaxLifeTime` | 3600 | 连接最大生命周期（秒），防止连接泄漏 |
| `connMaxIdleTime` | 1800 | 空闲连接最大生命周期（秒），释放空闲连接 |

### 不同 QPS 对应的连接池配置

```go
// 300 QPS
maxOpenConns: 50
maxIdleConns: 10

// 500 QPS（推荐）
maxOpenConns: 80
maxIdleConns: 15

// 1000 QPS
maxOpenConns: 120
maxIdleConns: 25
```

---

## 📊 监控和调优

### 1. 监控数据库连接数

```sql
-- 实时监控
SHOW STATUS LIKE 'Threads_connected';
SHOW STATUS LIKE 'Max_used_connections';

-- 查看历史最高连接数
SHOW STATUS LIKE 'Max_used_connections_time';
```

**健康指标**:
```
当前连接数 < 最大连接数 × 0.7  ✅ 健康
当前连接数 > 最大连接数 × 0.8  ⚠️ 警告
当前连接数 > 最大连接数 × 0.9  🔴 危险
```

### 2. 监控限流触发率

```bash
# 查看日志中 429 错误的数量
grep "429" files/logs/*.log | wc -l

# 计算限流触发率
限流触发率 = 429 错误数 / 总请求数

健康标准:
- < 1%   ✅ 健康
- 1-5%   ⚠️ 需要关注
- > 5%   🔴 需要增加 QPS
```

### 3. 监控平均响应时间

```bash
# 在 Gin 日志中查看平均响应时间
# [GIN] 2026/01/07 - 15:30:45 | 200 | 85.234ms | 192.168.1.100 | POST "/api/v1/xxx"
#                                      ^^^^^^^^ 这个就是响应时间
```

**响应时间标准**:
```
< 100ms   ✅ 优秀
100-200ms ✅ 良好
200-500ms ⚠️ 可接受
> 500ms   🔴 需要优化
```

---

## 🎯 分阶段调优策略

### 阶段 1: 上线初期（保守）

```go
QPS:   200
Burst: 50
```

**目标**: 确保系统稳定，收集真实数据

**监控指标**:
- 限流触发率 < 1%
- 数据库连接数 < 50
- 平均响应时间 < 100ms

---

### 阶段 2: 流量增长期（均衡）

```go
QPS:   500
Burst: 100
```

**目标**: 支撑业务增长，保持安全余量

**监控指标**:
- 限流触发率 < 3%
- 数据库连接数 < 80
- 平均响应时间 < 150ms

---

### 阶段 3: 高速增长期（激进）

```go
QPS:   1000
Burst: 200
```

**目标**: 最大化利用服务器资源

**监控指标**:
- 限流触发率 < 5%
- 数据库连接数 < 120
- 平均响应时间 < 200ms

**⚠️ 前提**: 必须完成数据库连接池优化和压力测试

---

## 🚨 故障排查

### 问题 1: 频繁触发限流（429）

**可能原因**:
- QPS 设置太低
- 突发流量太大
- 被恶意攻击

**解决方案**:
```go
// 方案 1: 增加 QPS
QPS:   500 → 800
Burst: 100 → 150

// 方案 2: 增加 Burst（应对突发）
Burst: 100 → 200

// 方案 3: 添加 IP 级别限流（防攻击）
// 需要额外开发
```

---

### 问题 2: 数据库连接数打满

**症状**:
```
Error: Too many connections
当前连接数: 150/151
```

**解决方案**:

```yaml
# 方案 1: 增加 MySQL 最大连接数
# 修改 MySQL 配置文件 my.cnf
[mysqld]
max_connections = 300

# 重启 MySQL
sudo systemctl restart mysql

# 方案 2: 降低 API QPS
QPS: 500 → 300

# 方案 3: 优化连接池配置
database:
  maxOpenConns: 80  # 降低最大连接数
  maxIdleConns: 15
  connMaxLifeTime: 1800  # 缩短连接生命周期
```

---

### 问题 3: 响应时间变慢

**症状**:
```
原来 50ms → 现在 300ms
```

**排查步骤**:

```bash
# 1. 检查数据库慢查询
mysql> SHOW FULL PROCESSLIST;
mysql> SELECT * FROM information_schema.PROCESSLIST WHERE TIME > 1;

# 2. 检查 Redis 延迟
redis-cli --latency

# 3. 检查服务器负载
top
htop

# 4. 检查网络延迟
ping 127.0.0.1
```

**常见原因**:
- 🔴 数据库慢查询（需要加索引）
- 🔴 Redis 内存不足（需要扩容）
- 🔴 服务器 CPU/内存打满（需要扩容）
- 🔴 数据库连接数不足（需要优化）

---

## ✅ 总结和建议

### 快速决策表

| 你的情况 | 推荐 QPS | Burst | 理由 |
|---------|---------|-------|------|
| 刚上线，不确定流量 | 200-300 | 50 | 保守安全 |
| 2核4G 服务器 | 200-300 | 50 | 硬件限制 |
| 4核8G 服务器 | **500** ✅ | **100** | **推荐配置** |
| 8核16G 服务器 | 1000 | 200 | 需要压测验证 |
| 日活 < 5000 | 200 | 50 | 流量较小 |
| 日活 5000-50000 | 500 | 100 | 中等流量 |
| 日活 > 50000 | 1000+ | 200+ | 大流量 |

### 我的最终建议

**如果你不确定，就用默认的 500 QPS！** ✅

```go
var DefaultRateLimitConfig = RateLimitConfig{
    QPS:   500,  // ✅ 推荐
    Burst: 100,  // ✅ 推荐
}
```

**理由**:
1. ✅ 4核8G 服务器可以稳定支撑
2. ✅ 有足够的安全余量（不会压垮服务器）
3. ✅ 有足够的容量（不会过度限流）
4. ✅ 适合大多数中小型应用
5. ✅ 后续可以根据监控数据调整

### 下一步行动

1. **上线运行**，使用默认 500 QPS
2. **监控 1 周**，收集真实数据:
   - 限流触发次数
   - 数据库连接数峰值
   - 平均响应时间
   - QPS 峰值
3. **根据数据调整**:
   - 限流触发率 > 5% → 增加 QPS
   - 数据库连接数 > 80 → 增加连接池
   - 响应时间 > 200ms → 优化慢查询
4. **压力测试**（可选）
5. **持续优化**

---

**记住**: 限流是为了保护系统，不是为了限制用户。合理的限流值应该是：

> **在 99% 的时间里，用户感觉不到限流的存在。** ✨

祝你的项目顺利上线！🚀
