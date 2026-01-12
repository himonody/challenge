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

type ChallengeTotalStat struct {
	service.Service
}

func (e *ChallengeTotalStat) GetPage(c *dto.ChallengeTotalStatQueryReq, p *middleware.DataPermission) ([]models.ChallengeTotalStat, int64, int, error) {
	var list []models.ChallengeTotalStat
	var data models.ChallengeTotalStat
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

func (e *ChallengeTotalStat) Create(c *dto.ChallengeTotalStatCreateReq) (bool, int, error) {
	if c.CurrUserId <= 0 || c.Id == 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := models.ChallengeTotalStat{
		Id:                 c.Id,
		TotalUserCnt:       c.TotalUserCnt,
		TotalJoinCnt:       c.TotalJoinCnt,
		TotalSuccessCnt:    c.TotalSuccessCnt,
		TotalFailCnt:       c.TotalFailCnt,
		TotalJoinAmount:    models.MustParseDecimal(c.TotalJoinAmount),
		TotalSuccessAmount: models.MustParseDecimal(c.TotalSuccessAmount),
		TotalFailAmount:    models.MustParseDecimal(c.TotalFailAmount),
		TotalPlatformBonus: models.MustParseDecimal(c.TotalPlatformBonus),
		TotalPoolAmount:    models.MustParseDecimal(c.TotalPoolAmount),
	}
	if err := e.Orm.Create(&data).Error; err != nil {
		return false, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}
	return true, baseLang.SuccessCode, nil
}

func (e *ChallengeTotalStat) Update(c *dto.ChallengeTotalStatUpdateReq) (bool, int, error) {
	if c.CurrUserId <= 0 || c.Id == 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	updates := map[string]interface{}{}
	if c.TotalUserCnt != 0 {
		updates["total_user_cnt"] = c.TotalUserCnt
	}
	if c.TotalJoinCnt != 0 {
		updates["total_join_cnt"] = c.TotalJoinCnt
	}
	if c.TotalSuccessCnt != 0 {
		updates["total_success_cnt"] = c.TotalSuccessCnt
	}
	if c.TotalFailCnt != 0 {
		updates["total_fail_cnt"] = c.TotalFailCnt
	}
	if c.TotalJoinAmount != "" {
		updates["total_join_amount"] = models.MustParseDecimal(c.TotalJoinAmount)
	}
	if c.TotalSuccessAmount != "" {
		updates["total_success_amount"] = models.MustParseDecimal(c.TotalSuccessAmount)
	}
	if c.TotalFailAmount != "" {
		updates["total_fail_amount"] = models.MustParseDecimal(c.TotalFailAmount)
	}
	if c.TotalPlatformBonus != "" {
		updates["total_platform_bonus"] = models.MustParseDecimal(c.TotalPlatformBonus)
	}
	if c.TotalPoolAmount != "" {
		updates["total_pool_amount"] = models.MustParseDecimal(c.TotalPoolAmount)
	}
	if len(updates) == 0 {
		return false, baseLang.SuccessCode, nil
	}
	if err := e.Orm.Model(&models.ChallengeTotalStat{}).Where("id = ?", c.Id).Updates(updates).Error; err != nil {
		return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}
	return true, baseLang.SuccessCode, nil
}

func (e *ChallengeTotalStat) Delete(id int, currUser int64) (bool, int, error) {
	if currUser <= 0 || id == 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if err := e.Orm.Delete(&models.ChallengeTotalStat{}, id).Error; err != nil {
		return false, baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	return true, baseLang.SuccessCode, nil
}
