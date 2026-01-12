package apis

import (
	"challenge/core/dto/api"
	"challenge/core/dto/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChallengeCheckIn struct {
	api.Api
}

func (a *ChallengeCheckIn) Start(c *gin.Context) {
	s := service.Service{}
	err := a.MakeContext(c).MakeOrm().MakeService(&s).MakeRuntime().Error
	if err != nil {
		a.Logger.Error(err)
		a.Error(http.StatusInternalServerError, "")
		return
	}

}
