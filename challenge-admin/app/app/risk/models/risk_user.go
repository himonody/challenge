package models

import "time"

type RiskUser struct {
	UserId    int64      `json:"userId" gorm:"column:user_id;primaryKey;comment:用户ID"`
	RiskLevel int8       `json:"riskLevel" gorm:"column:risk_level;type:tinyint;comment:0正常 1观察 2限制 3封禁"`
	RiskScore int        `json:"riskScore" gorm:"column:risk_score;type:int;comment:风险评分"`
	Reason    string     `json:"reason" gorm:"column:reason;type:varchar(255);comment:风险原因"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updated_at;type:datetime;comment:更新时间"`
}

func (RiskUser) TableName() string { return "app_risk_user" }
