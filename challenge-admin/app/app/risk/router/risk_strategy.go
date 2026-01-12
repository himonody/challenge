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
	r.GET("/strategy", api.RiskStrategyPage)
	r.GET("/strategy/export", api.RiskStrategyExport)
	r.POST("/strategy", api.RiskStrategyCreate)
	r.PUT("/strategy/:id", api.RiskStrategyUpdate)
	r.DELETE("/strategy/:id", api.RiskStrategyDelete)
}
