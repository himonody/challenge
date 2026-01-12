package router

import (
	"challenge-admin/app/app/user/apis"
	"challenge-admin/core/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerUserAccountLogRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerNoCheckUserAccountLogRouter)
}

// registerUserAccountLogRouter 注册需要认证的路由
func registerUserAccountLogRouter(v1 *gin.RouterGroup) {
	api := apis.UserAccountLog{}
	r := v1.Group("/app/user/user-account-log").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.GET("/export", api.Export)
	}
}

// registerNoCheckUserAccountLogRouter
func registerNoCheckUserAccountLogRouter(v1 *gin.RouterGroup) {}
