package models

import "time"

// RiskStrategy 风控策略
type RiskStrategy struct {
	ID            uint64    `json:"id" gorm:"primaryKey;autoIncrement;comment:策略ID"`
	Scene         string    `json:"scene" gorm:"column:scene;type:varchar(32);not null;uniqueIndex:uk_scene_rule,priority:1;index:idx_scene_status_priority,priority:1;index:idx_scene_identity,priority:1;comment:场景"`
	RuleCode      string    `json:"ruleCode" gorm:"column:rule_code;type:varchar(64);not null;uniqueIndex:uk_scene_rule,priority:2;comment:规则编码"`
	IdentityType  string    `json:"identityType" gorm:"column:identity_type;type:varchar(16);not null;index:idx_scene_identity,priority:2;comment:统计维度"`
	WindowSeconds int       `json:"windowSeconds" gorm:"column:window_seconds;type:int;not null;comment:统计窗口(秒)"`
	Threshold     int       `json:"threshold" gorm:"column:threshold;type:int;not null;comment:触发阈值（次数）"`
	Action        string    `json:"action" gorm:"column:action;type:varchar(32);not null;comment:触发动作编码"`
	ActionValue   int       `json:"actionValue" gorm:"column:action_value;type:int;not null;default:0;comment:动作值(秒/分数)"`
	Status        int8      `json:"status" gorm:"column:status;type:tinyint;not null;default:1;index:idx_scene_status_priority,priority:2;comment:1启用 0停用"`
	Priority      int       `json:"priority" gorm:"column:priority;type:int;not null;default:100;index:idx_scene_status_priority,priority:3;comment:优先级 数值越小越优先"`
	Remark        string    `json:"remark" gorm:"column:remark;type:varchar(255);comment:说明"`
	CreatedAt     time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`
}

// TableName 指定表名
func (RiskStrategy) TableName() string {
	return "app_risk_strategy"
}
