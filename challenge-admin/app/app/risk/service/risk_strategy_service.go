package service

import (
	"challenge-admin/app/app/risk/models"
	"challenge-admin/app/app/risk/service/dto"
	baseLang "challenge-admin/config/base/lang"
	cDto "challenge-admin/core/dto"
	"challenge-admin/core/dto/service"
	"challenge-admin/core/lang"
	"challenge-admin/core/middleware"
)

// RiskStrategy 策略
type RiskStrategy struct {
	service.Service
}

// GetRiskStrategyPage 分页查询策略
func (e *RiskStrategy) GetRiskStrategyPage(c *dto.RiskStrategyQueryReq, p *middleware.DataPermission) ([]models.RiskStrategy, int64, int, error) {
	var list []models.RiskStrategy
	var data models.RiskStrategy
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

// CreateRiskStrategy 新增
func (e *RiskStrategy) CreateRiskStrategy(c *dto.RiskStrategyCreateReq) (uint64, int, error) {
	if c.CurrUserId <= 0 || c.Scene == "" || c.RuleCode == "" || c.IdentityType == "" || c.WindowSeconds <= 0 || c.Threshold <= 0 || c.Action == "" {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := models.RiskStrategy{
		Scene:         c.Scene,
		RuleCode:      c.RuleCode,
		IdentityType:  c.IdentityType,
		WindowSeconds: c.WindowSeconds,
		Threshold:     c.Threshold,
		Action:        c.Action,
		ActionValue:   c.ActionValue,
		Status:        c.Status,
		Priority:      c.Priority,
		Remark:        c.Remark,
	}
	if data.Status == 0 {
		data.Status = 1
	}
	if data.Priority == 0 {
		data.Priority = 100
	}
	if err := e.Orm.Create(&data).Error; err != nil {
		return 0, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}
	return data.Id, baseLang.SuccessCode, nil
}

// UpdateRiskStrategy 更新
func (e *RiskStrategy) UpdateRiskStrategy(c *dto.RiskStrategyUpdateReq) (bool, int, error) {
	if c.CurrUserId <= 0 || c.Id == 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	updates := map[string]interface{}{}
	if c.Scene != "" {
		updates["scene"] = c.Scene
	}
	if c.RuleCode != "" {
		updates["rule_code"] = c.RuleCode
	}
	if c.IdentityType != "" {
		updates["identity_type"] = c.IdentityType
	}
	if c.WindowSeconds > 0 {
		updates["window_seconds"] = c.WindowSeconds
	}
	if c.Threshold > 0 {
		updates["threshold"] = c.Threshold
	}
	if c.Action != "" {
		updates["action"] = c.Action
	}
	if c.ActionValue != 0 {
		updates["action_value"] = c.ActionValue
	}
	if c.Status != 0 {
		updates["status"] = c.Status
	}
	if c.Priority != 0 {
		updates["priority"] = c.Priority
	}
	if c.Remark != "" {
		updates["remark"] = c.Remark
	}
	if len(updates) == 0 {
		return false, baseLang.SuccessCode, nil
	}
	err := e.Orm.Model(&models.RiskStrategy{}).Where("id = ?", c.Id).Updates(updates).Error
	if err != nil {
		return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}
	return true, baseLang.SuccessCode, nil
}

// DeleteRiskStrategy 删除
func (e *RiskStrategy) DeleteRiskStrategy(id uint64, currUserId int64) (bool, int, error) {
	if currUserId <= 0 || id == 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	err := e.Orm.Delete(&models.RiskStrategy{}, "id = ?", id).Error
	if err != nil {
		return false, baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	return true, baseLang.SuccessCode, nil
}
