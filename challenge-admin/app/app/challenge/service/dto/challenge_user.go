package dto

import "challenge-admin/core/dto"

type ChallengeUserQueryReq struct {
	dto.Pagination `search:"-"`
	UserId         uint64 `form:"userId" search:"type:exact;column:user_id;table:app_challenge_user"`
	Status         int8   `form:"status" search:"type:exact;column:status;table:app_challenge_user"`
	PoolId         uint64 `form:"poolId" search:"type:exact;column:pool_id;table:app_challenge_user"`
	ChallengeUserOrder
}

type ChallengeUserOrder struct {
	IdOrder        int64  `form:"idOrder" search:"type:order;column:id;table:app_challenge_user"`
	StartDateOrder string `form:"startDateOrder" search:"type:order;column:start_date;table:app_challenge_user"`
	CreatedAtOrder string `form:"createdAtOrder" search:"type:order;column:created_at;table:app_challenge_user"`
}

func (m *ChallengeUserQueryReq) GetNeedSearch() interface{} { return *m }

type ChallengeUserCreateReq struct {
	UserId          uint64 `json:"userId"`
	ConfigId        uint64 `json:"configId"`
	PoolId          uint64 `json:"poolId"`
	ChallengeAmount string `json:"challengeAmount"`
	StartDate       int    `json:"startDate"`
	EndDate         int    `json:"endDate"`
	Status          int8   `json:"status"`
	FailReason      int8   `json:"failReason"`
	CurrUserId      int64  `json:"-"`
}

type ChallengeUserUpdateReq struct {
	Id              uint64 `json:"-" uri:"id"`
	PoolId          uint64 `json:"poolId"`
	ChallengeAmount string `json:"challengeAmount"`
	StartDate       int    `json:"startDate"`
	EndDate         int    `json:"endDate"`
	Status          int8   `json:"status"`
	FailReason      int8   `json:"failReason"`
	CurrUserId      int64  `json:"-"`
}
