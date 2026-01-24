package router

import (
	adminHandler "habit/internal/admin/auth/handler"
	adminService "habit/internal/admin/auth/service"
	appHandler "habit/internal/app/auth/handler"
	appService "habit/internal/app/auth/service"
	"habit/internal/common/handler"
	"habit/internal/repo"
	"habit/pkg/database"
	"habit/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes sets up all application routes (all POST requests)
func SetupRoutes(app *fiber.App) {
	// Initialize repositories
	userRepo := repo.NewUserRepository(database.DB)
	walletRepo := repo.NewWalletRepository(database.DB)
	adminRepo := repo.NewAdminRepository(database.DB)

	// Initialize services
	authService := appService.NewAuthService(userRepo, walletRepo, logger.Logger)
	adminAuthService := adminService.NewAdminAuthService(adminRepo, logger.Logger)

	// Initialize handlers
	healthHandler := handler.NewHealthHandler()
	appAuthHandler := appHandler.NewAuthHandler(authService)
	adminAuthHandler := adminHandler.NewAdminAuthHandler(adminAuthService)

	// API v1 routes
	api := app.Group("/api/v1")

	// Health check (POST only)
	api.Post("/health", healthHandler.HealthCheck)

	// Setup App routes
	SetupAppRoutes(api, appAuthHandler, authService)

	// Setup Admin routes
	SetupAdminRoutes(api, adminAuthHandler, adminAuthService)
}
