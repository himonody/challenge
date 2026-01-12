package service

import (
	"challenge-admin/app/app/risk/models"
	"challenge-admin/app/app/risk/service/dto"
	baseLang "challenge-admin/config/base/lang"
	cDto "challenge-admin/core/dto"
	"challenge-admin/core/dto/service"
	"challenge-admin/core/lang"
	"challenge-admin/core/middleware"
	"errors"

	"gorm.io/gorm"
)

type RiskBlacklist struct {
	service.Service
}

// GetRiskBlacklistPage 分页查询黑名单
func (e *RiskBlacklist) GetRiskBlacklistPage(c *dto.RiskBlacklistQueryReq, p *middleware.DataPermission) ([]models.RiskBlacklist, int64, int, error) {
	var list []models.RiskBlacklist
	var data models.RiskBlacklist
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

// CreateRiskBlacklist 新增黑名单
func (e *RiskBlacklist) CreateRiskBlacklist(c *dto.RiskBlacklistCreateReq) (int64, int, error) {
	if c.CurrUserId <= 0 || c.Type == "" || c.Value == "" {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := models.RiskBlacklist{
		Type:      c.Type,
		Value:     c.Value,
		RiskLevel: c.RiskLevel,
		Reason:    c.Reason,
		Status:    c.Status,
	}
	if data.Status == "" {
		data.Status = "1"
	}
	err := e.Orm.Create(&data).Error
	if err != nil {
		return 0, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}
	return data.Id, baseLang.SuccessCode, nil
}

// UpdateRiskBlacklist 更新黑名单
func (e *RiskBlacklist) UpdateRiskBlacklist(c *dto.RiskBlacklistUpdateReq) (bool, int, error) {
	if c.CurrUserId <= 0 || c.Id <= 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := models.RiskBlacklist{}
	err := e.Orm.First(&data, c.Id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	updates := map[string]interface{}{}
	if c.Type != "" {
		updates["type"] = c.Type
	}
	if c.Value != "" {
		updates["value"] = c.Value
	}
	if c.RiskLevel > 0 {
		updates["risk_level"] = c.RiskLevel
	}
	if c.Reason != "" {
		updates["reason"] = c.Reason
	}
	if c.Status != "" {
		updates["status"] = c.Status
	}
	if len(updates) == 0 {
		return false, baseLang.SuccessCode, nil
	}
	err = e.Orm.Model(&data).Where("id=?", c.Id).Updates(updates).Error
	if err != nil {
		return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}
	return true, baseLang.SuccessCode, nil
}
