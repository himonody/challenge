package router

import (
	"challenge-admin/app/app/user/apis"
	"challenge-admin/core/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerUserLoginLogRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerNoCheckUserLoginLogRouter)
}

// registerUserLoginLogRouter 注册需要认证的路由
func registerUserLoginLogRouter(v1 *gin.RouterGroup) {
	api := apis.UserLoginLog{}
	r := v1.Group("/app/user/user-login-log").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/export", api.Export)
	}
}

// registerNoCheckUserLoginLogRouter
func registerNoCheckUserLoginLogRouter(v1 *gin.RouterGroup) {}
