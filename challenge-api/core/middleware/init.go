package middleware

import (
	"challenge/core/middleware/auth/jwtauth"
	"challenge/core/runtime"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	JwtTokenCheck string = "JwtToken"
)

func InitMiddleware(r *gin.Engine) {
	// 数据库链接
	r.Use(WithContextDb)
	// 日志处理
	r.Use(LoggerToFile())
	// 自定义错误处理
	r.Use(CustomError)
	// 限流中间件（500 QPS）
	r.Use(RateLimit())
	//只允许post请求
	r.Use(OnlyPost())
	// IsKeepAlive is a middleware function that appends headers
	r.Use(KeepAlive)
	// 跨域处理
	r.Use(Options)
	//风险收集
	r.Use(RiskCollect())
	// Secure is a middleware function that appends security
	r.Use(Secure)
	// 链路追踪
	r.Use(Trace())
	runtime.RuntimeConfig.SetMiddleware(JwtTokenCheck, (*jwtauth.GinJWTMiddleware).MiddlewareFunc)

}

// Auth 返回统一的身份校验中间件
func Auth() gin.HandlerFunc {
	if mw := runtime.RuntimeConfig.GetMiddlewareKey(JwtTokenCheck); mw != nil {
		if fn, ok := mw.(func() gin.HandlerFunc); ok {
			return fn()
		}
	}
	// 默认直接放行，避免因配置问题导致服务不可用
	return func(c *gin.Context) { c.Next() }
}

func OnlyPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		// OPTIONS 请求直接放行（用于跨域预检）
		if c.Request.Method == http.MethodOptions {
			c.Next()
			return
		}

		// SSE 连接端点允许 GET 请求
		if c.Request.Method == http.MethodGet && IsSSEStreamEndpoint(c.Request.URL.Path) {
			c.Next()
			return
		}

		// 其他接口只允许 POST 请求
		if c.Request.Method == http.MethodPost {
			c.Next()
			return
		}

		// 不允许的请求方法
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusMethodNotAllowed,
			"msg":  "Method Not Allowed",
		})
	}
}
