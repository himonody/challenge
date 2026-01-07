package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// KeepAlive is a middleware function that appends headers
// to prevent the client from caching the HTTP response.
func KeepAlive(c *gin.Context) {
	// SSE 连接端点使用自己的缓存控制策略，这里不设置
	if !IsSSEStreamEndpoint(c.Request.URL.Path) {
		c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
		c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
		c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	}
	c.Next()
}

// IsSSEStreamEndpoint 判断是否是 SSE 连接端点（公开函数）
func IsSSEStreamEndpoint(path string) bool {
	// SSE 连接端点都以 /api/v1/sse/stream 开头
	return len(path) >= 20 && path[:20] == "/api/v1/sse/stream"
}

// Options is a middleware function that appends headers
// for options requests and aborts then exits the middleware
// chain and ends the request.
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

// Secure is a middleware function that appends security
// and resource access headers.
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

	// Also consider adding Content-Security-Policy headers
	// c.Header("Content-Security-Policy", "script-src 'self' https://cdnjs.cloudflare.com")
}
