package router

import (
	"challenge-admin/app/app/user/apis"
	"challenge-admin/core/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerUserRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerNoCheckUserRouter)
}

// registerUserRouter 注册需要认证的路由
func registerUserRouter(v1 *gin.RouterGroup) {
	api := apis.User{}
	r := v1.Group("/app/user/user").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.GET("/export", api.Export)

		r.POST("/recharge/:id", api.Update)           //人工充值
		r.POST("/deduct/:id", api.Update)             //人工扣款
		r.POST("/reset/password/:id", api.Update)     //重置密码
		r.POST("/reset/pay-password/:id", api.Update) //重置支付密码
	}
}

// registerNoCheckUserRouter
func registerNoCheckUserRouter(v1 *gin.RouterGroup) {}
