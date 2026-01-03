package router

import (
	"challenge/app/auth/service"
	"challenge/core/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerNoCheckAuthRouter)
	routerCheckRole = append(routerCheckRole, registerAuthRouter)
}

func registerNoCheckAuthRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/app/auth")
	{
		r.POST("/captcha", service.Captcha)
		r.POST("/register", service.Register)
		r.POST("/login", service.Login)
	}
}

func registerAuthRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/app/auth").Use(middleware.Auth())
	{
		r.POST("/logout", service.Logout)
	}
}
