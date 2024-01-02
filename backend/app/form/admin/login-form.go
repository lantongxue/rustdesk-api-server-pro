package admin

type LoginForm struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Code      string `json:"code"`
	CaptchaId string `json:"captchaId"`
}
