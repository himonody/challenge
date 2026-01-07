package storage

import (
	"context"
	"fmt"

	"challenge/core/utils/storage"
)

// 通用过期时间（秒）
const (
	DefaultUserCacheExpire = 5 * 60 // 5分钟
)

// 缓存前缀
const (
	UserProfilePrefix = "challenge:user:profile" // 用户资料缓存
	UserStatPrefix    = "challenge:user:stat"    // 用户统计缓存
	UserInvitePrefix  = "challenge:user:invite"  // 用户邀请统计缓存
)

// ---------- 用户资料缓存 ----------

// SetUserProfileCache 缓存用户资料（Hash）
func SetUserProfileCache(ctx context.Context, cache storage.AdapterCache, userID uint64, data map[string]interface{}, expireSec int) error {
	if expireSec <= 0 {
		expireSec = DefaultUserCacheExpire
	}
	key := fmt.Sprintf("%d", userID)
	return cache.HashSet(expireSec, UserProfilePrefix, key, data)
}

// GetUserProfileCache 获取用户资料（Hash）
func GetUserProfileCache(ctx context.Context, cache storage.AdapterCache, userID uint64) (map[string]string, error) {
	key := fmt.Sprintf("%d", userID)
	return cache.HashGetAll(UserProfilePrefix, key)
}

// DelUserProfileCache 删除用户资料缓存
func DelUserProfileCache(ctx context.Context, cache storage.AdapterCache, userID uint64) error {
	key := fmt.Sprintf("%d", userID)
	return cache.Del(UserProfilePrefix, key)
}

// ---------- 用户统计缓存 ----------

// SetUserStatCache 缓存用户统计（Hash）
func SetUserStatCache(ctx context.Context, cache storage.AdapterCache, userID uint64, data map[string]interface{}, expireSec int) error {
	if expireSec <= 0 {
		expireSec = DefaultUserCacheExpire
	}
	key := fmt.Sprintf("%d", userID)
	return cache.HashSet(expireSec, UserStatPrefix, key, data)
}

// GetUserStatCache 获取用户统计（Hash）
func GetUserStatCache(ctx context.Context, cache storage.AdapterCache, userID uint64) (map[string]string, error) {
	key := fmt.Sprintf("%d", userID)
	return cache.HashGetAll(UserStatPrefix, key)
}

// DelUserStatCache 删除用户统计缓存
func DelUserStatCache(ctx context.Context, cache storage.AdapterCache, userID uint64) error {
	key := fmt.Sprintf("%d", userID)
	return cache.Del(UserStatPrefix, key)
}

// ---------- 用户邀请统计缓存 ----------

// SetUserInviteCache 缓存用户邀请统计（Hash）
func SetUserInviteCache(ctx context.Context, cache storage.AdapterCache, userID uint64, data map[string]interface{}, expireSec int) error {
	if expireSec <= 0 {
		expireSec = DefaultUserCacheExpire
	}
	key := fmt.Sprintf("%d", userID)
	return cache.HashSet(expireSec, UserInvitePrefix, key, data)
}

// GetUserInviteCache 获取用户邀请统计（Hash）
func GetUserInviteCache(ctx context.Context, cache storage.AdapterCache, userID uint64) (map[string]string, error) {
	key := fmt.Sprintf("%d", userID)
	return cache.HashGetAll(UserInvitePrefix, key)
}

// DelUserInviteCache 删除用户邀请统计缓存
func DelUserInviteCache(ctx context.Context, cache storage.AdapterCache, userID uint64) error {
	key := fmt.Sprintf("%d", userID)
	return cache.Del(UserInvitePrefix, key)
}
