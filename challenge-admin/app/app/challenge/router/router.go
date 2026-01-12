package router

import (
	"challenge-admin/core/global"
	"challenge-admin/core/runtime"

	"github.com/gin-gonic/gin"
)

var (
	routerCheckRole = make([]func(v1 *gin.RouterGroup), 0)
)

// InitRouter 初始化路由
func InitRouter() {
	h := runtime.RuntimeConfig.GetEngine()
	engine, ok := h.(*gin.Engine)
	if !ok {
		panic("not support other engine")
	}
	checkRoleRouter(engine)
}

// 需要认证的路由
func checkRoleRouter(r *gin.Engine) {
	v1 := r.Group(global.RouteRootPath + "/v1")
	for _, f := range routerCheckRole {
		f(v1)
	}
}
