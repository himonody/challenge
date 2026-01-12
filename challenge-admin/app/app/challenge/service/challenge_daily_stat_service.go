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

type ChallengeDailyStat struct {
	service.Service
}

func (e *ChallengeDailyStat) GetPage(c *dto.ChallengeDailyStatQueryReq, p *middleware.DataPermission) ([]models.ChallengeDailyStat, int64, int, error) {
	var list []models.ChallengeDailyStat
	var data models.ChallengeDailyStat
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

func (e *ChallengeDailyStat) Create(c *dto.ChallengeDailyStatCreateReq) (bool, int, error) {
	if c.CurrUserId <= 0 || c.StatDate == "" {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := models.ChallengeDailyStat{}
	if err := e.Orm.Table(data.TableName()).Create(&map[string]interface{}{
		"stat_date":        c.StatDate,
		"join_user_cnt":    c.JoinUserCnt,
		"success_user_cnt": c.SuccessUserCnt,
		"fail_user_cnt":    c.FailUserCnt,
		"join_amount":      c.JoinAmount,
		"success_amount":   c.SuccessAmount,
		"fail_amount":      c.FailAmount,
		"platform_bonus":   c.PlatformBonus,
		"pool_amount":      c.PoolAmount,
	}).Error; err != nil {
		return false, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}
	return true, baseLang.SuccessCode, nil
}

func (e *ChallengeDailyStat) Update(c *dto.ChallengeDailyStatUpdateReq) (bool, int, error) {
	if c.CurrUserId <= 0 || c.StatDate == "" {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	updates := map[string]interface{}{}
	if c.JoinUserCnt != 0 {
		updates["join_user_cnt"] = c.JoinUserCnt
	}
	if c.SuccessUserCnt != 0 {
		updates["success_user_cnt"] = c.SuccessUserCnt
	}
	if c.FailUserCnt != 0 {
		updates["fail_user_cnt"] = c.FailUserCnt
	}
	if c.JoinAmount != "" {
		updates["join_amount"] = c.JoinAmount
	}
	if c.SuccessAmount != "" {
		updates["success_amount"] = c.SuccessAmount
	}
	if c.FailAmount != "" {
		updates["fail_amount"] = c.FailAmount
	}
	if c.PlatformBonus != "" {
		updates["platform_bonus"] = c.PlatformBonus
	}
	if c.PoolAmount != "" {
		updates["pool_amount"] = c.PoolAmount
	}
	if len(updates) == 0 {
		return false, baseLang.SuccessCode, nil
	}
	var data models.ChallengeDailyStat
	if err := e.Orm.Table(data.TableName()).Where("stat_date = ?", c.StatDate).Updates(updates).Error; err != nil {
		return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}
	return true, baseLang.SuccessCode, nil
}

func (e *ChallengeDailyStat) Delete(statDate string, currUser int64) (bool, int, error) {
	if currUser <= 0 || statDate == "" {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	var data models.ChallengeDailyStat
	if err := e.Orm.Table(data.TableName()).Where("stat_date = ?", statDate).Delete(nil).Error; err != nil {
		return false, baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	return true, baseLang.SuccessCode, nil
}
