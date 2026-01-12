package dto

import "challenge-admin/core/dto"

type ChallengePoolFlowQueryReq struct {
	dto.Pagination `search:"-"`
	PoolId         uint64 `form:"poolId" search:"type:exact;column:pool_id;table:app_challenge_pool_flow"`
	UserId         uint64 `form:"userId" search:"type:exact;column:user_id;table:app_challenge_pool_flow"`
	Type           int8   `form:"type" search:"type:exact;column:type;table:app_challenge_pool_flow"`
	ChallengePoolFlowOrder
}

type ChallengePoolFlowOrder struct {
	IdOrder        int64  `form:"idOrder" search:"type:order;column:id;table:app_challenge_pool_flow"`
	CreatedAtOrder string `form:"createdAtOrder" search:"type:order;column:created_at;table:app_challenge_pool_flow"`
}

func (m *ChallengePoolFlowQueryReq) GetNeedSearch() interface{} { return *m }

type ChallengePoolFlowCreateReq struct {
	PoolId     uint64 `json:"poolId"`
	UserId     uint64 `json:"userId"`
	Amount     string `json:"amount"`
	Type       int8   `json:"type"`
	CurrUserId int64  `json:"-"`
}

type ChallengePoolFlowUpdateReq struct {
	Id         uint64 `json:"-" uri:"id"`
	Amount     string `json:"amount"`
	Type       int8   `json:"type"`
	CurrUserId int64  `json:"-"`
}
