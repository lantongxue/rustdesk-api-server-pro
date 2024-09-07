package api

import (
	"rustdesk-api-server-pro/app/form/api"
	"rustdesk-api-server-pro/app/model"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type AuditController struct {
	basicController
}

func (c *AuditController) PostAuditConn() mvc.Result {

	// [INFO] 2024/09/07 16:33 ▶ POST:/api/audit/conn
	// Accept-Encoding: [gzip]
	// Content-Length: [70]
	// Content-Type: [application/json]
	// Accept: [*/*]
	// {"id":"1139987256","note":"21312aaa","session_id":1511599550828610773}
	// {"action":"new","conn_id":762,"id":"182921366","ip":"103.156.242.225","session_id":0,"uuid":"MTA3ZjAwNTYtZjk1Ny00OTJhLWJmM2MtODM5YTcyNzMwNDY2"}
	// {"action":"close","conn_id":762,"id":"182921366","session_id":17409556129324805845,"uuid":"MTA3ZjAwNTYtZjk1Ny00OTJhLWJmM2MtODM5YTcyNzMwNDY2"}
	// {"conn_id":762,"id":"182921366","peer":["1139987256","SYSTEM"],"session_id":17409556129324805845,"type":0,"uuid":"MTA3ZjAwNTYtZjk1Ny00OTJhLWJmM2MtODM5YTcyNzMwNDY2"}

	var auditForm api.AuditForm
	err := c.Ctx.ReadJSON(&auditForm)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}
	user := c.GetUser()
	audit := &model.Audit{
		UserId:     user.Id,
		Action:     auditForm.Action,
		ConnId:     auditForm.ConnId,
		RustdeskId: auditForm.Id,
		IP:         auditForm.IP,
		SessionId:  auditForm.SessionId,
		Uuid:       auditForm.Uuid,
	}

	c.Db.Insert(audit)

	return mvc.Response{}
}
