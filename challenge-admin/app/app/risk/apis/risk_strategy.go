package apis

import (
	"challenge-admin/app/app/risk/service"
	"challenge-admin/app/app/risk/service/dto"
	baseLang "challenge-admin/config/base/lang"
	"challenge-admin/core/dto/api"
	"challenge-admin/core/lang"
	"challenge-admin/core/middleware"
	"challenge-admin/core/utils/dateutils"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type RiskStrategy struct {
	api.Api
}

// 策略分页
func (e RiskStrategy) RiskStrategyPage(c *gin.Context) {
	req := dto.RiskStrategyQueryReq{}
	s := service.RiskStrategy{}
	err := e.MakeContext(c).MakeOrm().Bind(&req).MakeService(&s.Service).Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	list, count, respCode, err := s.GetRiskStrategyPage(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.PageOK(list, nil, count, req.GetPageIndex(), req.GetPageSize(), lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// 策略导出
func (e RiskStrategy) RiskStrategyExport(c *gin.Context) {
	req := dto.RiskStrategyQueryReq{}
	s := service.RiskStrategy{}
	err := e.MakeContext(c).MakeOrm().Bind(&req).MakeService(&s.Service).Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	req.PageIndex = 1
	req.PageSize = 10000
	list, _, respCode, err := s.GetRiskStrategyPage(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	data, _ := s.ExportRiskStrategy(list)
	fileName := "risk_strategy_" + dateutils.ConvertToStr(time.Now(), 3) + ".xlsx"
	e.DownloadExcel(fileName, data)
}

// 新增
func (e RiskStrategy) RiskStrategyCreate(c *gin.Context) {
	req := dto.RiskStrategyCreateReq{}
	s := service.RiskStrategy{}
	err := e.MakeContext(c).MakeOrm().Bind(&req).MakeService(&s.Service).Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	req.CurrUserId = middleware.GetCurrUserId(c)
	id, code, err := s.CreateRiskStrategy(&req)
	if err != nil {
		e.Error(code, err.Error())
		return
	}
	e.OK(map[string]interface{}{"id": id}, lang.MsgByCode(code, e.Lang))
}

// 更新
func (e RiskStrategy) RiskStrategyUpdate(c *gin.Context) {
	req := dto.RiskStrategyUpdateReq{}
	s := service.RiskStrategy{}
	err := e.MakeContext(c).MakeOrm().Bind(&req).MakeService(&s.Service).Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	req.CurrUserId = middleware.GetCurrUserId(c)
	ok, code, err := s.UpdateRiskStrategy(&req)
	if err != nil {
		e.Error(code, err.Error())
		return
	}
	e.OK(ok, lang.MsgByCode(code, e.Lang))
}

// 删除
func (e RiskStrategy) RiskStrategyDelete(c *gin.Context) {
	idParam := c.Param("id")
	s := service.RiskStrategy{}
	err := e.MakeContext(c).MakeOrm().MakeService(&s.Service).Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	// simple parse to uint64
	var id uint64
	_, _ = fmt.Sscan(idParam, &id)
	ok, code, err := s.DeleteRiskStrategy(id, middleware.GetCurrUserId(c))
	if err != nil {
		e.Error(code, err.Error())
		return
	}
	e.OK(ok, lang.MsgByCode(code, e.Lang))
}
