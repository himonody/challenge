# 8核16G 单机部署优化方案

## 📊 你的部署架构

```
┌─────────────────────────────────────────┐
│         8核16G 单机服务器                │
│                                         │
│  ┌─────────────┐  ┌──────────────┐    │
│  │ Challenge   │  │   MySQL      │    │
│  │   API       │←→│ (3306)       │    │
│  │  (9000)     │  │ max_conn=151 │    │
│  │  QPS: 800   │  └──────────────┘    │
│  └─────────────┘                       │
│         ↓                               │
│  ┌─────────────┐                       │
│  │   Redis     │                       │
│  │  (6379)     │                       │
│  │  pool=100   │                       │
│  └─────────────┘                       │
│                                         │
└─────────────────────────────────────────┘
```

---

## 🎯 资源分配估算

### 内存分配（16G 总内存）

```
系统内核:          2G   (12.5%)
MySQL:            6G   (37.5%)  ← InnoDB buffer pool
Redis:            3G   (18.75%) ← maxmemory
API 应用:         4G   (25%)
预留/缓存:        1G   (6.25%)
──────────────────────────────
总计:            16G   (100%)
```

### CPU 分配（8核）

```
API 应用:      4核 (50%)  ← 处理 HTTP 请求
MySQL:         3核 (37.5%) ← 数据库查询
Redis:         0.5核 (6%)  ← 缓存操作
系统/其他:     0.5核 (6.5%)
──────────────────────────────
总计:          8核 (100%)
```

---

## ⚙️ 推荐配置

### 1️⃣ 限流配置（已更新）✅

```go
// core/middleware/ratelimit.go
var DefaultRateLimitConfig = RateLimitConfig{
    QPS:   800,  // 800 QPS
    Burst: 150,  // 150 突发
}
```

**为什么是 800 QPS？**

```
理论计算:
- 8核 CPU，平均每核处理 100-150 QPS
- 理论上限: 8 × 150 = 1200 QPS

实际考虑（单机部署）:
- MySQL 占用 3 核
- Redis 占用 0.5 核
- 应用可用: 4 核
- 安全系数: 0.7

实际 QPS = 4核 × 150 × 0.7 = 420 QPS（过于保守）

综合评估（数据库IO为主）:
- 假设平均响应时间: 80ms
- 数据库连接数: 100（限制后）
- QPS = 100 / 0.08 × 0.7 ≈ 875

建议设置: 800 QPS（取整，留有余量）
```

---

### 2️⃣ 数据库连接池配置（必须添加）⚠️

```yaml
# config/settings.yml
settings:
  database:
    driver: mysql
    source: root:root@tcp(127.0.0.1:3306)/app_db?charset=utf8&parseTime=True&loc=Local&timeout=1000ms
    
    # ⬇️ 必须添加以下配置
    maxOpenConns: 100      # 最大连接数（MySQL 默认 151，预留 51 给其他服务）
    maxIdleConns: 20       # 最大空闲连接数（20% of maxOpenConns）
    connMaxLifeTime: 3600  # 连接最大生命周期 1 小时
    connMaxIdleTime: 1800  # 空闲连接最大生命周期 30 分钟
```

**为什么 maxOpenConns = 100？**
```
MySQL max_connections: 151（默认）
预留给备份/监控/其他: 20
预留给后台任务: 10
预留给队列消费: 10
预留给 SSE 长连接: 11
──────────────────────────────
可用于 API 请求: 100
```

---

### 3️⃣ MySQL 优化配置

#### 检查当前配置

```bash
# 登录 MySQL
mysql -u root -p

# 查看关键配置
SHOW VARIABLES LIKE 'max_connections';
SHOW VARIABLES LIKE 'innodb_buffer_pool_size';
SHOW VARIABLES LIKE 'thread_cache_size';
```

#### 推荐配置（my.cnf）

```ini
# /etc/mysql/my.cnf 或 /etc/my.cnf

[mysqld]
# 基础配置
max_connections = 200              # 增加最大连接数（从默认 151 增加）
thread_cache_size = 32             # 线程缓存（减少创建线程开销）
table_open_cache = 4096            # 表缓存
max_allowed_packet = 64M           # 最大数据包

# InnoDB 配置（重要！）
innodb_buffer_pool_size = 6G       # InnoDB 缓冲池（内存的 37.5%）
innodb_log_file_size = 512M        # 重做日志大小
innodb_flush_log_at_trx_commit = 2 # 性能优化（牺牲一点安全性）
innodb_flush_method = O_DIRECT     # 避免双重缓冲

# 查询缓存（如果 MySQL < 8.0）
query_cache_type = 1
query_cache_size = 256M

# 慢查询日志
slow_query_log = 1
slow_query_log_file = /var/log/mysql/slow.log
long_query_time = 1                # 记录超过 1 秒的查询

# 连接超时
wait_timeout = 600                 # 10 分钟
interactive_timeout = 600
```

#### 应用配置后重启

```bash
# 检查配置文件语法
mysqld --help --verbose 2>&1 | grep "my.cnf"

# 重启 MySQL
sudo systemctl restart mysql

# 验证配置
mysql -u root -p -e "SHOW VARIABLES LIKE 'innodb_buffer_pool_size';"
```

