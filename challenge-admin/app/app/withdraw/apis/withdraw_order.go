package apis

import (
	adminService "challenge-admin/app/admin/sys/service"
	"challenge-admin/app/app/withdraw/service"
	"challenge-admin/app/app/withdraw/service/dto"
	baseLang "challenge-admin/config/base/lang"
	"challenge-admin/core/dto/api"
	_ "challenge-admin/core/dto/response"
	"challenge-admin/core/lang"
	"challenge-admin/core/middleware"
	"challenge-admin/core/middleware/auth"
	"time"

	"github.com/gin-gonic/gin"
)

type WithdrawOrder struct {
	api.Api
}

// GetPage 获取提现订单分页列表
func (e WithdrawOrder) GetPage(c *gin.Context) {
	req := dto.WithdrawOrderQueryReq{}
	s := service.WithdrawOrder{}
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

// Get 获取提现订单详情
func (e WithdrawOrder) Get(c *gin.Context) {
	req := struct {
		Id int64 `uri:"id"`
	}{}
	s := service.WithdrawOrder{}
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

// UpdateStatus 审核/更新状态
func (e WithdrawOrder) UpdateStatus(c *gin.Context) {
	req := dto.WithdrawOrderUpdateStatusReq{}
	s := service.WithdrawOrder{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	uid, rCode, err := auth.Auth.GetUserId(c)
	if err != nil {
		e.Error(rCode, err.Error())
		return
	}
	req.ReviewerId = uid
	p := middleware.GetPermissionFromContext(c)
	ok, respCode, err := s.UpdateStatus(&req, p)
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

// Export 导出提现订单
func (e WithdrawOrder) Export(c *gin.Context) {
	req := dto.WithdrawOrderQueryReq{}
	s := service.WithdrawOrder{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}

	sysConfService := adminService.NewSysConfigService(&s.Service)
	maxSize, respCode, err := sysConfService.GetWithKeyInt("admin_sys_max_export_size")
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	req.PageIndex = 1
	req.PageSize = maxSize
	list, _, respCode, err := s.GetPage(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	data, _ := s.Export(list)
	fileName := "withdraw_order_" + time.Now().Format("20060102150405") + ".xlsx"
	e.DownloadExcel(fileName, data)
}
