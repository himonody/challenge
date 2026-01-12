package dto

import "challenge-admin/core/dto"

// UserLoginLogQueryReq app 用户登录日志查询
type UserLoginLogQueryReq struct {
	dto.Pagination `search:"-"`
	UserId         int64  `form:"userId" search:"type:exact;column:user_id;table:app_user_login_log" comment:"用户ID"`
	LoginAt        string `form:"loginAt" search:"type:exact;column:login_at;table:app_user_login_log" comment:"登录时间"`
	LoginIp        string `form:"loginIp" search:"type:exact;column:login_ip;table:app_user_login_log" comment:"登录IP"`
	DeviceFp       string `form:"deviceFp" search:"type:exact;column:device_fp;table:app_user_login_log" comment:"设备指纹"`
	UserAgent      string `form:"userAgent" search:"type:contains;column:user_agent;table:app_user_login_log" comment:"UA信息"`
	Status         int8   `form:"status" search:"type:exact;column:status;table:app_user_login_log" comment:"登录状态 1成功 2失败 3风控拦截"`
	FailReason     string `form:"failReason" search:"type:contains;column:fail_reason;table:app_user_login_log" comment:"失败原因/拦截原因"`
	BeginCreatedAt string `form:"beginCreatedAt" search:"type:gte;column:created_at;table:app_user_login_log" comment:"创建时间起"`
	EndCreatedAt   string `form:"endCreatedAt" search:"type:lte;column:created_at;table:app_user_login_log" comment:"创建时间止"`
	UserLoginLogOrder
}

type UserLoginLogOrder struct {
	IdOrder        int64  `form:"idOrder" search:"type:order;column:id;table:app_user_login_log"`
	UserIdOrder    int64  `form:"userIdOrder" search:"type:order;column:user_id;table:app_user_login_log"`
	LoginAtOrder   string `form:"loginAtOrder" search:"type:order;column:login_at;table:app_user_login_log"`
	StatusOrder    string `form:"statusOrder" search:"type:order;column:status;table:app_user_login_log"`
	CreatedAtOrder string `form:"createdAtOrder" search:"type:order;column:created_at;table:app_user_login_log"`
}

func (m *UserLoginLogQueryReq) GetNeedSearch() interface{} {
	return *m
}