---

### 4️⃣ Redis 优化配置

#### 检查当前配置

```bash
# 连接 Redis
redis-cli

# 查看内存使用
INFO memory

# 查看连接数
INFO clients
```

#### 推荐配置（redis.conf）

```conf
# /etc/redis/redis.conf

# 内存配置
maxmemory 3gb                      # 最大内存 3G（内存的 18.75%）
maxmemory-policy allkeys-lru       # 内存淘汰策略（LRU）

# 持久化配置（根据需求选择）
# 方案 1: 性能优先（推荐）
save ""                            # 关闭 RDB（提升性能）
appendonly yes                     # 开启 AOF
appendfsync everysec               # 每秒同步（平衡性能和安全）

# 方案 2: 安全优先
save 900 1                         # 15分钟内有1个key变化就保存
save 300 10                        # 5分钟内有10个key变化就保存
save 60 10000                      # 1分钟内有10000个key变化就保存
appendonly yes
appendfsync always                 # 每次写入都同步（最安全，性能较低）

# 网络配置
tcp-backlog 511                    # TCP 连接队列
timeout 0                          # 客户端空闲超时（0=永不超时）
tcp-keepalive 300                  # TCP keepalive

# 性能优化
maxclients 10000                   # 最大客户端连接数
```

#### 应用配置后重启

```bash
# 重启 Redis
sudo systemctl restart redis

# 验证配置
redis-cli CONFIG GET maxmemory
redis-cli CONFIG GET maxmemory-policy
```

---

### 5️⃣ 系统优化（Linux）

#### 文件描述符限制

```bash
# 查看当前限制
ulimit -n

# 临时修改
ulimit -n 65535

# 永久修改
sudo vim /etc/security/limits.conf

# 添加以下内容
* soft nofile 65535
* hard nofile 65535
root soft nofile 65535
root hard nofile 65535

# 重启生效
sudo reboot
```

#### TCP 优化

```bash
# 编辑 sysctl.conf
sudo vim /etc/sysctl.conf

# 添加以下内容
net.core.somaxconn = 65535              # 监听队列最大长度
net.ipv4.tcp_max_syn_backlog = 65535    # SYN 队列最大长度
net.ipv4.tcp_tw_reuse = 1               # 复用 TIME_WAIT 连接
net.ipv4.tcp_fin_timeout = 30           # FIN-WAIT-2 超时时间
net.ipv4.ip_local_port_range = 1024 65535  # 本地端口范围

# 应用配置
sudo sysctl -p
```

---

## 📊 性能测试和验证

### 步骤 1: 压力测试

```bash
# 安装 Apache Bench
sudo apt install apache2-utils  # Ubuntu
brew install httpd              # macOS

# 测试 - 逐步增加并发
# 100 并发
ab -n 10000 -c 100 -p body.json -T application/json \
   http://localhost:9000/api/v1/sse/info

# 200 并发
ab -n 20000 -c 200 -p body.json -T application/json \
   http://localhost:9000/api/v1/sse/info

# 300 并发（观察是否出现错误）
ab -n 30000 -c 300 -p body.json -T application/json \
   http://localhost:9000/api/v1/sse/info
```

**关注指标**:
```
Requests per second: XXX [#/sec]  ← 实际 QPS
Time per request: XX [ms]         ← 平均响应时间
Failed requests: 0                ← 必须为 0
Percentage of requests served within XX ms
  50%  XXX ms  ← 中位数
  90%  XXX ms  ← 90% 请求的响应时间
  99%  XXX ms  ← 99% 请求的响应时间
```

**健康标准**:
```
✅ Failed requests = 0
✅ QPS > 800
✅ P99 < 200ms
✅ P90 < 100ms
```

---

### 步骤 2: 监控数据库

**在压测期间，另开一个终端监控 MySQL**:

```bash
# 监控连接数
watch -n 1 'mysql -u root -proot -e "SHOW STATUS LIKE \"Threads_connected\";"'

# 监控慢查询
watch -n 1 'mysql -u root -proot -e "SHOW STATUS LIKE \"Slow_queries\";"'

# 监控查询数
watch -n 1 'mysql -u root -proot -e "SHOW STATUS LIKE \"Questions\";"'
```

**健康标准**:
```
✅ Threads_connected < 120
✅ Slow_queries 不快速增长
✅ Questions 稳定增长
```

---

### 步骤 3: 监控 Redis

```bash
# 监控 Redis 状态
redis-cli --stat

# 实时监控命令
redis-cli MONITOR | head -n 100

# 查看内存使用
redis-cli INFO memory | grep used_memory_human

# 查看连接数
redis-cli INFO clients | grep connected_clients
```

**健康标准**:
```
✅ used_memory < 3GB
✅ connected_clients < 100
✅ ops/sec 稳定
```

---

### 步骤 4: 监控系统资源

```bash
# CPU 和内存
htop

# 网络连接
watch -n 1 'netstat -an | grep :9000 | wc -l'

# 磁盘 IO
iostat -x 1

# 系统负载
uptime
```

