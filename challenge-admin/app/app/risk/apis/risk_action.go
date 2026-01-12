package apis

import (
	"challenge-admin/app/app/risk/service"
	"challenge-admin/app/app/risk/service/dto"
	baseLang "challenge-admin/config/base/lang"
	"challenge-admin/core/dto/api"
	"challenge-admin/core/lang"
	"challenge-admin/core/middleware"
	"challenge-admin/core/utils/dateutils"
	"time"

	"github.com/gin-gonic/gin"
)

type RiskAction struct {
	api.Api
}

// 动作列表
func (e RiskAction) RiskActionPage(c *gin.Context) {
	req := dto.RiskActionQueryReq{}
	s := service.RiskAction{}
	err := e.MakeContext(c).MakeOrm().Bind(&req).MakeService(&s.Service).Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	list, count, respCode, err := s.GetRiskActionPage(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.PageOK(list, nil, count, req.GetPageIndex(), req.GetPageSize(), lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// 动作导出
func (e RiskAction) RiskActionExport(c *gin.Context) {
	req := dto.RiskActionQueryReq{}
	s := service.RiskAction{}
	err := e.MakeContext(c).MakeOrm().Bind(&req).MakeService(&s.Service).Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	req.PageIndex = 1
	req.PageSize = 10000
	list, _, respCode, err := s.GetRiskActionPage(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	data, _ := s.ExportRiskAction(list)
	fileName := "risk_action_" + dateutils.ConvertToStr(time.Now(), 3) + ".xlsx"
	e.DownloadExcel(fileName, data)
}

// 新增
func (e RiskAction) RiskActionCreate(c *gin.Context) {
	req := dto.RiskActionCreateReq{}
	s := service.RiskAction{}
	err := e.MakeContext(c).MakeOrm().Bind(&req).MakeService(&s.Service).Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	req.CurrUserId = middleware.GetCurrUserId(c)
	ok, code, err := s.CreateRiskAction(&req)
	if err != nil {
		e.Error(code, err.Error())
		return
	}
	e.OK(ok, lang.MsgByCode(code, e.Lang))
}

// 更新
func (e RiskAction) RiskActionUpdate(c *gin.Context) {
	req := dto.RiskActionUpdateReq{}
	s := service.RiskAction{}
	err := e.MakeContext(c).MakeOrm().Bind(&req).MakeService(&s.Service).Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	req.CurrUserId = middleware.GetCurrUserId(c)
	ok, code, err := s.UpdateRiskAction(&req)
	if err != nil {
		e.Error(code, err.Error())
		return
	}
	e.OK(ok, lang.MsgByCode(code, e.Lang))
}

// 删除
func (e RiskAction) RiskActionDelete(c *gin.Context) {
	codeParam := c.Param("code")
	s := service.RiskAction{}
	err := e.MakeContext(c).MakeOrm().MakeService(&s.Service).Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	ok, code, err := s.DeleteRiskAction(codeParam, middleware.GetCurrUserId(c))
	if err != nil {
		e.Error(code, err.Error())
		return
	}
	e.OK(ok, lang.MsgByCode(code, e.Lang))
}
