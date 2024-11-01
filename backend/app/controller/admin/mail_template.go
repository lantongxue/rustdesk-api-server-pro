package admin

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"rustdesk-api-server-pro/app/form/admin"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"
	"xorm.io/xorm"
)

type MailTemplateController struct {
	basicController
}

func (c *MailTemplateController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/mail/templates/list", "HandleList")
	b.Handle("POST", "/mail/templates/add", "HandleAdd")
	b.Handle("POST", "/mail/templates/edit", "HandleEdit")
}

func (c *MailTemplateController) HandleList() mvc.Result {
	currentPage := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("size", 10)
	name := c.Ctx.URLParamDefault("name", "")
	subject := c.Ctx.URLParamDefault("subject", "")
	_type := c.Ctx.URLParamDefault("type", "")
	created_at_0 := c.Ctx.URLParamDefault("created_at[0]", "")
	created_at_1 := c.Ctx.URLParamDefault("created_at[1]", "")

	query := func() *xorm.Session {
		q := c.Db.Table(&model.MailTemplate{})
		if name != "" {
			name = "%" + name + "%"
			q.Where("name like ?", name)
		}
		if subject != "" {
			subject = "%" + subject + "%"
			q.Where("subject like ?", subject)
		}
		if _type != "" {
			q.Where("type = ?", _type)
		}
		if created_at_0 != "" && created_at_1 != "" {
			q.Where("created_at BETWEEN ? AND ?", created_at_0, created_at_1)
		}
		q.Desc("id")
		return q
	}

	pagination := db.NewPagination(currentPage, pageSize)
	templateList := make([]model.MailTemplate, 0)
	err := pagination.Paginate(query, &model.MailTemplate{}, &templateList)
	if err != nil {
		return c.Error(nil, err.Error())
	}

	list := make([]iris.Map, 0)
	for _, u := range templateList {
		list = append(list, iris.Map{
			"id":         u.Id,
			"name":       u.Name,
			"type":       u.Type,
			"subject":    u.Subject,
			"contents":   u.Contents,
			"created_at": u.CreatedAt.Format(config.TimeFormat),
		})
	}
	return c.Success(iris.Map{
		"total":   pagination.TotalCount,
		"records": list,
		"current": currentPage,
		"size":    pageSize,
	}, "ok")
}

func (c *MailTemplateController) HandleAdd() mvc.Result {
	var form admin.MailTemplateForm
	err := c.Ctx.ReadJSON(&form)
	if err != nil {
		return c.Error(nil, err.Error())
	}

	if form.Name == "" {
		return c.Error(nil, "MailTemplateNameEmpty")
	}
	if form.Subject == "" {
		return c.Error(nil, "MailTemplateSubjectEmpty")
	}
	if form.Contents == "" {
		return c.Error(nil, "MailTemplateContentsEmpty")
	}

	template := &model.MailTemplate{
		Name:     form.Name,
		Type:     form.Type,
		Subject:  form.Subject,
		Contents: form.Contents,
	}

	_, err = c.Db.Insert(template)
	if err != nil {
		return c.Error(nil, err.Error())
	}

	return c.Success(nil, "MailTemplateAddSuccess")
}

func (c *MailTemplateController) HandleEdit() mvc.Result {
	var form admin.MailTemplateForm
	err := c.Ctx.ReadJSON(&form)
	if err != nil {
		return c.Error(nil, err.Error())
	}

	if form.Id <= 0 {
		return c.Error(nil, "DataError")
	}
	template := &model.MailTemplate{
		Name:     form.Name,
		Type:     form.Type,
		Subject:  form.Subject,
		Contents: form.Contents,
	}

	_, err = c.Db.Where("id = ?", form.Id).Update(template)
	if err != nil {
		return c.Error(nil, err.Error())
	}

	return c.Success(nil, "MailTemplateUpdateSuccess")
}
