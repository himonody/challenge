package common

import commonRouter "challenge/app/common/router"

func CommonRouter() []func() {
	var routers []func()
	routers = append(routers, commonRouter.InitRouter)
	return routers
}
