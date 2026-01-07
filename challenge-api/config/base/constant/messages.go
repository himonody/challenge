package constant

// ==================== 操作结果描述 ====================

const (
	// 注册相关
	MsgRegisterSuccess = "注册成功"
	MsgRegisterFail    = "注册失败"

	// 登录相关
	MsgLoginSuccess          = "登录成功"
	MsgLoginFail             = "登录失败"
	MsgLoginFailPassword     = "密码错误"
	MsgLoginFailUsername     = "用户名不存在"
	MsgLoginFailUserNotFound = "用户不存在"
	MsgLoginFailCaptcha      = "验证码错误"
	MsgLoginFailStatus       = "账号已被禁用"

	// 登出相关
	MsgLogoutSuccess = "用户主动登出"
	MsgLogout        = "登出成功"

	// 风控拦截
	MsgRiskBlock            = "风控拦截"
	MsgRiskBlacklist        = "黑名单拦截"
	MsgRiskIPLimit          = "IP限流拦截"
	MsgRiskDeviceLimit      = "设备限流拦截"
	MsgRiskLoginLocked      = "登录已锁定"
	MsgRiskLoginFailLimit   = "登录失败次数超限"
	MsgRiskRegisterIPLimit  = "注册IP限流"
	MsgRiskRegisterDevLimit = "注册设备限流"

	// 操作类型描述
	MsgActionRegister = "用户注册"
	MsgActionLogin    = "用户登录"
	MsgActionLogout   = "用户登出"
	MsgActionCheckin  = "打卡签到"

	// 通用
	MsgSuccess = "成功"
	MsgFail    = "失败"
)

// ==================== 登录日志失败原因 ====================

const (
	// app_user_login_log.fail_reason 字段
	LoginFailReasonPasswordError = "密码错误"
	LoginFailReasonUserNotFound  = "用户不存在"
	LoginFailReasonCaptchaError  = "验证码错误"
	LoginFailReasonStatusBlock   = "账号已禁用"
	LoginFailReasonRiskBlock     = "风控拦截"
	LoginFailReasonBlacklist     = "黑名单拦截"
	LoginFailReasonIPLimit       = "IP限流"
	LoginFailReasonDeviceLimit   = "设备限流"
	LoginFailReasonLocked        = "登录已锁定"
	LoginFailReasonTooManyFails  = "失败次数过多"
)

// ==================== 操作日志备注 ====================

const (
	// app_user_oper_log.remark 字段

	// 注册
	OperLogRemarkRegisterSuccess = "注册成功"
	OperLogRemarkRegisterFail    = "注册失败"

	// 登录
	OperLogRemarkLoginSuccess      = "登录成功"
	OperLogRemarkChangeLoginPwd    = "修改登录密码成功"
	OperLogRemarkChangePayPwd      = "修改支付密码成功"
	OperLogRemarkUpdateProfile     = "修改用户资料成功"
	OperLogRemarkLoginFail         = "登录失败"
	OperLogRemarkLoginFailPassword = "密码错误"
	OperLogRemarkLoginFailCaptcha  = "验证码错误"
	OperLogRemarkLoginFailRisk     = "风控拦截"

	// 登出
	OperLogRemarkLogoutSuccess = "用户主动登出"
	OperLogRemarkLogout        = "登出"

	// 修改操作
	OperLogRemarkUpdateNickname = "修改昵称"
	OperLogRemarkUpdateAvatar   = "修改头像"
	OperLogRemarkUpdatePassword = "修改密码"
	OperLogRemarkResetPassword  = "重置密码"

	// 绑定操作
	OperLogRemarkBindMobile   = "绑定手机号"
	OperLogRemarkBindEmail    = "绑定邮箱"
	OperLogRemarkUnbindMobile = "解绑手机号"
	OperLogRemarkUnbindEmail  = "解绑邮箱"

	// 提现操作
	OperLogRemarkWithdrawApply   = "申请提现"
	OperLogRemarkWithdrawCancel  = "取消提现"
	OperLogRemarkWithdrawApprove = "提现审核通过"
	OperLogRemarkWithdrawReject  = "提现审核拒绝"

	// 挑战操作
	OperLogRemarkChallengeJoin     = "加入挑战"
	OperLogRemarkChallengeCheckin  = "挑战打卡"
	OperLogRemarkChallengeSettle   = "挑战结算"
	OperLogRemarkChallengeSettleOK = "挑战结算成功"
	OperLogRemarkChallengeSettleNG = "挑战结算失败"

	// 邀请操作
	OperLogRemarkInviteGenerate = "生成邀请码"
	OperLogRemarkInviteUse      = "使用邀请码"
)

// ==================== 状态描述 ====================

