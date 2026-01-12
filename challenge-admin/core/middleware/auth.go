package middleware

import (
	"challenge-admin/core/middleware/auth"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return auth.Auth.AuthMiddlewareFunc()
}

func AuthCheckRole() gin.HandlerFunc {
	return auth.Auth.AuthCheckRoleMiddlewareFunc()
}
