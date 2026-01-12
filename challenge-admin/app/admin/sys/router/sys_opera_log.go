package router

import (
	"challenge-admin/app/admin/sys/apis"
	"github.com/gin-gonic/gin"

	"challenge-admin/core/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysOperLogRouter)
}

// 需认证的路由代码
func registerSysOperLogRouter(v1 *gin.RouterGroup) {
	api := apis.SysOperLog{}
	r := v1.Group("/admin/sys/sys-oper-log").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.DELETE("", api.Delete)
		r.GET("/export", api.Export)
	}
}
