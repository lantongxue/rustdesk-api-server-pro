package admin

import (
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"xorm.io/xorm"
)

type AuditController struct {
	basicController
}

func (c *AuditController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/audit/list", "HandleList")
}

func (c *AuditController) HandleList() mvc.Result {
	currentPage := c.Ctx.URLParamIntDefault("page", 1)
	pageSize := c.Ctx.URLParamIntDefault("pageSize", 10)
	action := c.Ctx.URLParamDefault("action", "")
	conn_id := c.Ctx.URLParamDefault("conn_id", "")
	rustdesk_id := c.Ctx.URLParamDefault("rustdesk_id", "")
	ip := c.Ctx.URLParamDefault("ip", "")
	session_id := c.Ctx.URLParamDefault("session_id", "")
	uuid := c.Ctx.URLParamDefault("uuid", "")
	created_at_0 := c.Ctx.URLParamDefault("created_at[0]", "")
	created_at_1 := c.Ctx.URLParamDefault("created_at[1]", "")

	query := func() *xorm.Session {
		q := c.Db.Table(&model.Audit{})
		if action != "" {
			q.Where("audit.action = ?", action)
		}
		if conn_id != "" {
			q.Where("audit.conn_id = ?", conn_id)
		}
		if rustdesk_id != "" {
			q.Where("audit.rustdesk_id = ?", rustdesk_id)
		}
		if ip != "" {
			q.Where("audit.ip = ?", ip)
		}
		if session_id != "" {
			q.Where("audit.session_id = ?", session_id)
		}
		if uuid != "" {
			q.Where("audit.uuid = ?", uuid)
		}
		if created_at_0 != "" && created_at_1 != "" {
			q.Where("audit.created_at BETWEEN ? AND ?", created_at_0, created_at_1)
		}
		q.Desc("id")
		return q
	}

	pagination := db.NewPagination(currentPage, pageSize)
	auditList := make([]model.Audit, 0)
	err := pagination.Paginate(query, &model.Audit{}, &auditList)
	if err != nil {
		return c.Error(nil, err.Error())
	}

	list := make([]iris.Map, 0)
	for _, a := range auditList {
		list = append(list, iris.Map{
			"id":          a.Id,
			"action":      a.Action,
			"conn_id":     a.ConnId,
			"rustdesk_id": a.RustdeskId,
			"ip":          a.IP,
			"session_id":  a.SessionId,
			"uuid":        a.Uuid,
			"created_at":  a.CreatedAt.Format(config.TimeFormat),
		})
	}
	return c.Success(iris.Map{
		"total":   pagination.TotalCount,
		"records": list,
		"current": currentPage,
		"size":    pageSize,
	}, "ok")
}
