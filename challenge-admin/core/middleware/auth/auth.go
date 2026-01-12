package auth

import (
	"challenge-admin/core/config"
	"challenge-admin/core/middleware/auth/jwtauth"
	"challenge-admin/core/middleware/auth/sessionauth"
	"github.com/gin-gonic/gin"
)

const (
	AuthTypeJwt     = "jwt"
	AuthTypeSession = "session"
)

var Auth AuthInter

type AuthInter interface {
	Init()
	Login(c *gin.Context)
	Logout(c *gin.Context)
	Get(c *gin.Context, key string) (interface{}, int, error)
	GetUserId(c *gin.Context) (int64, int, error)
	GetUserName(c *gin.Context) string
	GetRoleId(c *gin.Context) (int64, int, error)
	GetRoleKey(c *gin.Context) string
	GetDeptId(c *gin.Context) (int64, int, error)
	AuthMiddlewareFunc() gin.HandlerFunc
	AuthCheckRoleMiddlewareFunc() gin.HandlerFunc
}

// InitAuth
// @Description: 初始化
func InitAuth() {
	if config.AuthConfig.Use == AuthTypeJwt {
		Auth = &jwtauth.JwtAuth{}
		Auth.Init()
		return
	} else {
		//默认使用session
		Auth = &sessionauth.SessionAuth{}
		Auth.Init()
		return
	}
}
