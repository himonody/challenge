package dto

import "challenge-admin/core/dto"

// RiskActionQueryReq 风控动作查询
type RiskActionQueryReq struct {
	dto.Pagination `search:"-"`
	Code           string `form:"code" search:"type:exact;column:code;table:app_risk_action" comment:"动作编码"`
	Type           string `form:"type" search:"type:exact;column:type;table:app_risk_action" comment:"动作类型"`
	RiskActionOrder
}

type RiskActionOrder struct {
	CodeOrder    string `form:"codeOrder" search:"type:order;column:code;table:app_risk_action"`
	TypeOrder    string `form:"typeOrder" search:"type:order;column:type;table:app_risk_action"`
	DefaultOrder string `form:"defaultValueOrder" search:"type:order;column:default_value;table:app_risk_action"`
}

func (m *RiskActionQueryReq) GetNeedSearch() interface{} { return *m }

// Create / Update
type RiskActionCreateReq struct {
	Code         string `json:"code" comment:"动作编码"`
	Type         string `json:"type" comment:"动作类型"`
	DefaultValue int    `json:"defaultValue" comment:"默认动作值"`
	Remark       string `json:"remark" comment:"说明"`
	CurrUserId   int64  `json:"-" comment:"当前用户"`
}

type RiskActionUpdateReq struct {
	Code         string `json:"-" uri:"code" comment:"动作编码"`
	Type         string `json:"type" comment:"动作类型"`
	DefaultValue int    `json:"defaultValue" comment:"默认动作值"`
	Remark       string `json:"remark" comment:"说明"`
	CurrUserId   int64  `json:"-" comment:"当前用户"`
}
