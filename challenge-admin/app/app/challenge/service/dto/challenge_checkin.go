package dto

import "challenge-admin/core/dto"

type ChallengeCheckinQueryReq struct {
	dto.Pagination `search:"-"`
	UserId         uint64 `form:"userId" search:"type:exact;column:user_id;table:app_challenge_checkin" comment:"用户ID"`
	ChallengeId    uint64 `form:"challengeId" search:"type:exact;column:challenge_id;table:app_challenge_checkin" comment:"挑战ID"`
	Status         int8   `form:"status" search:"type:exact;column:status;table:app_challenge_checkin" comment:"状态"`
	CheckinOrder
}

type CheckinOrder struct {
	IdOrder          int64  `form:"idOrder" search:"type:order;column:id;table:app_challenge_checkin"`
	CheckinDateOrder string `form:"checkinDateOrder" search:"type:order;column:checkin_date;table:app_challenge_checkin"`
	CreatedAtOrder   string `form:"createdAtOrder" search:"type:order;column:created_at;table:app_challenge_checkin"`
}

func (m *ChallengeCheckinQueryReq) GetNeedSearch() interface{} { return *m }

type ChallengeCheckinCreateReq struct {
	ChallengeId uint64 `json:"challengeId"`
	UserId      uint64 `json:"userId"`
	CheckinDate string `json:"checkinDate"`
	CheckinTime string `json:"checkinTime"`
	MoodCode    int8   `json:"moodCode"`
	MoodText    string `json:"moodText"`
	ContentType int8   `json:"contentType"`
	Status      int8   `json:"status"`
	CurrUserId  int64  `json:"-"`
}

type ChallengeCheckinUpdateReq struct {
	Id          uint64 `json:"-" uri:"id"`
	CheckinDate string `json:"checkinDate"`
	CheckinTime string `json:"checkinTime"`
	MoodCode    int8   `json:"moodCode"`
	MoodText    string `json:"moodText"`
	ContentType int8   `json:"contentType"`
	Status      int8   `json:"status"`
	CurrUserId  int64  `json:"-"`
}
