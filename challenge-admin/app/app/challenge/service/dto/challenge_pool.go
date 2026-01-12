package dto

import "challenge-admin/core/dto"

type ChallengePoolQueryReq struct {
	dto.Pagination `search:"-"`
	ConfigId       uint64 `form:"configId" search:"type:exact;column:config_id;table:app_challenge_pool"`
	Settled        int8   `form:"settled" search:"type:exact;column:settled;table:app_challenge_pool"`
	ChallengePoolOrder
}

type ChallengePoolOrder struct {
	IdOrder        int64  `form:"idOrder" search:"type:order;column:id;table:app_challenge_pool"`
	StartDateOrder string `form:"startDateOrder" search:"type:order;column:start_date;table:app_challenge_pool"`
	CreatedAtOrder string `form:"createdAtOrder" search:"type:order;column:created_at;table:app_challenge_pool"`
}

func (m *ChallengePoolQueryReq) GetNeedSearch() interface{} { return *m }

type ChallengePoolCreateReq struct {
	ConfigId    uint64 `json:"configId"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	TotalAmount string `json:"totalAmount"`
	Settled     int8   `json:"settled"`
	CurrUserId  int64  `json:"-"`
}

type ChallengePoolUpdateReq struct {
	Id          uint64 `json:"-" uri:"id"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	TotalAmount string `json:"totalAmount"`
	Settled     int8   `json:"settled"`
	CurrUserId  int64  `json:"-"`
}
