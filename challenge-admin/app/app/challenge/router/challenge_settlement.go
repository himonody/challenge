package router

import (
	"challenge-admin/app/app/challenge/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerChallengeSettlementRoute)
}

func registerChallengeSettlementRoute(r *gin.RouterGroup) {
	api := apis.ChallengeSettlement{}
	r.GET("/challenge/settlement", api.ChallengeSettlementPage)
	r.GET("/challenge/settlement/export", api.ChallengeSettlementExport)
	r.POST("/challenge/settlement", api.ChallengeSettlementCreate)
	r.PUT("/challenge/settlement/:id", api.ChallengeSettlementUpdate)
	r.DELETE("/challenge/settlement/:id", api.ChallengeSettlementDelete)
}
