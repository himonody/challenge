package dto

import "challenge-admin/core/dto"

type WithdrawOrderQueryReq struct {
	dto.Pagination `search:"-"`
	UserId         int64  `form:"userId" search:"type:exact;column:user_id;table:app_withdraw_order" comment:"用户ID"`
	Status         int8   `form:"status" search:"type:exact;column:status;table:app_withdraw_order" comment:"状态 1待审核 2通过 3拒绝 4打款完成"`
	BeginCreatedAt string `form:"beginCreatedAt" search:"type:gte;column:created_at;table:app_withdraw_order" comment:"申请时间起"`
	EndCreatedAt   string `form:"endCreatedAt" search:"type:lte;column:created_at;table:app_withdraw_order" comment:"申请时间止"`
	WithdrawOrderOrder
}

type WithdrawOrderOrder struct {
	IdOrder        int64  `form:"idOrder" search:"type:order;column:id;table:app_withdraw_order"`
	UserIdOrder    int64  `form:"userIdOrder" search:"type:order;column:user_id;table:app_withdraw_order"`
	AmountOrder    string `form:"amountOrder" search:"type:order;column:amount;table:app_withdraw_order"`
	StatusOrder    string `form:"statusOrder" search:"type:order;column:status;table:app_withdraw_order"`
	CreatedAtOrder string `form:"createdAtOrder" search:"type:order;column:created_at;table:app_withdraw_order"`
}

func (m *WithdrawOrderQueryReq) GetNeedSearch() interface{} {
	return *m
}

type WithdrawOrderUpdateStatusReq struct {
	Id           int64  `uri:"id" json:"-" comment:"提现订单ID"`
	Status       int8   `json:"status" comment:"状态 1待审核 2通过 3拒绝 4打款完成"`
	RejectReason string `json:"rejectReason" comment:"拒绝原因"`
	ReviewIp     string `json:"reviewIp" comment:"审核IP"`
	ReviewerId   int64  `json:"-" comment:"审核人ID"`
}
