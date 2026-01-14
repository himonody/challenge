package router

import (
	"challenge-admin/app/app/risk/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerRiskActionRoute)
}

func registerRiskActionRoute(r *gin.RouterGroup) {
	api := apis.RiskAction{}
	r.GET("/risk/action", api.RiskActionPage)
	r.GET("/risk/action/export", api.RiskActionExport)
	r.POST("/risk/action", api.RiskActionCreate)
	r.PUT("/risk/action/:code", api.RiskActionUpdate)
	r.DELETE("/risk/action/:code", api.RiskActionDelete)
}
