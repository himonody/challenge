package dto

type CaptchaResp struct {
	CaptchaId   string `json:"captchaId"`
	ImageBase64 string `json:"imageBase64"`
}

type RegisterReq struct {
	UserName    string `json:"userName"`
	Password    string `json:"password"`
	RefCode     string `json:"refCode"`
	CaptchaId   string `json:"captchaId"`
	CaptchaCode string `json:"captchaCode"`
}

type LoginReq struct {
	UserName    string `json:"userName"`
	Password    string `json:"password"`
	CaptchaId   string `json:"captchaId"`
	CaptchaCode string `json:"captchaCode"`
}
