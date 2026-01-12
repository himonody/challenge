package router

import (
	"challenge-admin/app/app/challenge/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerChallengeCheckinImageRoute)
}

func registerChallengeCheckinImageRoute(r *gin.RouterGroup) {
	api := apis.ChallengeCheckinImage{}
	r.GET("/challenge/checkin/image", api.CheckinImagePage)
	r.GET("/challenge/checkin/image/export", api.CheckinImageExport)
	r.POST("/challenge/checkin/image", api.CheckinImageCreate)
	r.PUT("/challenge/checkin/image/:id", api.CheckinImageUpdate)
	r.DELETE("/challenge/checkin/image/:id", api.CheckinImageDelete)
}
