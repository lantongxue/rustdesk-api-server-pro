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
	act := c.Ctx.URLParamDefault("act", "")
	username := c.Ctx.URLParamDefault("username", "")

	query := func() *xorm.Session {
		q := c.Db.Table(&model.Audit{})
		q.Join("INNER", &model.User{}, "audit.user_id = user.id")
		if act != "" {
			q.Where("audit.action = ?", act)
		}
		if username != "" {
			username = "%" + username + "%"
			q.Where("user.username like ?", username)
		}
		return q
	}

	type Audit struct {
		model.Audit `xorm:"extends"`
		model.User  `xorm:"extends"`
	}

	pagination := db.NewPagination(currentPage, pageSize)
	auditList := make([]Audit, 0)
	err := pagination.Paginate(query, nil, &auditList)
	if err != nil {
		return c.Error(nil, err.Error())
	}

	list := make([]iris.Map, 0)
	for _, a := range auditList {
		list = append(list, iris.Map{
			"id":          a.Audit.Id,
			"username":    a.User.Username,
			"action":      a.Audit.Action,
			"conn_id":     a.Audit.ConnId,
			"rustdesk_id": a.Audit.MyId,
			"ip":          a.Audit.IP,
			"session_id":  a.Audit.SessionId,
			"uuid":        a.Audit.Uuid,
			"created_at":  a.Audit.CreatedAt.Format(config.TimeFormat),
		})
	}
	return c.Success(iris.Map{
		"total": pagination.TotalCount,
		"list":  list,
	}, "ok")
}
