package dto

import "challenge-admin/core/dto"

// RiskStrategyCacheQueryReq 策略缓存查询
type RiskStrategyCacheQueryReq struct {
	dto.Pagination `search:"-"`
	Scene          string `form:"scene" search:"type:exact;column:scene;table:app_risk_strategy_cache" comment:"场景"`
	IdentityType   string `form:"identityType" search:"type:exact;column:identity_type;table:app_risk_strategy_cache" comment:"统计维度"`
	RuleCode       string `form:"ruleCode" search:"type:exact;column:rule_code;table:app_risk_strategy_cache" comment:"规则编码"`
	RiskStrategyCacheOrder
}

type RiskStrategyCacheOrder struct {
	SceneOrder        string `form:"sceneOrder" search:"type:order;column:scene;table:app_risk_strategy_cache"`
	IdentityTypeOrder string `form:"identityTypeOrder" search:"type:order;column:identity_type;table:app_risk_strategy_cache"`
	RuleCodeOrder     string `form:"ruleCodeOrder" search:"type:order;column:rule_code;table:app_risk_strategy_cache"`
}

func (m *RiskStrategyCacheQueryReq) GetNeedSearch() interface{} { return *m }
