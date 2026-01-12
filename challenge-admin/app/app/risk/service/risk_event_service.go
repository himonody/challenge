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

type RiskEvent struct {
	service.Service
}

// GetRiskEventPage 分页查询事件
func (e *RiskEvent) GetRiskEventPage(c *dto.RiskEventQueryReq, p *middleware.DataPermission) ([]models.RiskEvent, int64, int, error) {
	var list []models.RiskEvent
	var data models.RiskEvent
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
