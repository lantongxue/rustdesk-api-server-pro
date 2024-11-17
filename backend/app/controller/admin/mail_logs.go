package admin

import (
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"xorm.io/xorm"
)

type MaiLogsController struct {
	basicController
}

func (c *MaiLogsController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/mail/logs/list", "HandleList")
	b.Handle("GET", "/mail/logs/info", "HandleInfo")
}

func (c *MaiLogsController) HandleList() mvc.Result {
	currentPage := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("size", 10)
	username := c.Ctx.URLParamDefault("username", "")
	uuid := c.Ctx.URLParamDefault("uuid", "")
	to := c.Ctx.URLParamDefault("to", "")
	status := c.Ctx.URLParamDefault("status", "")
	created_at_0 := c.Ctx.URLParamDefault("created_at[0]", "")
	created_at_1 := c.Ctx.URLParamDefault("created_at[1]", "")

	query := func() *xorm.Session {
		q := c.Db.Table(&model.MailLogs{})
		q.Join("INNER", &model.User{}, "mail_logs.user_id = user.id")
		if username != "" {
			q.Where("user.username = ?", username)
		}
		if uuid != "" {
			q.Where("mail_logs.uuid = ?", uuid)
		}
		if to != "" {
			q.Where("mail_logs.to = ?", to)
		}
		if status != "" {
			q.Where("mail_logs.status = ?", status)
		}

		if created_at_0 != "" && created_at_1 != "" {
			q.Where("mail_logs.created_at BETWEEN ? AND ?", created_at_0, created_at_1)
		}
		q.Desc("mail_logs.id")
		return q
	}

	type MailLog struct {
		model.MailLogs `xorm:"extends"`
		model.User     `xorm:"extends"`
	}

	pagination := db.NewPagination(currentPage, pageSize)
	mailLogList := make([]MailLog, 0)
	err := pagination.Paginate(query, &MailLog{}, &mailLogList)
	if err != nil {
		return c.Error(nil, err.Error())
	}

	list := make([]iris.Map, 0)
	for _, a := range mailLogList {
		list = append(list, iris.Map{
			"id":         a.MailLogs.Id,
			"username":   a.User.Username,
			"from":       a.MailLogs.From,
			"to":         a.MailLogs.To,
			"subject":    a.MailLogs.Subject,
			"uuid":       a.Uuid,
			"status":     a.MailLogs.Status,
			"created_at": a.MailLogs.CreatedAt.Format(config.TimeFormat),
		})
	}
	return c.Success(iris.Map{
		"total":   pagination.TotalCount,
		"records": list,
		"current": currentPage,
		"size":    pageSize,
	}, "ok")
}

func (c *MaiLogsController) HandleInfo() mvc.Result {

	return mvc.Response{
		Text: "213123",
	}
}
