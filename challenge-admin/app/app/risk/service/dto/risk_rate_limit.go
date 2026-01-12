package dto

import "challenge-admin/core/dto"

// RiskRateLimit
type RiskRateLimitQueryReq struct {
	dto.Pagination `search:"-"`
	Scene          string `form:"scene" search:"type:exact;column:scene;table:app_risk_rate_limit" comment:"场景"`
	IdentityType   string `form:"identityType" search:"type:exact;column:identity_type;table:app_risk_rate_limit" comment:"标识类型"`
	IdentityValue  string `form:"identityValue" search:"type:exact;column:identity_value;table:app_risk_rate_limit" comment:"标识值"`
	Blocked        string `form:"blocked" search:"type:exact;column:blocked;table:app_risk_rate_limit" comment:"是否拦截"`
	RiskRateLimitOrder
}

type RiskRateLimitOrder struct {
	IdOrder          int64  `form:"idOrder" search:"type:order;column:id;table:app_risk_rate_limit"`
	WindowStartOrder string `form:"windowStartOrder" search:"type:order;column:window_start;table:app_risk_rate_limit"`
	WindowEndOrder   string `form:"windowEndOrder" search:"type:order;column:window_end;table:app_risk_rate_limit"`
}

func (m *RiskRateLimitQueryReq) GetNeedSearch() interface{} { return *m }
