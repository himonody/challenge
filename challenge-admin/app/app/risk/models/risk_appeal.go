package models

import "time"

type RiskAppeal struct {
	Id             int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:申诉ID"`
	UserId         int64      `json:"userId" gorm:"column:user_id;type:bigint;comment:申诉用户ID"`
	RiskLevel      int8       `json:"riskLevel" gorm:"column:risk_level;type:tinyint;comment:申诉时风险等级"`
	RiskReason     string     `json:"riskReason" gorm:"column:risk_reason;type:varchar(255);comment:触发风控原因"`
	AppealType     int8       `json:"appealType" gorm:"column:appeal_type;type:tinyint;comment:申诉类型"`
	AppealReason   string     `json:"appealReason" gorm:"column:appeal_reason;type:varchar(500);comment:用户申诉说明"`
	AppealEvidence string     `json:"appealEvidence" gorm:"column:appeal_evidence;type:varchar(1000);comment:申诉凭证"`
	Ip             string     `json:"ip" gorm:"column:ip;type:varchar(45);comment:申诉时IP"`
	DeviceFp       string     `json:"deviceFp" gorm:"column:device_fp;type:varchar(64);comment:申诉设备指纹"`
	Status         int8       `json:"status" gorm:"column:status;type:tinyint;comment:状态 1待处理 2通过 3拒绝"`
	ReviewerId     *int64     `json:"reviewerId" gorm:"column:reviewer_id;type:bigint;comment:审核人ID"`
	ReviewRemark   string     `json:"reviewRemark" gorm:"column:review_remark;type:varchar(255);comment:审核备注"`
	ActionResult   int8       `json:"actionResult" gorm:"column:action_result;type:tinyint;comment:处理结果"`
	CreatedAt      *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:申诉时间"`
	ReviewedAt     *time.Time `json:"reviewedAt" gorm:"column:reviewed_at;type:datetime;comment:审核时间"`
}

func (RiskAppeal) TableName() string { return "app_risk_appeal" }
