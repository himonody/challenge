package storage

import (
	"challenge/core/utils/storage"
	"context"
	"fmt"
	"strconv"
)

// 推荐的过期时间常量（也可使用 config/base/constant/auth.go 中的定义）
const (
	DefaultTokenExpire    = 7 * 24 * 3600 // Token：7天
	DefaultSessionExpire  = 30 * 60       // Session：30分钟
	DefaultCaptchaExpire  = 5 * 60        // 验证码：5分钟
	DefaultSMSExpire      = 10 * 60       // 短信码：10分钟
	DefaultFailWindow     = 15 * 60       // 失败窗口：15分钟
	DefaultLock5M         = 5 * 60        // 锁定：5分钟
	DefaultLock30M        = 30 * 60       // 锁定：30分钟
	DefaultRegisterIPExp  = 60            // IP限流：1分钟
	DefaultRegisterDevExp = 24 * 3600     // 设备限流：24小时
)

const (
	// 登录相关缓存前缀
	AuthLoginFailPrefix    = "challenge:auth:login:fail"    // 登录失败计数
	AuthLoginLockPrefix    = "challenge:auth:login:lock"    // 登录锁定
	AuthLoginSessionPrefix = "challenge:auth:login:session" // 登录会话

	// 注册相关缓存前缀
	AuthRegisterIPPrefix     = "challenge:auth:register:ip"     // IP注册限流
	AuthRegisterDevicePrefix = "challenge:auth:register:device" // 设备注册限流

	// 验证码相关
	AuthCaptchaPrefix   = "challenge:auth:captcha"    // 验证码
	AuthSMSCodePrefix   = "challenge:auth:sms:code"   // 短信验证码
	AuthEmailCodePrefix = "challenge:auth:email:code" // 邮箱验证码
)

// ================ 登录失败计数管理 ================

// GetLoginFailCount 获取登录失败次数
func GetLoginFailCount(ctx context.Context, cache storage.AdapterCache, username string) (int, error) {
	val, err := cache.Get(AuthLoginFailPrefix, username)
	if err != nil || val == "" {
		return 0, nil
	}
	return strconv.Atoi(val)
}

// IncrLoginFailCount 增加登录失败次数
func IncrLoginFailCount(ctx context.Context, cache storage.AdapterCache, username string, expireSec int) error {
	val, err := cache.Get(AuthLoginFailPrefix, username)
	if err != nil || val == "" {
		// 第一次失败，设置为1
		return cache.Set(AuthLoginFailPrefix, username, "1", expireSec)
	}
	// 递增
	return cache.Increase(AuthLoginFailPrefix, username)
}

// ClearLoginFailCount 清除登录失败次数（登录成功后调用）
func ClearLoginFailCount(ctx context.Context, cache storage.AdapterCache, username string) error {
	return cache.Del(AuthLoginFailPrefix, username)
}

// ================ 登录锁定管理 ================

// LockLogin 锁定登录
func LockLogin(ctx context.Context, cache storage.AdapterCache, username string, lockSeconds int) error {
	return cache.Set(AuthLoginLockPrefix, username, "1", lockSeconds)
}

// IsLoginLocked 检查是否被锁定
func IsLoginLocked(ctx context.Context, cache storage.AdapterCache, username string) (bool, error) {
	val, err := cache.Get(AuthLoginLockPrefix, username)
	if err != nil || val == "" {
		return false, nil
	}
	return true, nil
}

// UnlockLogin 解除登录锁定
func UnlockLogin(ctx context.Context, cache storage.AdapterCache, username string) error {
	return cache.Del(AuthLoginLockPrefix, username)
}

// ================ 登录Session管理 ================

// SetLoginSession 设置登录会话信息
func SetLoginSession(ctx context.Context, cache storage.AdapterCache, sessionID string, userInfo map[string]interface{}, expireSec int) error {
	return cache.HashSet(expireSec, AuthLoginSessionPrefix, sessionID, userInfo)
}

// GetLoginSession 获取登录会话信息
func GetLoginSession(ctx context.Context, cache storage.AdapterCache, sessionID, field string) (string, error) {
	return cache.HashGet(AuthLoginSessionPrefix, sessionID, field)
}

// GetLoginSessionAll 获取完整会话信息
func GetLoginSessionAll(ctx context.Context, cache storage.AdapterCache, sessionID string) (map[string]string, error) {
	return cache.HashGetAll(AuthLoginSessionPrefix, sessionID)
}

// DelLoginSession 删除登录会话
func DelLoginSession(ctx context.Context, cache storage.AdapterCache, sessionID string) error {
	return cache.Del(AuthLoginSessionPrefix, sessionID)
}

