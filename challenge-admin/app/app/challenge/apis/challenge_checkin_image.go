package apis

import (
	"challenge-admin/app/app/challenge/service"
	"challenge-admin/app/app/challenge/service/dto"
	baseLang "challenge-admin/config/base/lang"
	"challenge-admin/core/dto/api"
	"challenge-admin/core/lang"
	"challenge-admin/core/middleware"
	"challenge-admin/core/utils/dateutils"
	"time"

	"github.com/gin-gonic/gin"
)

type ChallengeCheckinImage struct {
	api.Api
}

func (e ChallengeCheckinImage) CheckinImagePage(c *gin.Context) {
	req := dto.ChallengeCheckinImageQueryReq{}
	s := service.ChallengeCheckinImage{}
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

func (e ChallengeCheckinImage) CheckinImageExport(c *gin.Context) {
	req := dto.ChallengeCheckinImageQueryReq{}
	s := service.ChallengeCheckinImage{}
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
	data, _ := s.ExportCheckinImage(list)
	fileName := "challenge_checkin_image_" + dateutils.ConvertToStr(time.Now(), 3) + ".xlsx"
	e.DownloadExcel(fileName, data)
}

func (e ChallengeCheckinImage) CheckinImageCreate(c *gin.Context) {
	req := dto.ChallengeCheckinImageCreateReq{}
	s := service.ChallengeCheckinImage{}
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

func (e ChallengeCheckinImage) CheckinImageUpdate(c *gin.Context) {
	req := dto.ChallengeCheckinImageUpdateReq{}
	s := service.ChallengeCheckinImage{}
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

func (e ChallengeCheckinImage) CheckinImageDelete(c *gin.Context) {
	var uri struct {
		Id uint64 `uri:"id"`
	}
	s := service.ChallengeCheckinImage{}
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
