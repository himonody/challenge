package dto

import "challenge-admin/core/dto"

// RiskAppeal
type RiskAppealQueryReq struct {
	dto.Pagination `search:"-"`
	UserId         int64  `form:"userId" search:"type:exact;column:user_id;table:app_risk_appeal" comment:"用户ID"`
	Status         int8   `form:"status" search:"type:exact;column:status;table:app_risk_appeal" comment:"状态"`
	DeviceFp       string `form:"deviceFp" search:"type:exact;column:device_fp;table:app_risk_appeal" comment:"设备指纹"`
	BeginCreatedAt string `form:"beginCreatedAt" search:"type:gte;column:created_at;table:app_risk_appeal" comment:"申诉时间起"`
	EndCreatedAt   string `form:"endCreatedAt" search:"type:lte;column:created_at;table:app_risk_appeal" comment:"申诉时间止"`
	RiskAppealOrder
}

type RiskAppealOrder struct {
	IdOrder        int64  `form:"idOrder" search:"type:order;column:id;table:app_risk_appeal"`
	UserIdOrder    int64  `form:"userIdOrder" search:"type:order;column:user_id;table:app_risk_appeal"`
	StatusOrder    string `form:"statusOrder" search:"type:order;column:status;table:app_risk_appeal"`
	CreatedAtOrder string `form:"createdAtOrder" search:"type:order;column:created_at;table:app_risk_appeal"`
}

func (m *RiskAppealQueryReq) GetNeedSearch() interface{} { return *m }

// Appeal review
type RiskAppealReviewReq struct {
	Id           int64  `json:"-" uri:"id" comment:"申诉ID"`
	Status       int8   `json:"status" comment:"状态 1待处理 2通过 3拒绝"`
	ReviewRemark string `json:"reviewRemark" comment:"审核备注"`
	ActionResult int8   `json:"actionResult" comment:"处理结果"`
	ReviewerId   int64  `json:"-" comment:"审核人"`
}
