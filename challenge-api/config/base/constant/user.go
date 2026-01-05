package constant

const (
	UserStatusActive  = "1" // 正常
	UserStatusBlocked = "2" // 异常

	UserTreeSort0 = 0
	UserTreeLevel = 0

	UserPayStatusEnabled  = "1" // 提现启用
	UserPayStatusDisabled = "2" // 提现禁用

	UserTreeLeafNo  = "0" // 非末级
	UserTreeLeafYes = "1" // 末级

	// 登录开关（UserConf.can_login）
	UserCanLoginYes = "1"
	UserCanLoginNo  = "2"

	// 通用状态（1 正常 2 异常/停用）
	GeneralStatusOk    = "1"
	GeneralStatusBlock = "2"

	// 用户行为日志 by_type（1 app 用户，2 后台用户）
	UserOperByTypeApp   = "1"
	UserOperByTypeAdmin = "2"

	// 用户关键行为类型
	UserActionRegister          = "1"  // 注册
	UserActionLogin             = "2"  // 登录
	UserActionLogout            = "3"  // 登出
	UserActionCheckin           = "4"  // 打卡
	UserActionProfileNickname   = "5"  // 修改昵称
	UserActionProfileAvatar     = "6"  // 修改头像
	UserActionSecurityLoginPw   = "7"  // 修改/重置登录密码
	UserActionSecurityPayPw     = "8"  // 修改支付密码
	UserActionBindContact       = "9"  // 绑定联系方式（手机/邮箱）
	UserActionUnbindContact     = "10" // 解绑联系方式
	UserActionWithdrawApply     = "11" // 申请提现
	UserActionWithdrawCancel    = "12" // 取消提现
	UserActionWithdrawApprove   = "13" // 提现审核通过
	UserActionWithdrawReject    = "14" // 提现审核拒绝
	UserActionInviteGenerate    = "15" // 生成邀请码/邀请
	UserActionInviteUse         = "16" // 使用邀请码
	UserActionChallengeJoin     = "17" // 挑战报名
	UserActionChallengeSettleOK = "18" // 挑战结算成功
	UserActionChallengeSettleNG = "19" // 挑战结算失败

	// 行为结果（搭配 action_type 使用）
	UserActionResultSuccess = "成功"
	UserActionResultFail    = "失败"
)
