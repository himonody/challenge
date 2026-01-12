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
	r.GET("/action", api.RiskActionPage)
	r.GET("/action/export", api.RiskActionExport)
	r.POST("/action", api.RiskActionCreate)
	r.PUT("/action/:code", api.RiskActionUpdate)
	r.DELETE("/action/:code", api.RiskActionDelete)
}
