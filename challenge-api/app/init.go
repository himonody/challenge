// Package app
// @Description:路由汇总，如无特殊必要，勿操作本页代码
package app

import (
	authRouter "challenge/app/auth"
	userRouter "challenge/app/user"
)

// ChallengeRouter
// @Description: 汇总各大板块接口
// @return []func()
func ChallengeRouter() []func() {
	//初始化路由
	var routers []func()

	//app-应用
	routers = append(routers, authRouter.AuthRouter()...)
	routers = append(routers, userRouter.UserRouter()...)

	return routers
}
