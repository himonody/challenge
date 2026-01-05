package apis

import (
	"challenge/app/auth/service"
	"challenge/app/auth/service/dto"
	"challenge/config/lang"
	"challenge/core/dto/api"
	"challenge/core/middleware/auth"
	"challenge/core/middleware/auth/authdto"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	api.Api
}

func (a *Auth) Register(c *gin.Context) {
	req := dto.RegisterReq{}
	s := service.Auth{}
	err := a.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		MakeRuntime().Errors
	if err != nil {
		return
	}
	user, code := s.Register(&req)
	if code != 0 {
		a.Error(code, lang.Get(a.Lang, code))
		return
	}
	c.Set(authdto.UserId, user.ID)
	c.Set(authdto.Username, user.Username)
	auth.Auth.Login(c)
}

func (a *Auth) Login(c *gin.Context) {

}

func (a *Auth) Logout(c *gin.Context) {

}
