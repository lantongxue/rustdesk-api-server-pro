package admin

type UserForm struct {
	Id                  int    `json:"id"`
	Username            string `json:"username"`
	Password            string `json:"password"`
	Name                string `json:"name"`
	Email               string `json:"email"`
	LoginVerify         string `json:"login_verify"` // email_check tfa_check access_token
	TwoFactorAuthSecret string `json:"tfa_secret"`   // 2fa key
	Note                string `json:"note"`
	LicensedDevices     int    `json:"licensed_devices"`
	Status              int    `json:"status"`
	IsAdmin             bool   `json:"is_admin"`
}
