package models

// RiskAction 风控动作
type RiskAction struct {
	Code         string `json:"code" gorm:"column:code;primaryKey;type:varchar(32);not null;comment:动作编码"`
	Type         string `json:"type" gorm:"column:type;type:varchar(16);not null;comment:动作类型"`
	DefaultValue int    `json:"defaultValue" gorm:"column:default_value;type:int;not null;default:0;comment:默认动作值"`
	Remark       string `json:"remark" gorm:"column:remark;type:varchar(255);comment:说明"`
}

// TableName 指定表名
func (RiskAction) TableName() string {
	return "app_risk_action"
}
