package middleware

import (
	"challenge-admin/core/runtime"
	"github.com/gin-gonic/gin"
)

func WithContextDb(c *gin.Context) {
	c.Set("db", runtime.RuntimeConfig.GetDbByKey(c.Request.Host).WithContext(c))
	c.Next()
}
