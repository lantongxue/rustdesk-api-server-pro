package api

import "github.com/kataras/iris/v12/mvc"

type PeerController struct {
}

func (c *PeerController) GetPeers() mvc.Result {
	return mvc.Response{}
}