**健康标准**:
```
✅ CPU 使用率 < 80%
✅ 内存使用率 < 85%
✅ 磁盘 IO 等待 < 10%
✅ 系统负载 < CPU 核心数（8）
```

---

## 🎯 分阶段调优策略

### 阶段 1: 初始配置（当前）

```go
QPS:   800
Burst: 150
maxOpenConns: 100
```

**目标**: 稳定运行，收集数据

**监控周期**: 1 周

---

### 阶段 2: 优化后（如果阶段 1 表现良好）

```go
QPS:   1000
Burst: 200
maxOpenConns: 120
```

**前提条件**:
- ✅ 阶段 1 限流触发率 < 1%
- ✅ 数据库连接数峰值 < 80
- ✅ 平均响应时间 < 100ms
- ✅ 压力测试通过

---

### 阶段 3: 极限优化（如果需要）

```go
QPS:   1500
Burst: 300
maxOpenConns: 150
```

**前提条件**:
- ✅ 阶段 2 运行稳定
- ✅ MySQL 配置优化完成
- ✅ Redis 配置优化完成
- ✅ 系统参数优化完成
- ✅ 压力测试通过

**⚠️ 注意**: 此配置接近硬件极限，需要非常仔细的监控和调优。

---

## 📋 完整配置清单

### ✅ 已完成

- [x] 限流配置更新为 800 QPS

### ⚠️ 待完成（重要）

- [ ] **添加数据库连接池配置**（`settings.yml`）
- [ ] **优化 MySQL 配置**（`my.cnf`）
- [ ] **优化 Redis 配置**（`redis.conf`）
- [ ] **系统参数优化**（`ulimit`, `sysctl`）
- [ ] **压力测试验证**
- [ ] **监控告警配置**

---

## 🚨 常见问题和解决方案

### 问题 1: 数据库连接数打满

**症状**:
```sql
ERROR 1040 (HY000): Too many connections
```

**解决方案**:
```bash
# 1. 立即临时增加（不需要重启）
mysql> SET GLOBAL max_connections = 300;

# 2. 永久修改 my.cnf
max_connections = 300

# 3. 重启 MySQL
sudo systemctl restart mysql

# 4. 降低 API QPS（临时方案）
QPS: 800 → 600
```

---

### 问题 2: Redis 内存不足

**症状**:
```
redis> SET key value
(error) OOM command not allowed when used memory > 'maxmemory'
```

**解决方案**:
```bash
# 1. 临时增加内存（不需要重启）
redis-cli CONFIG SET maxmemory 4gb

# 2. 永久修改 redis.conf
maxmemory 4gb

# 3. 重启 Redis
sudo systemctl restart redis

# 4. 清理过期 key
redis-cli --scan --pattern "*" | xargs redis-cli DEL
```

---

### 问题 3: 服务器 CPU 打满

**症状**:
```bash
top
# CPU: 95%+ user, 系统响应变慢
```

**排查步骤**:
```bash
# 1. 查看是哪个进程占用 CPU
top

# 2. 查看 MySQL 慢查询
mysql> SHOW FULL PROCESSLIST;
mysql> SELECT * FROM information_schema.PROCESSLIST WHERE TIME > 1;

# 3. 查看是否有死循环的 goroutine
curl http://localhost:9000/debug/pprof/goroutine

# 4. 降低 QPS
QPS: 800 → 500
```

---

### 问题 4: 服务器内存不足

**症状**:
```bash
free -h
# available memory < 1GB
```

**解决方案**:
```bash
# 1. 查看内存占用
free -h
ps aux --sort=-%mem | head -n 10

# 2. 降低 MySQL buffer pool
innodb_buffer_pool_size = 6G → 4G

# 3. 降低 Redis maxmemory
maxmemory 3gb → 2gb

# 4. 重启服务
sudo systemctl restart mysql redis
```

---

## 🎯 总结

### 核心配置

```yaml
# 限流
QPS: 800
Burst: 150

# 数据库连接池
maxOpenConns: 100
maxIdleConns: 20

# MySQL
max_connections: 200
innodb_buffer_pool_size: 6G

# Redis
maxmemory: 3gb
maxmemory-policy: allkeys-lru
```

### 预期性能

```
✅ QPS: 800（峰值 1000）
✅ 平均响应时间: 80-120ms
✅ P99 响应时间: < 200ms
✅ 并发连接数: 100-150
✅ 数据库连接数: 50-100
✅ CPU 使用率: 60-70%
✅ 内存使用率: 70-80%
```

### 下一步行动

1. **立即**: 添加数据库连接池配置到 `settings.yml`
2. **今天**: 优化 MySQL 和 Redis 配置
3. **明天**: 进行压力测试验证
4. **本周**: 上线并监控 1 周
5. **下周**: 根据监控数据调优

---

**你的 8核16G 单机配置已经很不错了！按照这个方案优化后，应该可以稳定支撑 800 QPS，峰值可以达到 1000+ QPS！** 🚀

**记住**: 配置完成后一定要做压力测试验证！📊
