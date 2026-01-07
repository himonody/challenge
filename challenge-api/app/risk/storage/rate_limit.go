package storage

import (
	baseConstant "challenge/config/base/constant"
	"challenge/core/utils/storage"
	"context"
	"fmt"
	"strconv"
)

// 推荐的过期时间常量（秒）
const (
	DefaultRegisterIPWindow     = 60    // IP注册限流：1分钟
	DefaultRegisterDeviceWindow = 86400 // 设备注册限流：24小时
	DefaultLoginFailWindow      = 900   // 登录失败窗口：15分钟
	DefaultLockDuration5M       = 300   // 锁定：5分钟
	DefaultLockDuration30M      = 1800  // 锁定：30分钟
)

// 缓存Key前缀定义
const (
	// 注册限流前缀
	RegisterIPLimitPrefix     = "challenge:risk:register:ip"     // IP注册限流
	RegisterDeviceLimitPrefix = "challenge:risk:register:device" // 设备注册限流

	// 登录失败计数前缀（三维度）
	LoginUserFailPrefix   = "challenge:risk:login:fail:user"   // 用户登录失败计数
	LoginIPFailPrefix     = "challenge:risk:login:fail:ip"     // IP登录失败计数
	LoginDeviceFailPrefix = "challenge:risk:login:fail:device" // 设备登录失败计数

	// 锁定前缀（三维度）
	LoginUserLockPrefix   = "challenge:risk:login:lock:user"   // 用户登录锁定
	LoginIPLockPrefix     = "challenge:risk:login:lock:ip"     // IP锁定
	LoginDeviceLockPrefix = "challenge:risk:login:lock:device" // 设备锁定
)

// CheckRegisterIPLimit 检查IP注册限流（1分钟3次）
func CheckRegisterIPLimit(ctx context.Context, cache storage.AdapterCache, ip string, window, limit int) (bool, error) {
	val, err := cache.Get(RegisterIPLimitPrefix, ip)
	if err != nil {
		return false, nil // 无记录，允许
	}
	count, _ := strconv.Atoi(val)
	if count >= limit {
		return true, nil // 超限
	}
	return false, nil
}

// IncrRegisterIPLimit 增加IP注册计数
func IncrRegisterIPLimit(ctx context.Context, cache storage.AdapterCache, ip string, window int) error {
	val, err := cache.Get(RegisterIPLimitPrefix, ip)
	if err != nil || val == "" {
		// 第一次，设置为1
		return cache.Set(RegisterIPLimitPrefix, ip, "1", window)
	}
	// 递增
	return cache.Increase(RegisterIPLimitPrefix, ip)
}

// CheckRegisterDeviceLimit 检查设备注册限流（24小时2次）
func CheckRegisterDeviceLimit(ctx context.Context, cache storage.AdapterCache, deviceFP string, window, limit int) (bool, error) {
	val, err := cache.Get(RegisterDeviceLimitPrefix, deviceFP)
	if err != nil {
		return false, nil
	}
	count, _ := strconv.Atoi(val)
	if count >= limit {
		return true, nil
	}
	return false, nil
}

// IncrRegisterDeviceLimit 增加设备注册计数
func IncrRegisterDeviceLimit(ctx context.Context, cache storage.AdapterCache, deviceFP string, window int) error {
	val, err := cache.Get(RegisterDeviceLimitPrefix, deviceFP)
	if err != nil || val == "" {
		return cache.Set(RegisterDeviceLimitPrefix, deviceFP, "1", window)
	}
	return cache.Increase(RegisterDeviceLimitPrefix, deviceFP)
}

// GetLoginFailCount 获取登录失败次数
func GetLoginFailCount(ctx context.Context, cache storage.AdapterCache, dimension, key string) (int, error) {
	var prefix string
	switch dimension {
	case "user":
		prefix = LoginUserFailPrefix
	case "ip":
		prefix = LoginIPFailPrefix
	case "device":
		prefix = LoginDeviceFailPrefix
	default:
		return 0, fmt.Errorf("invalid dimension: %s", dimension)
	}

	val, err := cache.Get(prefix, key)
	if err != nil || val == "" {
		return 0, nil
	}
	return strconv.Atoi(val)
}

// IncrLoginFailCount 增加登录失败计数
func IncrLoginFailCount(ctx context.Context, cache storage.AdapterCache, dimension, key string, window int) error {
	var prefix string
	switch dimension {
	case "user":
		prefix = LoginUserFailPrefix
	case "ip":
		prefix = LoginIPFailPrefix
	case "device":
		prefix = LoginDeviceFailPrefix
	default:
		return fmt.Errorf("invalid dimension: %s", dimension)
	}

	val, err := cache.Get(prefix, key)
	if err != nil || val == "" {
		return cache.Set(prefix, key, "1", window)
	}
	return cache.Increase(prefix, key)
}

