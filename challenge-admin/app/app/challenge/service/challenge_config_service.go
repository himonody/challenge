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

	"github.com/shopspring/decimal"
)

type ChallengeConfig struct {
	service.Service
}

// GetPage 分页查询活动配置
func (e *ChallengeConfig) GetPage(c *dto.ChallengeConfigQueryReq, p *middleware.DataPermission) ([]models.ChallengeConfig, int64, int, error) {
	var list []models.ChallengeConfig
	var data models.ChallengeConfig
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

// Create 新增配置
func (e *ChallengeConfig) Create(c *dto.ChallengeConfigCreateReq) (uint64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	amount, _ := decimal.NewFromString(c.Amount)
	pBonus, _ := decimal.NewFromString(c.PlatformBonus)
	cfg := models.ChallengeConfig{
		DayCount:      c.DayCount,
		Amount:        amount,
		PlatformBonus: pBonus,
		Status:        c.Status,
		Sort:          c.Sort,
	}
	if c.CheckinStart != "" {
		if t, err := time.Parse("2006-01-02 15:04:05", c.CheckinStart); err == nil {
			cfg.CheckinStart = &t
		}
	}
	if c.CheckinEnd != "" {
		if t, err := time.Parse("2006-01-02 15:04:05", c.CheckinEnd); err == nil {
			cfg.CheckinEnd = &t
		}
	}
	if err := e.Orm.Create(&cfg).Error; err != nil {
		return 0, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}
	return cfg.Id, baseLang.SuccessCode, nil
}

// Update 更新配置
func (e *ChallengeConfig) Update(c *dto.ChallengeConfigUpdateReq) (bool, int, error) {
	if c.CurrUserId <= 0 || c.Id == 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	updates := map[string]interface{}{}
	if c.DayCount > 0 {
		updates["day_count"] = c.DayCount
	}
	if c.Amount != "" {
		if v, err := decimal.NewFromString(c.Amount); err == nil {
			updates["amount"] = v
		}
	}
	if c.CheckinStart != "" {
		if t, err := time.Parse("2006-01-02 15:04:05", c.CheckinStart); err == nil {
			updates["checkin_start"] = t
		}
	}
	if c.CheckinEnd != "" {
		if t, err := time.Parse("2006-01-02 15:04:05", c.CheckinEnd); err == nil {
			updates["checkin_end"] = t
		}
	}
	if c.PlatformBonus != "" {
		if v, err := decimal.NewFromString(c.PlatformBonus); err == nil {
			updates["platform_bonus"] = v
		}
	}
	if c.Status != 0 {
		updates["status"] = c.Status
	}
	if c.Sort != 0 {
		updates["sort"] = c.Sort
	}
	if len(updates) == 0 {
		return false, baseLang.SuccessCode, nil
	}
	if err := e.Orm.Model(&models.ChallengeConfig{}).Where("id = ?", c.Id).Updates(updates).Error; err != nil {
		return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}
	return true, baseLang.SuccessCode, nil
}

// Delete 删除
func (e *ChallengeConfig) Delete(id uint64, currUser int64) (bool, int, error) {
	if currUser <= 0 || id == 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if err := e.Orm.Delete(&models.ChallengeConfig{}, id).Error; err != nil {
		return false, baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	return true, baseLang.SuccessCode, nil
}
