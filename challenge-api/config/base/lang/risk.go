package lang

import "challenge/core/lang"

// 风控相关返回码，保留 40xxx 段
const (
	RiskStrategyNotFoundCode    = 40001
	RiskBlacklistHitCode        = 40002
	RiskRegisterIPLimitCode     = 40003
	RiskRegisterDeviceLimitCode = 40004
	RiskLoginLockedCode         = 40005
	RiskLoginFailTooManyCode    = 40006
)

func init() {
	lang.MsgInfo[RiskStrategyNotFoundCode] = "未找到可用风控策略"
	lang.MsgInfo[RiskBlacklistHitCode] = "命中风控黑名单，请联系客服"
	lang.MsgInfo[RiskRegisterIPLimitCode] = "注册过于频繁，请稍后再试"
	lang.MsgInfo[RiskRegisterDeviceLimitCode] = "该设备注册次数已达上限"
	lang.MsgInfo[RiskLoginLockedCode] = "账号已被锁定 %d 秒，请稍后再试"
	lang.MsgInfo[RiskLoginFailTooManyCode] = "登录失败次数过多，请稍后再试"
}
