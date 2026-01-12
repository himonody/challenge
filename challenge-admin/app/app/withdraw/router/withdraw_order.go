package router

import (
	"challenge-admin/app/app/withdraw/apis"
	"challenge-admin/core/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerWithdrawOrderRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerNoCheckWithdrawOrderRouter)
}

// registerWithdrawOrderRouter 注册需要认证的路由
func registerWithdrawOrderRouter(v1 *gin.RouterGroup) {
	api := apis.WithdrawOrder{}
	r := v1.Group("/app/withdraw/order").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.PUT("/:id/status", api.UpdateStatus)
		r.GET("/export", api.Export)
	}
}

// registerNoCheckWithdrawOrderRouter
func registerNoCheckWithdrawOrderRouter(v1 *gin.RouterGroup) {}
