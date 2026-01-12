package router

import (
	"challenge-admin/app/app/challenge/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerChallengeUserRoute)
}

func registerChallengeUserRoute(r *gin.RouterGroup) {
	api := apis.ChallengeUser{}
	r.GET("/challenge/user", api.ChallengeUserPage)
	r.GET("/challenge/user/export", api.ChallengeUserExport)
	r.POST("/challenge/user", api.ChallengeUserCreate)
	r.PUT("/challenge/user/:id", api.ChallengeUserUpdate)
	r.DELETE("/challenge/user/:id", api.ChallengeUserDelete)
}
