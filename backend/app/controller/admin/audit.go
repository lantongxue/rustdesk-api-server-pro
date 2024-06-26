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
	kw := c.Ctx.URLParamDefault("kw", "")

	query := func() *xorm.Session {
		q := c.Db.Table(&model.Audit{})
		q.Join("INNER", &model.User{}, "audit.user_id = user.id")
		if kw != "" {
			q.Or("audit.action = ?", kw)
			q.Or("audit.my_id = ?", kw)
			q.Or("audit.ip = ?", kw)
			q.Or("user.username like ?", "%"+kw+"%")
		}
		return q
	}

	type Audit struct {
		model.Audit `xorm:"extends"`
		model.User  `xorm:"extends"`
	}

	pagination := db.NewPagination(currentPage, pageSize)
	auditList := make([]Audit, 0)
	err := pagination.Paginate(query, &Audit{}, &auditList)
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
