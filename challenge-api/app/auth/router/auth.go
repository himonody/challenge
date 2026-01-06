package router

import (
	"challenge/app/auth/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerNoCheckAuthRouter)
	routerCheckRole = append(routerCheckRole, registerAuthRouter)
}

func registerNoCheckAuthRouter(v1 *gin.RouterGroup) {
	c := apis.Auth{}
	r := v1.Group("/app/auth")
	{
		r.POST("/register", c.Register)
		//r.POST("/login", service.Login)
	}
}

func registerAuthRouter(v1 *gin.RouterGroup) {
	//r := v1.Group("/app/auth").Use(middleware.Auth())
	{
		//r.POST("/logout", service.Logout)
	}
}
