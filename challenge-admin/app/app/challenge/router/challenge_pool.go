package router

import (
	"challenge-admin/app/app/challenge/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerChallengePoolRoute)
}

func registerChallengePoolRoute(r *gin.RouterGroup) {
	api := apis.ChallengePool{}
	r.GET("/challenge/pool", api.ChallengePoolPage)
	r.GET("/challenge/pool/export", api.ChallengePoolExport)
	r.POST("/challenge/pool", api.ChallengePoolCreate)
	r.PUT("/challenge/pool/:id", api.ChallengePoolUpdate)
	r.DELETE("/challenge/pool/:id", api.ChallengePoolDelete)
}
