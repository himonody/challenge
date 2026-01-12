package dto

import "challenge-admin/core/dto"

// RiskEvent
type RiskEventQueryReq struct {
	dto.Pagination `search:"-"`
	UserId         int64  `form:"userId" search:"type:exact;column:user_id;table:app_risk_event" comment:"用户ID"`
	EventType      int8   `form:"eventType" search:"type:exact;column:event_type;table:app_risk_event" comment:"事件类型"`
	BeginCreatedAt string `form:"beginCreatedAt" search:"type:gte;column:created_at;table:app_risk_event" comment:"时间起"`
	EndCreatedAt   string `form:"endCreatedAt" search:"type:lte;column:created_at;table:app_risk_event" comment:"时间止"`
	RiskEventOrder
}

type RiskEventOrder struct {
	IdOrder        int64  `form:"idOrder" search:"type:order;column:id;table:app_risk_event"`
	UserIdOrder    int64  `form:"userIdOrder" search:"type:order;column:user_id;table:app_risk_event"`
	EventTypeOrder string `form:"eventTypeOrder" search:"type:order;column:event_type;table:app_risk_event"`
	CreatedAtOrder string `form:"createdAtOrder" search:"type:order;column:created_at;table:app_risk_event"`
}

func (m *RiskEventQueryReq) GetNeedSearch() interface{} { return *m }
