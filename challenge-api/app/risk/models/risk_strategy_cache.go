package models

// RiskStrategyCache 风控策略缓存
type RiskStrategyCache struct {
	Scene         string `json:"scene" gorm:"column:scene;type:varchar(32);primaryKey;not null;index:idx_scene_identity,priority:1;comment:场景"`
	IdentityType  string `json:"identityType" gorm:"column:identity_type;type:varchar(16);primaryKey;not null;index:idx_scene_identity,priority:2;comment:统计维度"`
	RuleCode      string `json:"ruleCode" gorm:"column:rule_code;type:varchar(64);primaryKey;not null;comment:规则编码"`
	WindowSeconds int    `json:"windowSeconds" gorm:"column:window_seconds;type:int;not null;comment:统计窗口(秒)"`
	Threshold     int    `json:"threshold" gorm:"column:threshold;type:int;not null;comment:触发阈值（次数）"`
	Action        string `json:"action" gorm:"column:action;type:varchar(32);not null;comment:触发动作编码"`
	ActionValue   int    `json:"actionValue" gorm:"column:action_value;type:int;not null;comment:动作值(秒/分数)"`
}

// TableName 指定表名
func (RiskStrategyCache) TableName() string {
	return "app_risk_strategy_cache"
}
