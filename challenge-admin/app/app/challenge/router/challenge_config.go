package router

import (
	"challenge-admin/app/app/challenge/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerChallengeConfigRoute)
}

func registerChallengeConfigRoute(r *gin.RouterGroup) {
	api := apis.ChallengeConfig{}
	r.GET("/challenge/config", api.ConfigPage)
	r.GET("/challenge/config/export", api.ConfigExport)
	r.POST("/challenge/config", api.ConfigCreate)
	r.PUT("/challenge/config/:id", api.ConfigUpdate)
	r.DELETE("/challenge/config/:id", api.ConfigDelete)
}