const (
	// 用户状态描述
	StatusDescActive  = "正常"
	StatusDescBlocked = "已禁用"
	StatusDescLocked  = "已锁定"
	StatusDescDeleted = "已删除"

	// 审核状态描述
	StatusDescPending  = "待审核"
	StatusDescApproved = "已通过"
	StatusDescRejected = "已拒绝"

	// 开关状态描述
	StatusDescEnabled  = "已启用"
	StatusDescDisabled = "已禁用"

	// 通用状态
	StatusDescOK    = "正常"
	StatusDescError = "异常"

	// 挑战状态描述
	StatusDescChallengeNone    = "无进行中的挑战"
	StatusDescChallengeUnknown = "未知"
	StatusDescChallengeDoing   = "进行中"
	StatusDescChallengeSuccess = "成功"
	StatusDescChallengeFail    = "失败"
)

// ==================== 登录日志状态描述 ====================

const (
	// app_user_login_log.status 对应的中文描述
	LoginStatusDescSuccess   = "登录成功"
	LoginStatusDescFail      = "登录失败"
	LoginStatusDescRiskBlock = "风控拦截"
	LoginStatusDescLogout    = "用户登出"
)

// GetLoginStatusDesc 根据状态码获取中文描述
func GetLoginStatusDesc(status int8) string {
	switch status {
	case 1:
		return LoginStatusDescSuccess
	case 2:
		return LoginStatusDescFail
	case 3:
		return LoginStatusDescRiskBlock
	case 4:
		return LoginStatusDescLogout
	default:
		return "未知状态"
	}
}

// ==================== 操作类型描述 ====================

const (
	// app_user_oper_log.action_type 对应的中文描述
	ActionTypeDescRegister          = "注册"
	ActionTypeDescLogin             = "登录"
	ActionTypeDescLoginFail         = "登录失败"
	ActionTypeDescLogout            = "登出"
	ActionTypeDescCheckin           = "打卡"
	ActionTypeDescUpdateNickname    = "修改昵称"
	ActionTypeDescUpdateAvatar      = "修改头像"
	ActionTypeDescUpdateLoginPwd    = "修改登录密码"
	ActionTypeDescUpdatePayPwd      = "修改支付密码"
	ActionTypeDescBindContact       = "绑定联系方式"
	ActionTypeDescUnbindContact     = "解绑联系方式"
	ActionTypeDescWithdrawApply     = "申请提现"
	ActionTypeDescWithdrawCancel    = "取消提现"
	ActionTypeDescWithdrawApprove   = "审核通过"
	ActionTypeDescWithdrawReject    = "审核拒绝"
	ActionTypeDescInviteGenerate    = "生成邀请码"
	ActionTypeDescInviteUse         = "使用邀请码"
	ActionTypeDescChallengeJoin     = "加入挑战"
	ActionTypeDescChallengeSettleOK = "结算成功"
	ActionTypeDescChallengeSettleNG = "结算失败"
)

// GetActionTypeDesc 根据操作类型获取中文描述
func GetActionTypeDesc(actionType string) string {
	descMap := map[string]string{
		UserActionRegister:          ActionTypeDescRegister,
		UserActionLogin:             ActionTypeDescLogin,
		UserActionLoginFail:         ActionTypeDescLoginFail,
		UserActionLogout:            ActionTypeDescLogout,
		UserActionCheckin:           ActionTypeDescCheckin,
		UserActionProfileNickname:   ActionTypeDescUpdateNickname,
		UserActionProfileAvatar:     ActionTypeDescUpdateAvatar,
		UserActionSecurityLoginPw:   ActionTypeDescUpdateLoginPwd,
		UserActionSecurityPayPw:     ActionTypeDescUpdatePayPwd,
		UserActionBindContact:       ActionTypeDescBindContact,
		UserActionUnbindContact:     ActionTypeDescUnbindContact,
		UserActionWithdrawApply:     ActionTypeDescWithdrawApply,
		UserActionWithdrawCancel:    ActionTypeDescWithdrawCancel,
		UserActionWithdrawApprove:   ActionTypeDescWithdrawApprove,
		UserActionWithdrawReject:    ActionTypeDescWithdrawReject,
		UserActionInviteGenerate:    ActionTypeDescInviteGenerate,
		UserActionInviteUse:         ActionTypeDescInviteUse,
		UserActionChallengeJoin:     ActionTypeDescChallengeJoin,
		UserActionChallengeSettleOK: ActionTypeDescChallengeSettleOK,
		UserActionChallengeSettleNG: ActionTypeDescChallengeSettleNG,
	}
	if desc, ok := descMap[actionType]; ok {
		return desc
	}
	return "未知操作"
}
