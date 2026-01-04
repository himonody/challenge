package apis

import (
	"challenge/app/common/service"
	"challenge/core/dto/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Common struct {
	api.Api
}

func (co *Common) Captcha(c *gin.Context) {
	s := &service.Common{}
	err := co.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		co.Error(http.StatusBadRequest, "captcha generate failed")
		return
	}
	resp, err := s.Captcha()
	if err != nil {
		co.Error(http.StatusInternalServerError, "captcha generate failed")
		return
	}
	co.OK(resp, "")
}
