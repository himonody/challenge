package dto

import "challenge-admin/core/dto"

type ChallengeConfigQueryReq struct {
	dto.Pagination `search:"-"`
	Status         int8 `form:"status" search:"type:exact;column:status;table:app_challenge_config" comment:"状态"`
	DayCount       int  `form:"dayCount" search:"type:exact;column:day_count;table:app_challenge_config" comment:"挑战天数"`
	ConfigOrder
}

type ConfigOrder struct {
	IdOrder        int64  `form:"idOrder" search:"type:order;column:id;table:app_challenge_config"`
	SortOrder      string `form:"sortOrder" search:"type:order;column:sort;table:app_challenge_config"`
	CreatedAtOrder string `form:"createdAtOrder" search:"type:order;column:created_at;table:app_challenge_config"`
}

func (m *ChallengeConfigQueryReq) GetNeedSearch() interface{} { return *m }

type ChallengeConfigCreateReq struct {
	DayCount      int    `json:"dayCount"`
	Amount        string `json:"amount"`
	CheckinStart  string `json:"checkinStart"`
	CheckinEnd    string `json:"checkinEnd"`
	PlatformBonus string `json:"platformBonus"`
	Status        int8   `json:"status"`
	Sort          int8   `json:"sort"`
	CurrUserId    int64  `json:"-"`
}

type ChallengeConfigUpdateReq struct {
	Id            uint64 `json:"-" uri:"id"`
	DayCount      int    `json:"dayCount"`
	Amount        string `json:"amount"`
	CheckinStart  string `json:"checkinStart"`
	CheckinEnd    string `json:"checkinEnd"`
	PlatformBonus string `json:"platformBonus"`
	Status        int8   `json:"status"`
	Sort          int8   `json:"sort"`
	CurrUserId    int64  `json:"-"`
}
