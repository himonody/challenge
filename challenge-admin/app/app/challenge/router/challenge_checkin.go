package router

import (
	"challenge-admin/app/app/challenge/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerChallengeCheckinRoute)
}

func registerChallengeCheckinRoute(r *gin.RouterGroup) {
	api := apis.ChallengeCheckin{}
	r.GET("/challenge/checkin", api.CheckinPage)
	r.GET("/challenge/checkin/export", api.CheckinExport)
	r.POST("/challenge/checkin", api.CheckinCreate)
	r.PUT("/challenge/checkin/:id", api.CheckinUpdate)
	r.DELETE("/challenge/checkin/:id", api.CheckinDelete)
}
