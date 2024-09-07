package api

type AuditForm struct {
	Action    string   `json:"action"`
	ConnId    string   `json:"conn_id"`
	Id        string   `json:"id"`
	IP        string   `json:"ip"`
	SessionId string   `json:"session_id"`
	Uuid      string   `json:"uuid"`
	Note      string   `json:"note"`
	Peer      []string `json:"peer"`
	Type      int      `json:"type"`
}
