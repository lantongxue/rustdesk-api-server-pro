package api

type LoginForm struct {
	Username   string     `json:"username"`
	Password   string     `json:"password"`
	Id         string     `json:"id"`
	Uuid       string     `json:"uuid"`
	AutoLogin  bool       `json:"autoLogin"`
	Type       string     `json:"type"`
	DeviceInfo DeviceInfo `json:"deviceInfo"`
}

type DeviceInfo struct {
	OS   string `json:"os"`
	Type string `json:"type"`
	Name string `json:"name"`
}
