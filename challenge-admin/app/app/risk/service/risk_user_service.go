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

type RiskUser struct {
	service.Service
}

// GetRiskUserPage 分页查询风控用户
func (e *RiskUser) GetRiskUserPage(c *dto.RiskUserQueryReq, p *middleware.DataPermission) ([]models.RiskUser, int64, int, error) {
	var list []models.RiskUser
	var data models.RiskUser
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
