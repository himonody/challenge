package models

import "time"

// RiskBlacklist 风控黑名单
type RiskBlacklist struct {
	ID        uint64     `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	Type      string     `json:"type" gorm:"column:type;type:varchar(16);not null;uniqueIndex:uk_type_value,priority:1;index:idx_type_status,priority:1;comment:ip/device/country/mobile/email"`
	Value     string     `json:"value" gorm:"column:value;type:varchar(128);not null;uniqueIndex:uk_type_value,priority:2;comment:命中值"`
	RiskLevel int8       `json:"riskLevel" gorm:"column:risk_level;type:tinyint;not null;default:3;comment:风险等级"`
	Reason    string     `json:"reason" gorm:"column:reason;type:varchar(255);not null;default:'';comment:封禁原因"`
	Status    string     `json:"status" gorm:"column:status;type:char(1);not null;default:'1';index:idx_status;index:idx_type_status,priority:2;comment:1生效 2失效"`
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
}

// TableName 指定表名
func (RiskBlacklist) TableName() string {
	return "app_risk_blacklist"
}
