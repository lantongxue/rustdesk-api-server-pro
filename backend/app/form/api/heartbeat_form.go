package api

type HeartbeatForm struct {
	RustdeskId string `json:"id"`
	Uuid       string `json:"uuid"`
	ModifiedAt int64  `json:"modified_at"`
	Ver        int64  `json:"ver"`
	Version    string `json:"version"`
	Conns      []int  `json:"conns"`
}
