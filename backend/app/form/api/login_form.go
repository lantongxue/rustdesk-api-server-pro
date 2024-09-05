package api

type LoginForm struct {
	Username         string     `json:"username"`
	Password         string     `json:"password"`
	RustdeskId       string     `json:"id"`
	Uuid             string     `json:"uuid"`
	AutoLogin        bool       `json:"autoLogin"`
	Type             string     `json:"type"` // 取值范围：account mobile sms_code email_code tfa_code
	VerificationCode string     `json:"verificationCode"`
	TfaCode          string     `json:"tfaCode"`
	Secret           string     `json:"secret"`
	DeviceInfo       DeviceInfo `json:"deviceInfo"`
}

type DeviceInfo struct {
	OS   string `json:"os"`
	Type string `json:"type"`
	Name string `json:"name"`
}
