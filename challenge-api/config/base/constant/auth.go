package constant

// Auth 认证相关常量
const (
	// Token 过期时间（秒）
	AuthTokenExpire   = 7 * 24 * 3600 // 7天 = 604800秒
	AuthSessionExpire = 30 * 60       // 会话：30分钟 = 1800秒

	// 验证码过期时间
	AuthCaptchaExpire   = 5 * 60  // 图形验证码：5分钟 = 300秒
	AuthSMSCodeExpire   = 10 * 60 // 短信验证码：10分钟 = 600秒
	AuthEmailCodeExpire = 10 * 60 // 邮箱验证码：10分钟 = 600秒

	// 登录失败相关
	AuthLoginFailWindow = 15 * 60 // 失败计数窗口：15分钟 = 900秒
	AuthLoginLock5M     = 5 * 60  // 锁定5分钟 = 300秒
	AuthLoginLock30M    = 30 * 60 // 锁定30分钟 = 1800秒

	// 注册限流
	AuthRegisterIPWindow     = 60        // IP限流窗口：1分钟
	AuthRegisterIPLimit      = 3         // IP限流：1分钟3次
	AuthRegisterDeviceWindow = 24 * 3600 // 设备限流窗口：24小时
	AuthRegisterDeviceLimit  = 2         // 设备限流：24小时2次
)
