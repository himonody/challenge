package dto

import "challenge-admin/core/dto"

type ChallengeCheckinImageQueryReq struct {
	dto.Pagination `search:"-"`
	CheckinId      uint64 `form:"checkinId" search:"type:exact;column:checkin_id;table:app_challenge_checkin_image" comment:"打卡ID"`
	UserId         uint64 `form:"userId" search:"type:exact;column:user_id;table:app_challenge_checkin_image" comment:"用户ID"`
	Status         int8   `form:"status" search:"type:exact;column:status;table:app_challenge_checkin_image" comment:"状态"`
	CheckinImageOrder
}

type CheckinImageOrder struct {
	IdOrder        int64  `form:"idOrder" search:"type:order;column:id;table:app_challenge_checkin_image"`
	SortNoOrder    string `form:"sortNoOrder" search:"type:order;column:sort_no;table:app_challenge_checkin_image"`
	CreatedAtOrder string `form:"createdAtOrder" search:"type:order;column:created_at;table:app_challenge_checkin_image"`
}

func (m *ChallengeCheckinImageQueryReq) GetNeedSearch() interface{} { return *m }

type ChallengeCheckinImageCreateReq struct {
	CheckinId  uint64 `json:"checkinId"`
	UserId     uint64 `json:"userId"`
	ImageUrl   string `json:"imageUrl"`
	ImageHash  string `json:"imageHash"`
	SortNo     int8   `json:"sortNo"`
	Status     int8   `json:"status"`
	CurrUserId int64  `json:"-"`
}

type ChallengeCheckinImageUpdateReq struct {
	Id         uint64 `json:"-" uri:"id"`
	ImageUrl   string `json:"imageUrl"`
	ImageHash  string `json:"imageHash"`
	SortNo     int8   `json:"sortNo"`
	Status     int8   `json:"status"`
	CurrUserId int64  `json:"-"`
}
