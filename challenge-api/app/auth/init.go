package app

import (
	authRouter "challenge/app/auth/router"
)

func AuthRouter() []func() {
	var routers []func()
	routers = append(routers, authRouter.InitRouter)
	return routers
}
