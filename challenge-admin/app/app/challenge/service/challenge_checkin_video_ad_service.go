package service

import (
	"challenge-admin/app/app/challenge/models"
	"challenge-admin/app/app/challenge/service/dto"
	baseLang "challenge-admin/config/base/lang"
	cDto "challenge-admin/core/dto"
	"challenge-admin/core/dto/service"
	"challenge-admin/core/lang"
	"challenge-admin/core/middleware"
)

type ChallengeCheckinVideoAd struct {
	service.Service
}

func (e *ChallengeCheckinVideoAd) GetPage(c *dto.ChallengeCheckinVideoAdQueryReq, p *middleware.DataPermission) ([]models.ChallengeCheckinVideoAd, int64, int, error) {
	var list []models.ChallengeCheckinVideoAd
	var data models.ChallengeCheckinVideoAd
	var count int64
	err := e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			middleware.Permission(data.TableName(), p),
		).Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	return list, count, baseLang.SuccessCode, nil
}

func (e *ChallengeCheckinVideoAd) Create(c *dto.ChallengeCheckinVideoAdCreateReq) (uint64, int, error) {
	if c.CurrUserId <= 0 || c.CheckinId == 0 {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := models.ChallengeCheckinVideoAd{
		CheckinId:     c.CheckinId,
		UserId:        c.UserId,
		AdPlatform:    c.AdPlatform,
		AdUnitId:      c.AdUnitId,
		AdOrderNo:     c.AdOrderNo,
		VideoDuration: c.VideoDuration,
		WatchDuration: c.WatchDuration,
		VerifyStatus:  c.VerifyStatus,
	}
	if c.RewardAmount != "" {
		if v, err := models.ParseDecimal(c.RewardAmount); err == nil {
			data.RewardAmount = v
		}
	}
	if err := e.Orm.Create(&data).Error; err != nil {
		return 0, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}
	return data.Id, baseLang.SuccessCode, nil
}

func (e *ChallengeCheckinVideoAd) Update(c *dto.ChallengeCheckinVideoAdUpdateReq) (bool, int, error) {
	if c.CurrUserId <= 0 || c.Id == 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	updates := map[string]interface{}{}
	if c.AdPlatform != "" {
		updates["ad_platform"] = c.AdPlatform
	}
	if c.AdUnitId != "" {
		updates["ad_unit_id"] = c.AdUnitId
	}
	if c.AdOrderNo != "" {
		updates["ad_order_no"] = c.AdOrderNo
	}
	if c.VideoDuration != 0 {
		updates["video_duration"] = c.VideoDuration
	}
	if c.WatchDuration != 0 {
		updates["watch_duration"] = c.WatchDuration
	}
	if c.RewardAmount != "" {
		if v, err := models.ParseDecimal(c.RewardAmount); err == nil {
			updates["reward_amount"] = v
		}
	}
	if c.VerifyStatus != 0 {
		updates["verify_status"] = c.VerifyStatus
	}
	if len(updates) == 0 {
		return false, baseLang.SuccessCode, nil
	}
	if err := e.Orm.Model(&models.ChallengeCheckinVideoAd{}).Where("id = ?", c.Id).Updates(updates).Error; err != nil {
		return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}
	return true, baseLang.SuccessCode, nil
}

func (e *ChallengeCheckinVideoAd) Delete(id uint64, currUser int64) (bool, int, error) {
	if currUser <= 0 || id == 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if err := e.Orm.Delete(&models.ChallengeCheckinVideoAd{}, id).Error; err != nil {
		return false, baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	return true, baseLang.SuccessCode, nil
}
