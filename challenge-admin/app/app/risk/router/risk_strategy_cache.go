package router

import (
	"challenge-admin/app/app/risk/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerRiskStrategyCacheRoute)
}

func registerRiskStrategyCacheRoute(r *gin.RouterGroup) {
	api := apis.RiskStrategyCache{}
	r.GET("/strategy/cache", api.RiskStrategyCachePage)
	r.GET("/strategy/cache/export", api.RiskStrategyCacheExport)
}
