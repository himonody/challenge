package dto

import "challenge-admin/core/dto"

type ChallengeTotalStatQueryReq struct {
	dto.Pagination `search:"-"`
	ChallengeTotalStatOrder
}

type ChallengeTotalStatOrder struct {
	IdOrder        int64  `form:"idOrder" search:"type:order;column:id;table:app_challenge_total_stat"`
	UpdatedAtOrder string `form:"updatedAtOrder" search:"type:order;column:updated_at;table:app_challenge_total_stat"`
}

func (m *ChallengeTotalStatQueryReq) GetNeedSearch() interface{} { return *m }

type ChallengeTotalStatCreateReq struct {
	Id                 int    `json:"id"`
	TotalUserCnt       int    `json:"totalUserCnt"`
	TotalJoinCnt       int    `json:"totalJoinCnt"`
	TotalSuccessCnt    int    `json:"totalSuccessCnt"`
	TotalFailCnt       int    `json:"totalFailCnt"`
	TotalJoinAmount    string `json:"totalJoinAmount"`
	TotalSuccessAmount string `json:"totalSuccessAmount"`
	TotalFailAmount    string `json:"totalFailAmount"`
	TotalPlatformBonus string `json:"totalPlatformBonus"`
	TotalPoolAmount    string `json:"totalPoolAmount"`
	CurrUserId         int64  `json:"-"`
}

type ChallengeTotalStatUpdateReq struct {
	Id                 int    `json:"id" uri:"id"`
	TotalUserCnt       int    `json:"totalUserCnt"`
	TotalJoinCnt       int    `json:"totalJoinCnt"`
	TotalSuccessCnt    int    `json:"totalSuccessCnt"`
	TotalFailCnt       int    `json:"totalFailCnt"`
	TotalJoinAmount    string `json:"totalJoinAmount"`
	TotalSuccessAmount string `json:"totalSuccessAmount"`
	TotalFailAmount    string `json:"totalFailAmount"`
	TotalPlatformBonus string `json:"totalPlatformBonus"`
	TotalPoolAmount    string `json:"totalPoolAmount"`
	CurrUserId         int64  `json:"-"`
}
