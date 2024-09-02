package api

type DeviceForm struct {
	Cpu        string `json:"cpu"`
	Hostname   string `json:"hostname"`
	RustdeskId string `json:"id"`
	Memory     string `json:"memory"`
	Os         string `json:"os"`
	Username   string `json:"username"`
	Uuid       string `json:"uuid"`
	Version    string `json:"version"`
}
