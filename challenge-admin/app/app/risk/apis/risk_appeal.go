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

type RiskAppeal struct {
	api.Api
}

// 申诉
func (e RiskAppeal) RiskAppealPage(c *gin.Context) {
	req := dto.RiskAppealQueryReq{}
	s := service.RiskAppeal{}
	err := e.MakeContext(c).MakeOrm().Bind(&req).MakeService(&s.Service).Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	list, count, respCode, err := s.GetRiskAppealPage(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.PageOK(list, nil, count, req.GetPageIndex(), req.GetPageSize(), lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

func (e RiskAppeal) RiskAppealExport(c *gin.Context) {
	req := dto.RiskAppealQueryReq{}
	s := service.RiskAppeal{}
	err := e.MakeContext(c).MakeOrm().Bind(&req).MakeService(&s.Service).Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	req.PageIndex = 1
	req.PageSize = 10000
	list, _, respCode, err := s.GetRiskAppealPage(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	data, _ := s.ExportRiskAppeal(list)
	fileName := "risk_appeal_" + dateutils.ConvertToStr(time.Now(), 3) + ".xlsx"
	e.DownloadExcel(fileName, data)
}

func (e RiskAppeal) RiskAppealReview(c *gin.Context) {
	req := dto.RiskAppealReviewReq{}
	s := service.RiskAppeal{}
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
	req.ReviewerId = uid
	ok, respCode, err := s.ReviewRiskAppeal(&req)
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
