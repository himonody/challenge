package models

import "time"

type RiskBlacklist struct {
	Id        int64      `json:"id" gorm:"primaryKey;autoIncrement"`
	Type      string     `json:"type" gorm:"column:type;type:varchar(16);comment:ip/device/country/mobile/email"`
	Value     string     `json:"value" gorm:"column:value;type:varchar(128);comment:命中值"`
	RiskLevel int8       `json:"riskLevel" gorm:"column:risk_level;type:tinyint;comment:风险等级"`
	Reason    string     `json:"reason" gorm:"column:reason;type:varchar(255);comment:封禁原因"`
	Status    string     `json:"status" gorm:"column:status;type:char(1);comment:1生效 2失效"`
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
}

func (RiskBlacklist) TableName() string { return "app_risk_blacklist" }
