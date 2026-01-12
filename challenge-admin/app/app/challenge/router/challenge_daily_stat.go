package router

import (
	"challenge-admin/app/app/challenge/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerChallengeDailyStatRoute)
}

func registerChallengeDailyStatRoute(r *gin.RouterGroup) {
	api := apis.ChallengeDailyStat{}
	r.GET("/challenge/daily_stat", api.ChallengeDailyStatPage)
	r.GET("/challenge/daily_stat/export", api.ChallengeDailyStatExport)
	r.POST("/challenge/daily_stat", api.ChallengeDailyStatCreate)
	r.PUT("/challenge/daily_stat/:statDate", api.ChallengeDailyStatUpdate)
	r.DELETE("/challenge/daily_stat/:statDate", api.ChallengeDailyStatDelete)
}
