package models

import "time"

// RiskAppeal 风控申诉
type RiskAppeal struct {
	ID             uint64     `json:"id" gorm:"primaryKey;autoIncrement;comment:申诉ID"`
	UserID         uint64     `json:"userId" gorm:"column:user_id;type:bigint;not null;index:idx_user_id;comment:申诉用户ID"`
	RiskLevel      int8       `json:"riskLevel" gorm:"column:risk_level;type:tinyint;not null;default:0;comment:申诉时风险等级"`
	RiskReason     string     `json:"riskReason" gorm:"column:risk_reason;type:varchar(255);not null;default:'';comment:触发风控原因"`
	AppealType     int8       `json:"appealType" gorm:"column:appeal_type;type:tinyint;not null;default:1;comment:申诉类型 1账号封禁 2登录限制 3设备封禁"`
	AppealReason   string     `json:"appealReason" gorm:"column:appeal_reason;type:varchar(500);not null;comment:用户申诉说明"`
	AppealEvidence string     `json:"appealEvidence" gorm:"column:appeal_evidence;type:varchar(1000);comment:申诉凭证"`
	IP             string     `json:"ip" gorm:"column:ip;type:varchar(45);comment:申诉时IP"`
	DeviceFP       string     `json:"deviceFp" gorm:"column:device_fp;type:varchar(64);index:idx_device_fp;comment:申诉设备指纹"`
	Status         int8       `json:"status" gorm:"column:status;type:tinyint;not null;default:1;index:idx_status;comment:状态 1待处理 2通过 3拒绝"`
	ReviewerID     *uint64    `json:"reviewerId" gorm:"column:reviewer_id;type:bigint;comment:审核人ID"`
	ReviewRemark   string     `json:"reviewRemark" gorm:"column:review_remark;type:varchar(255);comment:审核备注"`
	ActionResult   int8       `json:"actionResult" gorm:"column:action_result;type:tinyint;not null;default:0;comment:处理结果 0无操作 1已解封账号 2已解封设备"`
	CreatedAt      time.Time  `json:"createdAt" gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:申诉时间"`
	ReviewedAt     *time.Time `json:"reviewedAt" gorm:"column:reviewed_at;type:datetime;comment:审核时间"`
}

// TableName 指定表名
func (RiskAppeal) TableName() string {
	return "app_risk_appeal"
}
