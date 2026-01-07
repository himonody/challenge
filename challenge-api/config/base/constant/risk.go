package constant

const (
	// 风控缓存前缀
	RiskStrategyCachePrefix = "challenge:risk:strategy"
	RiskRateLimitPrefix     = "challenge:risk:ratelimit"
	RiskBlacklistPrefix     = "challenge:risk:blacklist"

	// 风控事件类型
	RiskEventRegister      = 1 // 注册
	RiskEventLoginSuccess  = 2 // 登录成功
	RiskEventLoginFail     = 3 // 登录失败
	RiskEventDeviceBinding = 4 // 设备绑定
	RiskEventScoreChange   = 5 // 分数变化
	RiskEventBlacklist     = 6 // 加入黑名单
	RiskEventUnlock        = 7 // 解除锁定

	// 风险等级
	RiskLevelNormal   = "normal"   // 正常
	RiskLevelObserve  = "observe"  // 观察
	RiskLevelRestrict = "restrict" // 限制
	RiskLevelBan      = "ban"      // 封禁
)
