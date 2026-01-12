package dto

import "challenge-admin/core/dto"

type ChallengeDailyStatQueryReq struct {
	dto.Pagination `search:"-"`
	StatDate       string `form:"statDate" search:"type:exact;column:stat_date;table:app_challenge_daily_stat"`
	ChallengeDailyStatOrder
}

type ChallengeDailyStatOrder struct {
	StatDateOrder  string `form:"statDateOrder" search:"type:order;column:stat_date;table:app_challenge_daily_stat"`
	CreatedAtOrder string `form:"createdAtOrder" search:"type:order;column:created_at;table:app_challenge_daily_stat"`
}

func (m *ChallengeDailyStatQueryReq) GetNeedSearch() interface{} { return *m }

type ChallengeDailyStatCreateReq struct {
	StatDate       string `json:"statDate"`
	JoinUserCnt    int    `json:"joinUserCnt"`
	SuccessUserCnt int    `json:"successUserCnt"`
	FailUserCnt    int    `json:"failUserCnt"`
	JoinAmount     string `json:"joinAmount"`
	SuccessAmount  string `json:"successAmount"`
	FailAmount     string `json:"failAmount"`
	PlatformBonus  string `json:"platformBonus"`
	PoolAmount     string `json:"poolAmount"`
	CurrUserId     int64  `json:"-"`
}

type ChallengeDailyStatUpdateReq struct {
	StatDate       string `json:"statDate" uri:"statDate"`
	JoinUserCnt    int    `json:"joinUserCnt"`
	SuccessUserCnt int    `json:"successUserCnt"`
	FailUserCnt    int    `json:"failUserCnt"`
	JoinAmount     string `json:"joinAmount"`
	SuccessAmount  string `json:"successAmount"`
	FailAmount     string `json:"failAmount"`
	PlatformBonus  string `json:"platformBonus"`
	PoolAmount     string `json:"poolAmount"`
	CurrUserId     int64  `json:"-"`
}
