package dto

import "challenge-admin/core/dto"

// RiskUser
type RiskUserQueryReq struct {
	dto.Pagination `search:"-"`
	UserId         int64  `form:"userId" search:"type:exact;column:user_id;table:app_risk_user" comment:"用户ID"`
	RiskLevel      int8   `form:"riskLevel" search:"type:exact;column:risk_level;table:app_risk_user" comment:"风险等级"`
	BeginUpdatedAt string `form:"beginUpdatedAt" search:"type:gte;column:updated_at;table:app_risk_user" comment:"更新时间起"`
	EndUpdatedAt   string `form:"endUpdatedAt" search:"type:lte;column:updated_at;table:app_risk_user" comment:"更新时间止"`
	RiskUserOrder
}

type RiskUserOrder struct {
	UserIdOrder    int64  `form:"userIdOrder" search:"type:order;column:user_id;table:app_risk_user"`
	RiskLevelOrder string `form:"riskLevelOrder" search:"type:order;column:risk_level;table:app_risk_user"`
	UpdatedAtOrder string `form:"updatedAtOrder" search:"type:order;column:updated_at;table:app_risk_user"`
}

func (m *RiskUserQueryReq) GetNeedSearch() interface{} { return *m }
