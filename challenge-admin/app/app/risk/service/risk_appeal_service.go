package service

import (
	"challenge-admin/app/app/risk/models"
	"challenge-admin/app/app/risk/service/dto"
	baseLang "challenge-admin/config/base/lang"
	cDto "challenge-admin/core/dto"
	"challenge-admin/core/dto/service"
	"challenge-admin/core/lang"
	"challenge-admin/core/middleware"
	"time"

	"gorm.io/gorm"
)

type RiskAppeal struct {
	service.Service
}

// GetRiskAppealPage 分页查询申诉
func (e *RiskAppeal) GetRiskAppealPage(c *dto.RiskAppealQueryReq, p *middleware.DataPermission) ([]models.RiskAppeal, int64, int, error) {
	var list []models.RiskAppeal
	var data models.RiskAppeal
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

// ReviewRiskAppeal 审核申诉
func (e *RiskAppeal) ReviewRiskAppeal(c *dto.RiskAppealReviewReq) (bool, int, error) {
	if c.Id <= 0 || c.ReviewerId <= 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := models.RiskAppeal{}
	err := e.Orm.First(&data, c.Id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return false, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	updates := map[string]interface{}{}
	if c.Status > 0 {
		updates["status"] = c.Status
	}
	if c.ReviewRemark != "" {
		updates["review_remark"] = c.ReviewRemark
	}
	if c.ActionResult > 0 {
		updates["action_result"] = c.ActionResult
	}
	if len(updates) == 0 {
		return false, baseLang.SuccessCode, nil
	}
	now := time.Now()
	updates["reviewer_id"] = c.ReviewerId
	updates["reviewed_at"] = &now
	err = e.Orm.Model(&data).Where("id=?", c.Id).Updates(updates).Error
	if err != nil {
		return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}
	return true, baseLang.SuccessCode, nil
}
