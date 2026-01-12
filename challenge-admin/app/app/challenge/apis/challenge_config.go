package apis

import (
	"challenge-admin/app/app/challenge/service"
	"challenge-admin/app/app/challenge/service/dto"
	baseLang "challenge-admin/config/base/lang"
	"challenge-admin/core/dto/api"
	"challenge-admin/core/lang"
	"challenge-admin/core/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

type ChallengeConfig struct {
	api.Api
}

// 分页
func (e ChallengeConfig) ConfigPage(c *gin.Context) {
	req := dto.ChallengeConfigQueryReq{}
	s := service.ChallengeConfig{}
	err := e.MakeContext(c).MakeOrm().Bind(&req).MakeService(&s.Service).Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	list, count, respCode, err := s.GetPage(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.PageOK(list, nil, count, req.GetPageIndex(), req.GetPageSize(), lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// 导出
func (e ChallengeConfig) ConfigExport(c *gin.Context) {
	req := dto.ChallengeConfigQueryReq{}
	s := service.ChallengeConfig{}
	err := e.MakeContext(c).MakeOrm().Bind(&req).MakeService(&s.Service).Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	req.PageIndex = 1
	req.PageSize = 10000
	list, _, respCode, err := s.GetPage(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	data, _ := s.ExportConfig(list)
	fileName := "challenge_config_" + time.Now().Format("20060102150405") + ".xlsx"
	e.DownloadExcel(fileName, data)
}

// 新增
func (e ChallengeConfig) ConfigCreate(c *gin.Context) {
	req := dto.ChallengeConfigCreateReq{}
	s := service.ChallengeConfig{}
	err := e.MakeContext(c).MakeOrm().Bind(&req).MakeService(&s.Service).Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	req.CurrUserId = middleware.GetCurrUserId(c)
	id, code, err := s.Create(&req)
	if err != nil {
		e.Error(code, err.Error())
		return
	}
	e.OK(map[string]interface{}{"id": id}, lang.MsgByCode(code, e.Lang))
}

// 更新
func (e ChallengeConfig) ConfigUpdate(c *gin.Context) {
	req := dto.ChallengeConfigUpdateReq{}
	s := service.ChallengeConfig{}
	err := e.MakeContext(c).MakeOrm().Bind(&req).MakeService(&s.Service).Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	req.CurrUserId = middleware.GetCurrUserId(c)
	ok, code, err := s.Update(&req)
	if err != nil {
		e.Error(code, err.Error())
		return
	}
	e.OK(ok, lang.MsgByCode(code, e.Lang))
}

// 删除
func (e ChallengeConfig) ConfigDelete(c *gin.Context) {
	var uri struct {
		Id uint64 `uri:"id"`
	}
	s := service.ChallengeConfig{}
	err := e.MakeContext(c).MakeOrm().Bind(&uri).MakeService(&s.Service).Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	ok, code, err := s.Delete(uri.Id, middleware.GetCurrUserId(c))
	if err != nil {
		e.Error(code, err.Error())
		return
	}
	e.OK(ok, lang.MsgByCode(code, e.Lang))
}
