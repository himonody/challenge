package router

import (
	"challenge-admin/app/app/risk/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerRiskDeviceRoute)
}

func registerRiskDeviceRoute(r *gin.RouterGroup) {
	api := apis.RiskDevice{}
	r.GET("/device", api.RiskDevicePage)
	r.GET("/device/export", api.RiskDeviceExport)
}
