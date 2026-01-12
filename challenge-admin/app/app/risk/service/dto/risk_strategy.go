package dto

import "challenge-admin/core/dto"

// RiskStrategyQueryReq 风控策略查询
type RiskStrategyQueryReq struct {
	dto.Pagination `search:"-"`
	Scene          string `form:"scene" search:"type:exact;column:scene;table:app_risk_strategy" comment:"场景"`
	RuleCode       string `form:"ruleCode" search:"type:exact;column:rule_code;table:app_risk_strategy" comment:"规则编码"`
	IdentityType   string `form:"identityType" search:"type:exact;column:identity_type;table:app_risk_strategy" comment:"统计维度"`
	Status         int8   `form:"status" search:"type:exact;column:status;table:app_risk_strategy" comment:"状态"`
	RiskStrategyOrder
}

type RiskStrategyOrder struct {
	IdOrder       int64  `form:"idOrder" search:"type:order;column:id;table:app_risk_strategy"`
	PriorityOrder string `form:"priorityOrder" search:"type:order;column:priority;table:app_risk_strategy"`
	UpdatedOrder  string `form:"updatedAtOrder" search:"type:order;column:updated_at;table:app_risk_strategy"`
}

func (m *RiskStrategyQueryReq) GetNeedSearch() interface{} { return *m }

// Create / Update
type RiskStrategyCreateReq struct {
	Scene         string `json:"scene" comment:"场景"`
	RuleCode      string `json:"ruleCode" comment:"规则编码"`
	IdentityType  string `json:"identityType" comment:"统计维度"`
	WindowSeconds int    `json:"windowSeconds" comment:"窗口秒"`
	Threshold     int    `json:"threshold" comment:"阈值次数"`
	Action        string `json:"action" comment:"动作编码"`
	ActionValue   int    `json:"actionValue" comment:"动作值"`
	Status        int8   `json:"status" comment:"状态"`
	Priority      int    `json:"priority" comment:"优先级"`
	Remark        string `json:"remark" comment:"说明"`
	CurrUserId    int64  `json:"-" comment:"当前用户"`
}

type RiskStrategyUpdateReq struct {
	Id            uint64 `json:"id" uri:"id" comment:"策略ID"`
	Scene         string `json:"scene" comment:"场景"`
	RuleCode      string `json:"ruleCode" comment:"规则编码"`
	IdentityType  string `json:"identityType" comment:"统计维度"`
	WindowSeconds int    `json:"windowSeconds" comment:"窗口秒"`
	Threshold     int    `json:"threshold" comment:"阈值次数"`
	Action        string `json:"action" comment:"动作编码"`
	ActionValue   int    `json:"actionValue" comment:"动作值"`
	Status        int8   `json:"status" comment:"状态"`
	Priority      int    `json:"priority" comment:"优先级"`
	Remark        string `json:"remark" comment:"说明"`
	CurrUserId    int64  `json:"-" comment:"当前用户"`
}
