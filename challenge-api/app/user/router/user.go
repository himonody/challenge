package router

import (
	"challenge/app/user/apis"
	"challenge/core/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerUserRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerNoCheckUserRouter)
}

// registerUserRouter 注册需要认证的路由
func registerUserRouter(v1 *gin.RouterGroup) {
	userApis := apis.User{}

	r := v1.Group("/app/user").Use(middleware.Auth())
	{
		// 用户资料相关
		r.POST("/profile", userApis.GetProfile)               // 获取用户资料
		r.POST("/profile/update", userApis.UpdateProfile)     // 修改用户资料
		r.POST("/pwd/change", userApis.ChangeLoginPassword)   // 修改登录密码
		r.POST("/pay-pwd/change", userApis.ChangePayPassword) // 修改支付密码

		// 邀请相关
		r.POST("/invite/info", userApis.GetInviteInfo) // 邀请好友
		r.POST("/invites", userApis.GetMyInvites)      // 我的邀请

		// 统计相关
		r.POST("/stats", userApis.GetStatistics)            // 统计
		r.POST("/stats/today", userApis.GetTodayStatistics) // 今日统计
	}
}

// registerNoCheckUserRouter 注册不需要认证的路由
func registerNoCheckUserRouter(v1 *gin.RouterGroup) {}
