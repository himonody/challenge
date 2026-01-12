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

// RiskAction 动作字典
type RiskAction struct {
	service.Service
}

// GetRiskActionPage 分页查询动作
func (e *RiskAction) GetRiskActionPage(c *dto.RiskActionQueryReq, p *middleware.DataPermission) ([]models.RiskAction, int64, int, error) {
	var list []models.RiskAction
	var data models.RiskAction
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

// CreateRiskAction 新增
func (e *RiskAction) CreateRiskAction(c *dto.RiskActionCreateReq) (bool, int, error) {
	if c.CurrUserId <= 0 || c.Code == "" || c.Type == "" {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := models.RiskAction{
		Code:         c.Code,
		Type:         c.Type,
		DefaultValue: c.DefaultValue,
		Remark:       c.Remark,
	}
	if err := e.Orm.Create(&data).Error; err != nil {
		return false, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}
	return true, baseLang.SuccessCode, nil
}

// UpdateRiskAction 更新
func (e *RiskAction) UpdateRiskAction(c *dto.RiskActionUpdateReq) (bool, int, error) {
	if c.CurrUserId <= 0 || c.Code == "" {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	updates := map[string]interface{}{}
	if c.Type != "" {
		updates["type"] = c.Type
	}
	if c.DefaultValue != 0 {
		updates["default_value"] = c.DefaultValue
	}
	if c.Remark != "" {
		updates["remark"] = c.Remark
	}
	if len(updates) == 0 {
		return false, baseLang.SuccessCode, nil
	}
	err := e.Orm.Model(&models.RiskAction{}).Where("code = ?", c.Code).Updates(updates).Error
	if err != nil {
		return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}
	return true, baseLang.SuccessCode, nil
}

// DeleteRiskAction 删除
func (e *RiskAction) DeleteRiskAction(code string, currUserId int64) (bool, int, error) {
	if currUserId <= 0 || code == "" {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	err := e.Orm.Delete(&models.RiskAction{}, "code = ?", code).Error
	if err != nil {
		return false, baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	return true, baseLang.SuccessCode, nil
}
