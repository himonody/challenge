package apis

import (
	"challenge-admin/app/plugins/msg/service"
	"challenge-admin/app/plugins/msg/service/dto"
	baseLang "challenge-admin/config/base/lang"
	"challenge-admin/core/dto/api"
	_ "challenge-admin/core/dto/response"
	"challenge-admin/core/lang"
	"challenge-admin/core/middleware"
	"github.com/gin-gonic/gin"
)

type MsgCode struct {
	api.Api
}

// GetPage plugins-获取验证码管理分页列表
func (e MsgCode) GetPage(c *gin.Context) {
	req := dto.MsgCodeQueryReq{}
	s := service.MsgCode{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
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

// Get plugins-获取验证码管理详情
func (e MsgCode) Get(c *gin.Context) {
	req := dto.MsgCodeGetReq{}
	s := service.MsgCode{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	result, respCode, err := s.Get(req.Id, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(result, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}
