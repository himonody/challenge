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
	r.GET("/risk/strategy/cache", api.RiskStrategyCachePage)
	r.GET("/risk/strategy/cache/export", api.RiskStrategyCacheExport)
}
