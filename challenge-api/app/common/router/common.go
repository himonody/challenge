package router

import (
	"challenge/app/common/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerCommonRouter)
	routerCheckRole = append(routerCheckRole, registerNoCheckCommonRouter)
}

// registerCommonRouter 注册需要认证的路由
func registerCommonRouter(v1 *gin.RouterGroup) {

}

// registerNoCheckCommonRouter
func registerNoCheckCommonRouter(v1 *gin.RouterGroup) {
	api := apis.Common{}
	r := v1.Group("/app/common")
	{

		r.POST("/captcha", api.Captcha)
		r.POST("/upload", api.Upload)

	}
}
