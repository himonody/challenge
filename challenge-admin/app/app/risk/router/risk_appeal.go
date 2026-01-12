package router

import (
	"challenge-admin/app/app/risk/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerRiskAppealRoute)
}

func registerRiskAppealRoute(r *gin.RouterGroup) {
	api := apis.RiskAppeal{}
	r.GET("/appeal", api.RiskAppealPage)
	r.GET("/appeal/export", api.RiskAppealExport)
	r.PUT("/appeal/:id/review", api.RiskAppealReview)
}
