package middleware

import (
	"challenge/app/risk/service/dto"

	"github.com/gin-gonic/gin"
)

func RiskCollect() gin.HandlerFunc {
	return func(c *gin.Context) {
		rc := dto.RiskContext{
			IP:       c.ClientIP(),
			UA:       c.GetHeader("User-Agent"),
			DeviceFP: c.GetHeader("X-Device-FP"),
		}
		if rc.DeviceFP == "" {
			rc.DeviceFP = "unknown"
		}
		c.Set("risk_ctx", rc)
		c.Next()
	}
}
