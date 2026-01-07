package middleware

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var (
	// 全局限流器
	globalLimiter *rate.Limiter
	limiterOnce   sync.Once
)

// RateLimitConfig 限流配置
type RateLimitConfig struct {
	QPS   int // 每秒请求数
	Burst int // 突发流量大小（令牌桶容量）
}

// DefaultRateLimitConfig 默认限流配置
// 当前配置适用于: 8核16G 单机部署（服务+MySQL+Redis 同机）
var DefaultRateLimitConfig = RateLimitConfig{
	QPS:   800, // 800 QPS（8核16G 单机部署推荐值）
	Burst: 150, // 允许150个突发请求
}

// initGlobalLimiter 初始化全局限流器
func initGlobalLimiter() {
	limiterOnce.Do(func() {
		globalLimiter = rate.NewLimiter(
			rate.Limit(DefaultRateLimitConfig.QPS),
			DefaultRateLimitConfig.Burst,
		)
	})
}

// RateLimit 限流中间件（全局限流）
func RateLimit() gin.HandlerFunc {
	initGlobalLimiter()

	return func(c *gin.Context) {
		// SSE 连接端点不进行限流（长连接，不应该计入 QPS）
		if IsSSEStreamEndpoint(c.Request.URL.Path) {
			c.Next()
			return
		}

		// 尝试获取令牌
		if !globalLimiter.Allow() {
			// 限流触发
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code": http.StatusTooManyRequests,
				"msg":  "Too Many Requests - Rate limit exceeded",
			})
			return
		}

		c.Next()
	}
}

// RateLimitWithConfig 自定义配置的限流中间件
func RateLimitWithConfig(config RateLimitConfig) gin.HandlerFunc {
	limiter := rate.NewLimiter(
		rate.Limit(config.QPS),
		config.Burst,
	)

	return func(c *gin.Context) {
		// SSE 连接端点不进行限流
		if IsSSEStreamEndpoint(c.Request.URL.Path) {
			c.Next()
			return
		}

		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code": http.StatusTooManyRequests,
				"msg":  "Too Many Requests - Rate limit exceeded",
			})
			return
		}

		c.Next()
	}
}

// GetLimiterStats 获取限流器统计信息（调试用）
func GetLimiterStats() map[string]interface{} {
	if globalLimiter == nil {
		return map[string]interface{}{
			"initialized": false,
		}
	}

	return map[string]interface{}{
		"initialized": true,
		"qps":         DefaultRateLimitConfig.QPS,
		"burst":       DefaultRateLimitConfig.Burst,
		"tokens":      globalLimiter.Tokens(), // 当前可用令牌数
	}
}
