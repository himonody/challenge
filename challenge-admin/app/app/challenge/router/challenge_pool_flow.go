package router

import (
	"challenge-admin/app/app/challenge/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerChallengePoolFlowRoute)
}

func registerChallengePoolFlowRoute(r *gin.RouterGroup) {
	api := apis.ChallengePoolFlow{}
	r.GET("/challenge/pool/flow", api.ChallengePoolFlowPage)
	r.GET("/challenge/pool/flow/export", api.ChallengePoolFlowExport)
	r.POST("/challenge/pool/flow", api.ChallengePoolFlowCreate)
	r.PUT("/challenge/pool/flow/:id", api.ChallengePoolFlowUpdate)
	r.DELETE("/challenge/pool/flow/:id", api.ChallengePoolFlowDelete)
}
