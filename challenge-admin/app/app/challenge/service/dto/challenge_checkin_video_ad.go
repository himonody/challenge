package dto

import "challenge-admin/core/dto"

type ChallengeCheckinVideoAdQueryReq struct {
	dto.Pagination `search:"-"`
	CheckinId      uint64 `form:"checkinId" search:"type:exact;column:checkin_id;table:app_challenge_checkin_video_ad"`
	UserId         uint64 `form:"userId" search:"type:exact;column:user_id;table:app_challenge_checkin_video_ad"`
	VerifyStatus   int8   `form:"verifyStatus" search:"type:exact;column:verify_status;table:app_challenge_checkin_video_ad"`
	VideoAdOrder
}

type VideoAdOrder struct {
	IdOrder         int64  `form:"idOrder" search:"type:order;column:id;table:app_challenge_checkin_video_ad"`
	CreatedAtOrder  string `form:"createdAtOrder" search:"type:order;column:created_at;table:app_challenge_checkin_video_ad"`
	VerifiedAtOrder string `form:"verifiedAtOrder" search:"type:order;column:verified_at;table:app_challenge_checkin_video_ad"`
}

func (m *ChallengeCheckinVideoAdQueryReq) GetNeedSearch() interface{} { return *m }

type ChallengeCheckinVideoAdCreateReq struct {
	CheckinId     uint64 `json:"checkinId"`
	UserId        uint64 `json:"userId"`
	AdPlatform    string `json:"adPlatform"`
	AdUnitId      string `json:"adUnitId"`
	AdOrderNo     string `json:"adOrderNo"`
	VideoDuration int    `json:"videoDuration"`
	WatchDuration int    `json:"watchDuration"`
	RewardAmount  string `json:"rewardAmount"`
	VerifyStatus  int8   `json:"verifyStatus"`
	CurrUserId    int64  `json:"-"`
}

type ChallengeCheckinVideoAdUpdateReq struct {
	Id            uint64 `json:"-" uri:"id"`
	AdPlatform    string `json:"adPlatform"`
	AdUnitId      string `json:"adUnitId"`
	AdOrderNo     string `json:"adOrderNo"`
	VideoDuration int    `json:"videoDuration"`
	WatchDuration int    `json:"watchDuration"`
	RewardAmount  string `json:"rewardAmount"`
	VerifyStatus  int8   `json:"verifyStatus"`
	CurrUserId    int64  `json:"-"`
}
