// Package app
// @Description:路由汇总，如无特殊必要，勿操作本页代码
package app

import (
	authRouter "challenge/app/auth"
	sseRouter "challenge/app/sse"
	userRouter "challenge/app/user"
)

// InitModules 初始化各模块
func InitModules() {
	// 初始化 SSE 模块
	sseRouter.Init()
}

// ChallengeRouter
// @Description: 汇总各大板块接口
// @return []func()
func ChallengeRouter() []func() {
	//初始化路由
	var routers []func()

	//app-应用
	routers = append(routers, authRouter.AuthRouter()...)
	routers = append(routers, userRouter.UserRouter()...)
	routers = append(routers, sseRouter.SSERouter()...) // SSE 路由

	return routers
}
