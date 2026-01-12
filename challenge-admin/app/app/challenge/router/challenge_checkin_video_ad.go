package router

import (
	"challenge-admin/app/app/challenge/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerChallengeCheckinVideoAdRoute)
}

func registerChallengeCheckinVideoAdRoute(r *gin.RouterGroup) {
	api := apis.ChallengeCheckinVideoAd{}
	r.GET("/challenge/checkin/video_ad", api.CheckinVideoAdPage)
	r.GET("/challenge/checkin/video_ad/export", api.CheckinVideoAdExport)
	r.POST("/challenge/checkin/video_ad", api.CheckinVideoAdCreate)
	r.PUT("/challenge/checkin/video_ad/:id", api.CheckinVideoAdUpdate)
	r.DELETE("/challenge/checkin/video_ad/:id", api.CheckinVideoAdDelete)
}
