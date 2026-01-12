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
	r.GET("/user", api.RiskUserPage)
	r.GET("/user/export", api.RiskUserExport)
}
