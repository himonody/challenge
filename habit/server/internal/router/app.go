package router

import (
	appAuthHandler "habit/internal/app/auth/handler"
	appAuthService "habit/internal/app/auth/service"
	walletHandler "habit/internal/app/wallet/handler"
	walletService "habit/internal/app/wallet/service"
	"habit/internal/repo"
	"habit/pkg/database"
	"habit/pkg/logger"
	"habit/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

// SetupAppRoutes 设置 App 端路由
func SetupAppRoutes(api fiber.Router, appAuthHdl *appAuthHandler.AuthHandler, authService *appAuthService.AuthService) {
	// Initialize wallet dependencies
	walletRepo := repo.NewWalletRepository(database.DB)
	walletSvc := walletService.NewWalletService(walletRepo, logger.Logger)
	walletHdl := walletHandler.NewWalletHandler(walletSvc)

	// App 路由分组
	app := api.Group("/app")

	// 公开路由
	app.Post("/auth/register", appAuthHdl.Register)
	app.Post("/auth/login", appAuthHdl.Login)

	// 需要认证的路由
	appProtected := app.Group("")
	appProtected.Use(middleware.AuthMiddleware(authService))
	appProtected.Post("/auth/logout", appAuthHdl.Logout)
	appProtected.Post("/auth/me", appAuthHdl.GetUserInfo)
	appProtected.Post("/auth/change-password", appAuthHdl.ChangePassword)
	appProtected.Post("/auth/set-pay-password", appAuthHdl.SetPayPassword)
	appProtected.Post("/auth/update-profile", appAuthHdl.UpdateProfile)

	// 钱包路由
	appProtected.Post("/wallet/info", walletHdl.GetWalletInfo)
}
