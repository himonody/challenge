// Package sse SSE实时推送模块
// @Description: SSE (Server-Sent Events) 实时推送功能
package sse

import (
	sseRouter "challenge/app/sse/router"
	"challenge/core/sse"
)

// Init 初始化 SSE 模块
func Init() {
	// 初始化核心 SSE 服务
	sse.InitSSE()
}

// SSERouter 返回 SSE 路由
func SSERouter() []func() {
	return sseRouter.SSERouter()
}
