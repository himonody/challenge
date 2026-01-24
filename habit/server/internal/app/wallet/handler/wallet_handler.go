package handler

import (
	"habit/internal/app/wallet/service"
	"habit/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type WalletHandler struct {
	walletService *service.WalletService
}

func NewWalletHandler(walletService *service.WalletService) *WalletHandler {
	return &WalletHandler{
		walletService: walletService,
	}
}

// GetWalletInfo 获取钱包信息
func (h *WalletHandler) GetWalletInfo(c *fiber.Ctx) error {
	// 获取用户ID
	userID, ok := c.Locals("userID").(int64)
	if !ok {
		return response.Error(c, response.CodeUnauthorized, "unauthorized")
	}

	// 调用 service
	walletInfo, err := h.walletService.GetWalletInfo(userID)
	if err != nil {
		return response.Error(c, response.CodeInternalError, "internal_error")
	}

	return response.Success(c, walletInfo)
}
