package router

import (
	"challenge-admin/app/app/risk/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerRiskBlacklistRoute)
}

func registerRiskBlacklistRoute(r *gin.RouterGroup) {
	api := apis.RiskBlacklist{}
	r.GET("/blacklist", api.RiskBlacklistPage)
	r.GET("/blacklist/export", api.RiskBlacklistExport)
	r.POST("/blacklist", api.RiskBlacklistCreate)
	r.PUT("/blacklist/:id", api.RiskBlacklistUpdate)
}
