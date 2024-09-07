package api

type AbForm struct {
	Data string `json:"data"`
}

type AbData struct {
	Tags      []string `json:"tags"`
	Peers     []AbPeer `json:"peers"`
	TagColors string   `json:"tag_colors"`
}

type AbPeer struct {
	Id               string   `json:"id"`
	Hash             string   `json:"hash"`
	Username         string   `json:"username"`
	Password         string   `json:"password"`
	Hostname         string   `json:"hostname"`
	Platform         string   `json:"platform"`
	Alias            string   `json:"alias"`
	Tags             []string `json:"tags"`
	ForceAlwaysRelay string   `json:"forceAlwaysRelay"`
	RdpPort          string   `json:"rdpPort"`
	RdpUsername      string   `json:"rdpUsername"`
	LoginName        string   `json:"loginName"`
	SameServer       string   `json:"same_server"`
}

type AbTagForm struct {
	Name  string `json:"name"`
	Color int64  `json:"color"`
}

type AbTagRenameForm struct {
	Old string `json:"old"`
	New string `json:"new"`
}
