package api

import "github.com/kataras/iris/v12/mvc"

type AuditController struct {
}

func (c *AuditController) PostAuditConn() mvc.Result {
	return mvc.Response{}
}