// ================ 注册限流管理 ================

// CheckRegisterIPLimit 检查IP注册限流
func CheckRegisterIPLimit(ctx context.Context, cache storage.AdapterCache, ip string, limit int) (bool, error) {
	val, err := cache.Get(AuthRegisterIPPrefix, ip)
	if err != nil || val == "" {
		return false, nil // 未达到限制
	}
	count, _ := strconv.Atoi(val)
	return count >= limit, nil // 超过限制返回true
}

// IncrRegisterIPCount 增加IP注册计数
func IncrRegisterIPCount(ctx context.Context, cache storage.AdapterCache, ip string, expireSec int) error {
	val, err := cache.Get(AuthRegisterIPPrefix, ip)
	if err != nil || val == "" {
		return cache.Set(AuthRegisterIPPrefix, ip, "1", expireSec)
	}
	return cache.Increase(AuthRegisterIPPrefix, ip)
}

// CheckRegisterDeviceLimit 检查设备注册限流
func CheckRegisterDeviceLimit(ctx context.Context, cache storage.AdapterCache, deviceFP string, limit int) (bool, error) {
	if deviceFP == "" {
		return false, nil
	}
	val, err := cache.Get(AuthRegisterDevicePrefix, deviceFP)
	if err != nil || val == "" {
		return false, nil
	}
	count, _ := strconv.Atoi(val)
	return count >= limit, nil
}

// IncrRegisterDeviceCount 增加设备注册计数
func IncrRegisterDeviceCount(ctx context.Context, cache storage.AdapterCache, deviceFP string, expireSec int) error {
	if deviceFP == "" {
		return nil
	}
	val, err := cache.Get(AuthRegisterDevicePrefix, deviceFP)
	if err != nil || val == "" {
		return cache.Set(AuthRegisterDevicePrefix, deviceFP, "1", expireSec)
	}
	return cache.Increase(AuthRegisterDevicePrefix, deviceFP)
}

// ================ 验证码管理 ================

// SetCaptcha 设置图形验证码
func SetCaptcha(ctx context.Context, cache storage.AdapterCache, captchaID, code string, expireSec int) error {
	return cache.Set(AuthCaptchaPrefix, captchaID, code, expireSec)
}

// GetCaptcha 获取图形验证码
func GetCaptcha(ctx context.Context, cache storage.AdapterCache, captchaID string) (string, error) {
	return cache.Get(AuthCaptchaPrefix, captchaID)
}

// DelCaptcha 删除验证码（验证后删除）
func DelCaptcha(ctx context.Context, cache storage.AdapterCache, captchaID string) error {
	return cache.Del(AuthCaptchaPrefix, captchaID)
}

// SetSMSCode 设置短信验证码
func SetSMSCode(ctx context.Context, cache storage.AdapterCache, mobile, code string, expireSec int) error {
	return cache.Set(AuthSMSCodePrefix, mobile, code, expireSec)
}

// GetSMSCode 获取短信验证码
func GetSMSCode(ctx context.Context, cache storage.AdapterCache, mobile string) (string, error) {
	return cache.Get(AuthSMSCodePrefix, mobile)
}

// DelSMSCode 删除短信验证码
func DelSMSCode(ctx context.Context, cache storage.AdapterCache, mobile string) error {
	return cache.Del(AuthSMSCodePrefix, mobile)
}

// SetEmailCode 设置邮箱验证码
func SetEmailCode(ctx context.Context, cache storage.AdapterCache, email, code string, expireSec int) error {
	return cache.Set(AuthEmailCodePrefix, email, code, expireSec)
}

// GetEmailCode 获取邮箱验证码
func GetEmailCode(ctx context.Context, cache storage.AdapterCache, email string) (string, error) {
	return cache.Get(AuthEmailCodePrefix, email)
}

// DelEmailCode 删除邮箱验证码
func DelEmailCode(ctx context.Context, cache storage.AdapterCache, email string) error {
	return cache.Del(AuthEmailCodePrefix, email)
}

// ================ 辅助函数 ================

// BuildUserKey 构建用户相关的缓存key
func BuildUserKey(userID int) string {
	return fmt.Sprintf("user:%d", userID)
}

// BuildIPKey 构建IP相关的缓存key
func BuildIPKey(ip string) string {
	return fmt.Sprintf("ip:%s", ip)
}

// BuildDeviceKey 构建设备相关的缓存key
func BuildDeviceKey(deviceFP string) string {
	return fmt.Sprintf("device:%s", deviceFP)
}
