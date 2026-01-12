package models

// RiskStrategyCache 风控策略缓存
type RiskStrategyCache struct {
	Scene         string `json:"scene" gorm:"column:scene;primaryKey;type:varchar(32);comment:场景"`
	IdentityType  string `json:"identityType" gorm:"column:identity_type;primaryKey;type:varchar(16);comment:统计维度"`
	RuleCode      string `json:"ruleCode" gorm:"column:rule_code;primaryKey;type:varchar(64);comment:规则编码"`
	WindowSeconds int    `json:"windowSeconds" gorm:"column:window_seconds;type:int;comment:统计窗口(秒)"`
	Threshold     int    `json:"threshold" gorm:"column:threshold;type:int;comment:触发阈值（次数）"`
	Action        string `json:"action" gorm:"column:action;type:varchar(32);comment:触发动作编码"`
	ActionValue   int    `json:"actionValue" gorm:"column:action_value;type:int;comment:动作值(秒/分数)"`
}

func (RiskStrategyCache) TableName() string { return "app_risk_strategy_cache" }
