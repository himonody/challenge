package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"challenge/core/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerUserRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerNoCheckUserRouter)
}

// registerUserRouter 注册需要认证的路由
func registerUserRouter(v1 *gin.RouterGroup) {

	r := v1.Group("/app/user").Use(middleware.Auth())
	{
		r.POST("/profile", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "success",
			})
		})
		r.POST("/profile/update")
		r.POST("/profile/detail")
		r.POST("/profile/stat")

	}
}

// registerNoCheckUserRouter
func registerNoCheckUserRouter(v1 *gin.RouterGroup) {}
