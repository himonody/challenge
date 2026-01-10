package router

import (
	"challenge/core/global"
	"challenge/core/runtime"

	"github.com/gin-gonic/gin"
)

var (
	routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
	routerCheckRole   = make([]func(v1 *gin.RouterGroup), 0)
)

func InitRouter() {
	var r *gin.Engine
	h := runtime.RuntimeConfig.GetEngine()
	if h == nil {
		panic("not found engine...")
	}
	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		panic("not support other engine")
	}

	noCheckRoleRouter(r)
	checkRoleRouter(r)
}

func noCheckRoleRouter(r *gin.Engine) {
	v1 := r.Group(global.RouteRootPath + "/v1")
	for _, f := range routerNoCheckRole {
		f(v1)
	}
}

func checkRoleRouter(r *gin.Engine) {
	v1 := r.Group(global.RouteRootPath + "/v1")
	for _, f := range routerCheckRole {
		f(v1)
	}
}
