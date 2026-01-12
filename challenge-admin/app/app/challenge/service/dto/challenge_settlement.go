package dto

import "challenge-admin/core/dto"

type ChallengeSettlementQueryReq struct {
	dto.Pagination `search:"-"`
	ChallengeId    uint64 `form:"challengeId" search:"type:exact;column:challenge_id;table:app_challenge_settlement"`
	UserId         uint64 `form:"userId" search:"type:exact;column:user_id;table:app_challenge_settlement"`
	ChallengeSettlementOrder
}

type ChallengeSettlementOrder struct {
	IdOrder        int64  `form:"idOrder" search:"type:order;column:id;table:app_challenge_settlement"`
	CreatedAtOrder string `form:"createdAtOrder" search:"type:order;column:created_at;table:app_challenge_settlement"`
}

func (m *ChallengeSettlementQueryReq) GetNeedSearch() interface{} { return *m }

type ChallengeSettlementCreateReq struct {
	ChallengeId uint64 `json:"challengeId"`
	UserId      uint64 `json:"userId"`
	Reward      string `json:"reward"`
	CurrUserId  int64  `json:"-"`
}

type ChallengeSettlementUpdateReq struct {
	Id         uint64 `json:"-" uri:"id"`
	Reward     string `json:"reward"`
	CurrUserId int64  `json:"-"`
}
