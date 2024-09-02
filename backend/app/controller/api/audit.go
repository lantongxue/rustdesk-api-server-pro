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

	_, _ = c.Db.Insert(audit)

	return mvc.Response{}
}
