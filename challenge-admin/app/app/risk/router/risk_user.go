package router

import (
	"challenge-admin/app/app/risk/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerRiskUserRoute)
}

func registerRiskUserRoute(r *gin.RouterGroup) {
	api := apis.RiskUser{}
	r.GET("/risk/user", api.RiskUserPage)
	r.GET("/risk/user/export", api.RiskUserExport)
}