// ClearLoginFailCount 清除登录失败计数
func ClearLoginFailCount(ctx context.Context, cache storage.AdapterCache, dimension, key string) error {
	var prefix string
	switch dimension {
	case "user":
		prefix = LoginUserFailPrefix
	case "ip":
		prefix = LoginIPFailPrefix
	case "device":
		prefix = LoginDeviceFailPrefix
	default:
		return fmt.Errorf("invalid dimension: %s", dimension)
	}
	return cache.Del(prefix, key)
}

// LockLoginUser 锁定用户登录
func LockLoginUser(ctx context.Context, cache storage.AdapterCache, username string, seconds int) error {
	return cache.Set(LoginUserLockPrefix, username, "1", seconds)
}

// IsLoginUserLocked 检查用户是否被锁定
func IsLoginUserLocked(ctx context.Context, cache storage.AdapterCache, username string) (bool, error) {
	val, err := cache.Get(LoginUserLockPrefix, username)
	if err != nil || val == "" {
		return false, nil
	}
	return true, nil
}

// GetLockTTL 获取锁定剩余时间（秒）
func GetLockTTL(ctx context.Context, cache storage.AdapterCache, username string) (int, error) {
	// 简化实现，返回固定值
	locked, _ := IsLoginUserLocked(ctx, cache, username)
	if locked {
		return baseConstant.AuthLoginLock5M, nil // 默认5分钟
	}
	return 0, nil
}

// ================ IP/设备锁定 ================

// LockIP 锁定IP
func LockIP(ctx context.Context, cache storage.AdapterCache, ip string, seconds int) error {
	return cache.Set(LoginIPLockPrefix, ip, "1", seconds)
}

// IsIPLocked 检查IP是否被锁定
func IsIPLocked(ctx context.Context, cache storage.AdapterCache, ip string) (bool, error) {
	val, err := cache.Get(LoginIPLockPrefix, ip)
	if err != nil || val == "" {
		return false, nil
	}
	return true, nil
}

// UnlockIP 解除IP锁定
func UnlockIP(ctx context.Context, cache storage.AdapterCache, ip string) error {
	return cache.Del(LoginIPLockPrefix, ip)
}

// LockDevice 锁定设备
func LockDevice(ctx context.Context, cache storage.AdapterCache, deviceFP string, seconds int) error {
	if deviceFP == "" {
		return nil
	}
	return cache.Set(LoginDeviceLockPrefix, deviceFP, "1", seconds)
}

// IsDeviceLocked 检查设备是否被锁定
func IsDeviceLocked(ctx context.Context, cache storage.AdapterCache, deviceFP string) (bool, error) {
	if deviceFP == "" {
		return false, nil
	}
	val, err := cache.Get(LoginDeviceLockPrefix, deviceFP)
	if err != nil || val == "" {
		return false, nil
	}
	return true, nil
}

// UnlockDevice 解除设备锁定
func UnlockDevice(ctx context.Context, cache storage.AdapterCache, deviceFP string) error {
	if deviceFP == "" {
		return nil
	}
	return cache.Del(LoginDeviceLockPrefix, deviceFP)
}

// ================ 批量操作 ================

// ClearAllLoginFails 清除所有维度的登录失败计数
func ClearAllLoginFails(ctx context.Context, cache storage.AdapterCache, username, ip, deviceFP string) error {
	_ = ClearLoginFailCount(ctx, cache, "user", username)
	_ = ClearLoginFailCount(ctx, cache, "ip", ip)
	if deviceFP != "" {
		_ = ClearLoginFailCount(ctx, cache, "device", deviceFP)
	}
	return nil
}

// ================ 辅助函数 ================

// BuildRiskKey 构建风控相关的缓存key
func BuildRiskKey(prefix, identifier string) string {
	return fmt.Sprintf("%s:%s", prefix, identifier)
}

// GetDimensionPrefix 根据维度获取前缀
func GetDimensionPrefix(dimension string) (string, error) {
	switch dimension {
	case "user":
		return LoginUserFailPrefix, nil
	case "ip":
		return LoginIPFailPrefix, nil
	case "device":
		return LoginDeviceFailPrefix, nil
	default:
		return "", fmt.Errorf("invalid dimension: %s", dimension)
	}
}
