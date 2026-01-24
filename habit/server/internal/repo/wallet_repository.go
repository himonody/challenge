package repo

import (
	"context"
	"errors"

	"habit/internal/model"
	"habit/pkg/cache"
	"habit/pkg/database"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type WalletRepository struct {
	db    *gorm.DB
	cache *cache.Cache
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{
		db:    db,
		cache: cache.NewCache(database.RedisClient),
	}
}

// FindByUserID 根据用户ID查询钱包（带缓存）
func (r *WalletRepository) FindByUserID(userID int64) (*model.AppUserWallet, error) {
	ctx := context.Background()
	cacheKey := cache.UserWalletKey(userID)

	// 尝试从缓存获取
	var wallet model.AppUserWallet
	err := r.cache.Get(ctx, cacheKey, &wallet)
	if err == nil {
		return &wallet, nil
	}

	// 缓存未命中，查询数据库
	if !errors.Is(err, redis.Nil) {
		// 记录缓存错误但继续
	}

	err = r.db.Where("user_id = ?", userID).First(&wallet).Error
	if err != nil {
		return nil, err
	}

	// 存入缓存
	_ = r.cache.Set(ctx, cacheKey, &wallet, cache.UserCacheExpiration)

	return &wallet, nil
}

// Create 创建钱包
func (r *WalletRepository) Create(wallet *model.AppUserWallet) error {
	return r.db.Create(wallet).Error
}

// UpdatePayPassword 更新支付密码
func (r *WalletRepository) UpdatePayPassword(userID int64, payPasswordHash string) error {
	err := r.db.Model(&model.AppUserWallet{}).
		Where("user_id = ?", userID).
		Update("pay_pwd", payPasswordHash).Error

	if err == nil {
		// 清除缓存
		r.invalidateWalletCache(userID)
	}

	return err
}

// invalidateWalletCache 清除钱包缓存
func (r *WalletRepository) invalidateWalletCache(userID int64) {
	ctx := context.Background()
	cacheKey := cache.UserWalletKey(userID)
	_ = r.cache.Delete(ctx, cacheKey)
}
