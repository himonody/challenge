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

type ChallengeUser struct {
	service.Service
}

// GetPage 分页查询用户挑战
func (e *ChallengeUser) GetPage(c *dto.ChallengeUserQueryReq, p *middleware.DataPermission) ([]models.ChallengeUser, int64, int, error) {
	var list []models.ChallengeUser
	var data models.ChallengeUser
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
func (e *ChallengeUser) Create(c *dto.ChallengeUserCreateReq) (uint64, int, error) {
	if c.CurrUserId <= 0 || c.UserId == 0 {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := models.ChallengeUser{
		UserId:          c.UserId,
		ConfigId:        c.ConfigId,
		PoolId:          c.PoolId,
		ChallengeAmount: c.ChallengeAmount,
		StartDate:       c.StartDate,
		EndDate:         c.EndDate,
		Status:          c.Status,
		FailReason:      c.FailReason,
	}
	if err := e.Orm.Create(&data).Error; err != nil {
		return 0, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}
	return data.Id, baseLang.SuccessCode, nil
}

// Update 更新
func (e *ChallengeUser) Update(c *dto.ChallengeUserUpdateReq) (bool, int, error) {
	if c.CurrUserId <= 0 || c.Id == 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	updates := map[string]interface{}{}
	if c.PoolId != 0 {
		updates["pool_id"] = c.PoolId
	}
	if c.ChallengeAmount != "" {
		updates["challenge_amount"] = c.ChallengeAmount
	}
	if c.StartDate != 0 {
		updates["start_date"] = c.StartDate
	}
	if c.EndDate != 0 {
		updates["end_date"] = c.EndDate
	}
	if c.Status != 0 {
		updates["status"] = c.Status
	}
	if c.FailReason != 0 {
		updates["fail_reason"] = c.FailReason
	}
	if len(updates) == 0 {
		return false, baseLang.SuccessCode, nil
	}
	if err := e.Orm.Model(&models.ChallengeUser{}).Where("id = ?", c.Id).Updates(updates).Error; err != nil {
		return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}
	return true, baseLang.SuccessCode, nil
}

// Delete 删除
func (e *ChallengeUser) Delete(id uint64, currUser int64) (bool, int, error) {
	if currUser <= 0 || id == 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if err := e.Orm.Delete(&models.ChallengeUser{}, id).Error; err != nil {
		return false, baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	return true, baseLang.SuccessCode, nil
}
