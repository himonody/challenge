package lang

import "challenge/config/base/lang"

var zh = map[int]string{
	lang.SuccessCode:       "操作成功",
	lang.RequestErr:        "请求失败",
	lang.AuthErr:           "状态失效，请重新登录",
	lang.ForbitErr:         "对不起，您权限不足，操作异常，请联系管理员",
	lang.ServerErr:         "内部错误",
	lang.ParamErrCode:      "参数错误",
	lang.OpErrCode:         "操作异常，请检查",
	lang.DataDecodeCode:    "数据解析异常",
	lang.DataDecodeLogCode: "数据解析异常：%s",
	lang.DataQueryCode:     "数据查询失败",
	lang.DataQueryLogCode:  "数据查询失败：%s",
	lang.DataInsertLogCode: "数据新增失败：%s",
	lang.DataInsertCode:    "数据新增失败",
	lang.DataNotUpdateCode: "数据未变更",
	lang.DataUpdateCode:    "数据更新异常",
	lang.DataUpdateLogCode: "数据更新异常：%s",
	lang.DataDeleteCode:    "数据删除失败",
	lang.DataDeleteLogCode: "数据删除失败：%s",
	lang.DataNotFoundCode:  "数据不存在",
	lang.ServerErrLogCode:  "内部错误：%s",

	lang.AuthUsernameErrorCode:           "用户名格式错误",
	lang.AuthPasswordErrorCode:           "密码格式错误",
	lang.AuthVerificationCodeErrorCode:   "验证码错误",
	lang.AuthUserAlreadyExistsCode:       "用户已存在",
	lang.AuthInviteCodeNotFoundErrorCode: "推荐码不存在",

	// 风控
	lang.RiskStrategyNotFoundCode: "未找到可用风控策略",
	lang.RiskBlacklistHitCode:     "命中风控黑名单",
	lang.RepeatOperationCode:      "正在处理中,请勿重复操作",
}
