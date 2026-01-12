package apis

import (
	"challenge-admin/app/app/risk/service"
	"challenge-admin/app/app/risk/service/dto"
	baseLang "challenge-admin/config/base/lang"
	"challenge-admin/core/dto/api"
	"challenge-admin/core/lang"
	"challenge-admin/core/middleware"
	"challenge-admin/core/middleware/auth"
	"challenge-admin/core/utils/dateutils"
	"time"

	"github.com/gin-gonic/gin"
)

type RiskBlacklist struct {
	api.Api
}

// 黑名单
func (e RiskBlacklist) RiskBlacklistPage(c *gin.Context) {
	req := dto.RiskBlacklistQueryReq{}
	s := service.RiskBlacklist{}
	err := e.MakeContext(c).MakeOrm().Bind(&req).MakeService(&s.Service).Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	list, count, respCode, err := s.GetRiskBlacklistPage(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.PageOK(list, nil, count, req.GetPageIndex(), req.GetPageSize(), lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

func (e RiskBlacklist) RiskBlacklistExport(c *gin.Context) {
	req := dto.RiskBlacklistQueryReq{}
	s := service.RiskBlacklist{}
	err := e.MakeContext(c).MakeOrm().Bind(&req).MakeService(&s.Service).Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	req.PageIndex = 1
	req.PageSize = 10000
	list, _, respCode, err := s.GetRiskBlacklistPage(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	data, _ := s.ExportRiskBlacklist(list)
	fileName := "risk_blacklist_" + dateutils.ConvertToStr(time.Now(), 3) + ".xlsx"
	e.DownloadExcel(fileName, data)
}

func (e RiskBlacklist) RiskBlacklistCreate(c *gin.Context) {
	req := dto.RiskBlacklistCreateReq{}
	s := service.RiskBlacklist{}
	err := e.MakeContext(c).MakeOrm().Bind(&req).MakeService(&s.Service).Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	uid, rCode, err := auth.Auth.GetUserId(c)
	if err != nil {
		e.Error(rCode, err.Error())
		return
	}
	req.CurrUserId = uid
	id, respCode, err := s.CreateRiskBlacklist(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(gin.H{"id": id}, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

func (e RiskBlacklist) RiskBlacklistUpdate(c *gin.Context) {
	req := dto.RiskBlacklistUpdateReq{}
	s := service.RiskBlacklist{}
	err := e.MakeContext(c).MakeOrm().Bind(&req).MakeService(&s.Service).Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	uid, rCode, err := auth.Auth.GetUserId(c)
	if err != nil {
		e.Error(rCode, err.Error())
		return
	}
	req.CurrUserId = uid
	ok, respCode, err := s.UpdateRiskBlacklist(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	if !ok {
		e.OK(nil, lang.MsgByCode(baseLang.DataNotUpdateCode, e.Lang))
		return
	}
	e.OK(nil, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}
