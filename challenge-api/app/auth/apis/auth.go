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
	s := service.AuthRegister{}
	err := a.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		MakeRuntime().Errors
	if err != nil {
		return
	}
	user, code := s.Register(&req)
	if code != 200 {
		a.Error(code, lang.Get(a.Lang, code))
		return
	}
	c.Set(authdto.UserId, user.ID)
	c.Set(authdto.Username, user.Username)
	auth.Auth.Login(c)
}

func (a *Auth) Login(c *gin.Context) {
	req := dto.LoginReq{}
	s := service.AuthLogin{}
	err := a.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		MakeRuntime().Errors
	if err != nil {
		return
	}
	user, code := s.Login(&req)
	if code != 200 {
		a.Error(code, lang.Get(a.Lang, code))
		return
	}
	c.Set(authdto.UserId, user.ID)
	c.Set(authdto.Username, user.Username)
	auth.Auth.Login(c)
}

func (a *Auth) Logout(c *gin.Context) {
	// 获取用户信息
	userId := c.GetInt64(authdto.UserId)
	username := c.GetString(authdto.Username)

	// 调用登出服务记录日志
	s := service.AuthLogout{}
	err := a.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		MakeRuntime().Errors
	if err != nil {
		a.Error(500, "登出失败")
		return
	}

	// 记录登出日志
	_ = s.Logout(int(userId), username)

	// 调用认证中间件的登出
	auth.Auth.Logout(c)

}
