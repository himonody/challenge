package service

import (
	"challenge/app/common/service/dto"

	"challenge/core/dto/service"
	"challenge/core/utils/captchautils"
)

type Common struct {
	service.Service
}

func (c *Common) Captcha() (dto.CaptchaResp, error) {
	id, b64s, _, err := captchautils.DriverStringFunc()
	if err != nil {
		c.Log.Errorf("common.common.Captcha DriverStringFunc error:%w", err)
		return dto.CaptchaResp{}, err
	}
	resp := dto.CaptchaResp{
		CaptchaId:   id,
		ImageBase64: b64s,
	}
	c.Log.Infof("common.common.Captcha DriverStringFunc data:%v", resp)
	return resp, nil
}
