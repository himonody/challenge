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

type ChallengePool struct {
	service.Service
}

func (e *ChallengePool) GetPage(c *dto.ChallengePoolQueryReq, p *middleware.DataPermission) ([]models.ChallengePool, int64, int, error) {
	var list []models.ChallengePool
	var data models.ChallengePool
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

func (e *ChallengePool) Create(c *dto.ChallengePoolCreateReq) (uint64, int, error) {
	if c.CurrUserId <= 0 || c.ConfigId == 0 {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := models.ChallengePool{
		ConfigId:    c.ConfigId,
		TotalAmount: c.TotalAmount,
		Settled:     c.Settled,
	}
	if c.StartDate != "" {
		if t, err := time.Parse("2006-01-02 15:04:05", c.StartDate); err == nil {
			data.StartDate = &t
		}
	}
	if c.EndDate != "" {
		if t, err := time.Parse("2006-01-02 15:04:05", c.EndDate); err == nil {
			data.EndDate = &t
		}
	}
	if err := e.Orm.Create(&data).Error; err != nil {
		return 0, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}
	return data.Id, baseLang.SuccessCode, nil
}

func (e *ChallengePool) Update(c *dto.ChallengePoolUpdateReq) (bool, int, error) {
	if c.CurrUserId <= 0 || c.Id == 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	updates := map[string]interface{}{}
	if c.StartDate != "" {
		if t, err := time.Parse("2006-01-02 15:04:05", c.StartDate); err == nil {
			updates["start_date"] = t
		}
	}
	if c.EndDate != "" {
		if t, err := time.Parse("2006-01-02 15:04:05", c.EndDate); err == nil {
			updates["end_date"] = t
		}
	}
	if c.TotalAmount != "" {
		updates["total_amount"] = c.TotalAmount
	}
	if c.Settled != 0 {
		updates["settled"] = c.Settled
	}
	if len(updates) == 0 {
		return false, baseLang.SuccessCode, nil
	}
	if err := e.Orm.Model(&models.ChallengePool{}).Where("id = ?", c.Id).Updates(updates).Error; err != nil {
		return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}
	return true, baseLang.SuccessCode, nil
}

func (e *ChallengePool) Delete(id uint64, currUser int64) (bool, int, error) {
	if currUser <= 0 || id == 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if err := e.Orm.Delete(&models.ChallengePool{}, id).Error; err != nil {
		return false, baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	return true, baseLang.SuccessCode, nil
}
