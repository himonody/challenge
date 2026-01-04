package dto

type RegisterReq struct {
	UserName    string `json:"username"`
	Password    string `json:"password"`
	RefCode     string `json:"refCode"`
	CaptchaId   string `json:"captchaId"`
	CaptchaCode string `json:"captchaCode"`
}

type LoginReq struct {
	UserName    string `json:"username"`
	Password    string `json:"password"`
	CaptchaId   string `json:"captchaId"`
	CaptchaCode string `json:"captchaCode"`
}
