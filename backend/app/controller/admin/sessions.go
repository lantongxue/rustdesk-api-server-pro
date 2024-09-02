package admin

import (
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"
	"rustdesk-api-server-pro/util"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"xorm.io/xorm"
)

type SessionsController struct {
	basicController
}

func (c *SessionsController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/sessions/list", "HandleList")
	b.Handle("POST", "/sessions/kill", "HandleKill")
}

func (c *SessionsController) HandleList() mvc.Result {
	currentPage := c.Ctx.URLParamIntDefault("page", 1)
	pageSize := c.Ctx.URLParamIntDefault("pageSize", 10)
	username := c.Ctx.URLParamDefault("username", "")
	created_at_0 := c.Ctx.URLParamDefault("created_at[0]", "")
	created_at_1 := c.Ctx.URLParamDefault("created_at[1]", "")
	query := func() *xorm.Session {
		q := c.Db.Table(&model.AuthToken{})
		q.Join("INNER", &model.User{}, "auth_token.user_id = user.id")
		q.Where("auth_token.status = 1 and auth_token.is_admin = 0")
		if username != "" {
			q.Where("user.username = ?", username)
		}
		if created_at_0 != "" && created_at_1 != "" {
			q.Where("created_at BETWEEN ? AND ?", created_at_0, created_at_1)
		}
		return q
	}

	type Session struct {
		model.AuthToken `xorm:"extends"`
		model.User      `xorm:"extends"`
	}

	pagination := db.NewPagination(currentPage, pageSize)
	sessionList := make([]Session, 0)
	err := pagination.Paginate(query, &Session{}, &sessionList)
	if err != nil {
		return c.Error(nil, err.Error())
	}

	list := make([]iris.Map, 0)
	for _, s := range sessionList {
		list = append(list, iris.Map{
			"id":          s.AuthToken.Id,
			"username":    s.User.Username,
			"rustdesk_id": s.AuthToken.RustdeskId,
			"expired":     s.AuthToken.Expired.Format(config.TimeFormat),
			"created_at":  s.AuthToken.CreatedAt.Format(config.TimeFormat),
		})
	}
	return c.Success(iris.Map{
		"total":   pagination.TotalCount,
		"records": list,
		"current": currentPage,
		"size":    pageSize,
	}, "ok")
}

func (c *SessionsController) HandleKill() mvc.Result {
	type killParams struct {
		Ids []int `json:"ids"`
	}
	var params killParams
	err := c.Ctx.ReadJSON(&params)
	if err != nil {
		return c.Error(nil, err.Error())
	}
	ids := util.RemoveElement(params.Ids, 1)
	_, err = c.Db.In("id", ids).Cols("status").Update(&model.AuthToken{
		Status: 0,
	})
	if err != nil {
		return c.Error(nil, err.Error())
	}
	return c.Success(nil, "SessionKillSuccess")
}
