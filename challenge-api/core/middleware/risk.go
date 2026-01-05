package middleware

import "github.com/gin-gonic/gin"

type RiskContext struct {
	IP       string
	UA       string
	DeviceFP string
}

func RiskCollect() gin.HandlerFunc {
	return func(c *gin.Context) {
		rc := RiskContext{
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
