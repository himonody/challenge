package router

import (
	"challenge/app/sse/apis"
	"challenge/app/sse/service"
	"challenge/core/runtime"

	"github.com/gin-gonic/gin"
)

// SSERouter SSE 路由注册
func SSERouter() []func() {
	var routers []func()

	// SSE 路由组
	routers = append(routers, func() {
		r := runtime.RuntimeConfig.GetEngine().(*gin.Engine)
		api := r.Group("/api/v1/sse")

		// 创建 Service
		orm := runtime.RuntimeConfig.GetDbByKey("*")
		run := runtime.RuntimeConfig
		sseService := service.NewSSEService(orm, run, "zh-CN")

		// 创建 APIs
		sseApis := &apis.SSEApis{
			Service: sseService,
		}

		// ========== SSE 连接端点（必须使用 GET）==========
		// 支持多种连接方式
		api.GET("/stream", sseApis.Connect)            // ?id=xxx&user_id=xxx&group=xxx
		api.GET("/stream/:group/:id", sseApis.Connect) // /stream/notifications/user123
		api.GET("/stream/:id", sseApis.Connect)        // /stream/user123

		// ========== 管理接口（改为 POST）==========
		api.POST("/info", sseApis.GetInfo)            // 获取管理器信息
		api.POST("/group/info", sseApis.GetGroupInfo) // 获取分组信息
		api.POST("/disconnect", sseApis.Disconnect)   // 断开客户端连接

		// ========== 消息发送接口 ==========
		api.POST("/send", sseApis.SendMessage)            // 发送消息给指定用户
		api.POST("/send/group", sseApis.SendGroupMessage) // 发送消息到分组
		api.POST("/broadcast", sseApis.Broadcast)         // 广播消息

		// ========== 消息管理接口（改为 POST）==========
		api.POST("/messages/pending", sseApis.GetPendingMessages) // 获取待发送消息
		api.POST("/messages/read", sseApis.MarkRead)              // 标记消息已读
		api.POST("/messages/unread", sseApis.GetUnreadCount)      // 获取未读消息数量

		// ========== 订阅管理接口 ==========
		api.POST("/subscribe", sseApis.Subscribe)            // 订阅分组
		api.POST("/unsubscribe", sseApis.Unsubscribe)        // 取消订阅
		api.POST("/subscriptions", sseApis.GetSubscriptions) // 获取订阅列表
	})

	return routers
}
