package dto

import "challenge-admin/core/dto"

// RiskDevice
type RiskDeviceQueryReq struct {
	dto.Pagination `search:"-"`
	DeviceFp       string `form:"deviceFp" search:"type:exact;column:device_fp;table:app_risk_device" comment:"设备指纹"`
	UserId         int64  `form:"userId" search:"type:exact;column:user_id;table:app_risk_device" comment:"用户ID"`
	BeginCreatedAt string `form:"beginCreatedAt" search:"type:gte;column:created_at;table:app_risk_device" comment:"记录时间起"`
	EndCreatedAt   string `form:"endCreatedAt" search:"type:lte;column:created_at;table:app_risk_device" comment:"记录时间止"`
	RiskDeviceOrder
}

type RiskDeviceOrder struct {
	IdOrder        int64  `form:"idOrder" search:"type:order;column:id;table:app_risk_device"`
	UserIdOrder    int64  `form:"userIdOrder" search:"type:order;column:user_id;table:app_risk_device"`
	CreatedAtOrder string `form:"createdAtOrder" search:"type:order;column:created_at;table:app_risk_device"`
}

func (m *RiskDeviceQueryReq) GetNeedSearch() interface{} { return *m }
