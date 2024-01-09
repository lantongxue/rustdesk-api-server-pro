package api

type HeartbeatForm struct {
	Id         string `json:"id"`
	Uuid       string `json:"uuid"`
	ModifiedAt int64  `json:"modified_at"`
	Ver        int64  `json:"ver"`
}
