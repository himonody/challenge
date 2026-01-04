package dto

type CaptchaResp struct {
	CaptchaId   string `json:"captchaId"`
	ImageBase64 string `json:"imageBase64"`
}
