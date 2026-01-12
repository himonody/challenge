package dto

import "challenge-admin/core/dto"

// RiskBlacklist
type RiskBlacklistQueryReq struct {
	dto.Pagination `search:"-"`
	Type           string `form:"type" search:"type:exact;column:type;table:app_risk_blacklist" comment:"类型"`
	Value          string `form:"value" search:"type:exact;column:value;table:app_risk_blacklist" comment:"命中值"`
	Status         string `form:"status" search:"type:exact;column:status;table:app_risk_blacklist" comment:"状态"`
	RiskLevel      int8   `form:"riskLevel" search:"type:exact;column:risk_level;table:app_risk_blacklist" comment:"风险等级"`
	BlacklistOrder
}

type BlacklistOrder struct {
	IdOrder        int64  `form:"idOrder" search:"type:order;column:id;table:app_risk_blacklist"`
	RiskLevelOrder string `form:"riskLevelOrder" search:"type:order;column:risk_level;table:app_risk_blacklist"`
	CreatedAtOrder string `form:"createdAtOrder" search:"type:order;column:created_at;table:app_risk_blacklist"`
}

func (m *RiskBlacklistQueryReq) GetNeedSearch() interface{} { return *m }

// Blacklist create/update
type RiskBlacklistCreateReq struct {
	Type       string `json:"type" comment:"类型"`
	Value      string `json:"value" comment:"命中值"`
	RiskLevel  int8   `json:"riskLevel" comment:"风险等级"`
	Reason     string `json:"reason" comment:"封禁原因"`
	Status     string `json:"status" comment:"状态"`
	CurrUserId int64  `json:"-" comment:"当前用户"`
}

type RiskBlacklistUpdateReq struct {
	Id         int64  `json:"-" uri:"id" comment:"ID"`
	Type       string `json:"type" comment:"类型"`
	Value      string `json:"value" comment:"命中值"`
	RiskLevel  int8   `json:"riskLevel" comment:"风险等级"`
	Reason     string `json:"reason" comment:"封禁原因"`
	Status     string `json:"status" comment:"状态"`
	CurrUserId int64  `json:"-" comment:"当前用户"`
}
