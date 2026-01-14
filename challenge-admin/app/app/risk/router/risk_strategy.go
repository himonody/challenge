package router

import (
	"challenge-admin/app/app/risk/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerRiskStrategyRoute)
}

func registerRiskStrategyRoute(r *gin.RouterGroup) {
	api := apis.RiskStrategy{}
	r.GET("/risk/strategy", api.RiskStrategyPage)
	r.GET("/risk/strategy/export", api.RiskStrategyExport)
	r.POST("/risk/strategy", api.RiskStrategyCreate)
	r.PUT("/risk/strategy/:id", api.RiskStrategyUpdate)
	r.DELETE("/risk/strategy/:id", api.RiskStrategyDelete)
}
