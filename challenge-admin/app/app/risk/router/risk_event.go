package router

import (
	"challenge-admin/app/app/risk/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerRiskEventRoute)
}

func registerRiskEventRoute(r *gin.RouterGroup) {
	api := apis.RiskEvent{}
	r.GET("/risk/event", api.RiskEventPage)
	r.GET("/risk/event/export", api.RiskEventExport)
}
