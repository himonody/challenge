package apis

import (
	"challenge/app/auth/service"
	"challenge/app/auth/service/dto"
	"challenge/core/dto/api"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	api.Api
}

func (a *Auth) Register(c *gin.Context) {
	dto := dto.RegisterReq{}
	s := service.Auth{}
	err := a.MakeContext(c).
		MakeOrm().
		Bind(&dto).
		MakeService(&s.Service).
		MakeRuntime().Errors
	if err != nil {
		return
	}
	err = s.Register(&dto)
	if err != nil {
		return
	}

}

func (a *Auth) Login(c *gin.Context) {

}

func (a *Auth) Logout(c *gin.Context) {

}
