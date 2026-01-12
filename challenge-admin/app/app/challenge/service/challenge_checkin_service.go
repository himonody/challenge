package service

import (
	"challenge-admin/app/app/challenge/models"
	"challenge-admin/app/app/challenge/service/dto"
	baseLang "challenge-admin/config/base/lang"
	cDto "challenge-admin/core/dto"
	"challenge-admin/core/dto/service"
	"challenge-admin/core/lang"
	"challenge-admin/core/middleware"
	"time"
)

type ChallengeCheckin struct {
	service.Service
}

// GetPage 分页查询打卡
func (e *ChallengeCheckin) GetPage(c *dto.ChallengeCheckinQueryReq, p *middleware.DataPermission) ([]models.ChallengeCheckin, int64, int, error) {
	var list []models.ChallengeCheckin
	var data models.ChallengeCheckin
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

// Create 新增
func (e *ChallengeCheckin) Create(c *dto.ChallengeCheckinCreateReq) (uint64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := models.ChallengeCheckin{
		ChallengeId: c.ChallengeId,
		UserId:      c.UserId,
		MoodCode:    c.MoodCode,
		MoodText:    c.MoodText,
		ContentType: c.ContentType,
		Status:      c.Status,
	}
	if t, err := time.Parse("2006-01-02", c.CheckinDate); err == nil {
		data.CheckinDate = t
	}
	if c.CheckinTime != "" {
		if t, err := time.Parse("2006-01-02 15:04:05", c.CheckinTime); err == nil {
			data.CheckinTime = &t
		}
	}
	if err := e.Orm.Create(&data).Error; err != nil {
		return 0, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}
	return data.Id, baseLang.SuccessCode, nil
}

// Update 更新
func (e *ChallengeCheckin) Update(c *dto.ChallengeCheckinUpdateReq) (bool, int, error) {
	if c.CurrUserId <= 0 || c.Id == 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	updates := map[string]interface{}{}
	if c.CheckinDate != "" {
		if t, err := time.Parse("2006-01-02", c.CheckinDate); err == nil {
			updates["checkin_date"] = t
		}
	}
	if c.CheckinTime != "" {
		if t, err := time.Parse("2006-01-02 15:04:05", c.CheckinTime); err == nil {
			updates["checkin_time"] = t
		}
	}
	if c.MoodCode != 0 {
		updates["mood_code"] = c.MoodCode
	}
	if c.MoodText != "" {
		updates["mood_text"] = c.MoodText
	}
	if c.ContentType != 0 {
		updates["content_type"] = c.ContentType
	}
	if c.Status != 0 {
		updates["status"] = c.Status
	}
	if len(updates) == 0 {
		return false, baseLang.SuccessCode, nil
	}
	if err := e.Orm.Model(&models.ChallengeCheckin{}).Where("id = ?", c.Id).Updates(updates).Error; err != nil {
		return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}
	return true, baseLang.SuccessCode, nil
}

// Delete 删除
func (e *ChallengeCheckin) Delete(id uint64, currUser int64) (bool, int, error) {
	if currUser <= 0 || id == 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if err := e.Orm.Delete(&models.ChallengeCheckin{}, id).Error; err != nil {
		return false, baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	return true, baseLang.SuccessCode, nil
}
