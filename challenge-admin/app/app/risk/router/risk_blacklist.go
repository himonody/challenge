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
	r.GET("/risk/blacklist", api.RiskBlacklistPage)
	r.GET("/risk/blacklist/export", api.RiskBlacklistExport)
	r.POST("/risk/blacklist", api.RiskBlacklistCreate)
	r.PUT("/risk/blacklist/:id", api.RiskBlacklistUpdate)
}
