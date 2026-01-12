package router

import (
	"challenge-admin/app/app/challenge/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerChallengeTotalStatRoute)
}

func registerChallengeTotalStatRoute(r *gin.RouterGroup) {
	api := apis.ChallengeTotalStat{}
	r.GET("/challenge/total_stat", api.ChallengeTotalStatPage)
	r.GET("/challenge/total_stat/export", api.ChallengeTotalStatExport)
	r.POST("/challenge/total_stat", api.ChallengeTotalStatCreate)
	r.PUT("/challenge/total_stat/:id", api.ChallengeTotalStatUpdate)
	r.DELETE("/challenge/total_stat/:id", api.ChallengeTotalStatDelete)
}
