package router

import (
	"challenge-admin/app/app/risk/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerRiskRateLimitRoute)
}

func registerRiskRateLimitRoute(r *gin.RouterGroup) {
	api := apis.RiskRateLimit{}
	r.GET("/rate-limit", api.RiskRateLimitPage)
	r.GET("/rate-limit/export", api.RiskRateLimitExport)
}
