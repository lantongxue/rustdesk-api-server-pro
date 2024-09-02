package api

import (
	"rustdesk-api-server-pro/app/form/api"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/db"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"xorm.io/xorm"
)

type UserController struct {
	basicController
}

func (c *UserController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/currentUser", "HandleCurrentUser")
}

func (c *UserController) HandleCurrentUser() mvc.Result {
	user := c.GetUser()
	return mvc.Response{
		Object: iris.Map{
			"name":     user.Name,
			"email":    user.Email,
			"note":     user.Note,
			"status":   user.Status,
			"is_admin": user.IsAdmin,
		},
	}
}

func (c *UserController) GetUsers() mvc.Result {
	user := c.GetUser()
	if !user.IsAdmin {
		return mvc.Response{
			Object: iris.Map{
				"error": "not admin",
			},
		}
	}
	current := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("pageSize", 10)
	status := c.Ctx.URLParamIntDefault("status", 1)

	query := func() *xorm.Session {
		q := c.Db.Table(&model.User{}).Where("status = ?", status)
		return q.Desc("id")
	}

	pagination := db.NewPagination(current, pageSize)
	userList := make([]model.User, 0)
	err := pagination.Paginate(query, &model.User{}, &userList)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}
	data := make([]iris.Map, 0)
	for _, u := range userList {
		data = append(data, iris.Map{
			"name":     u.Name,
			"email":    u.Email,
			"note":     u.Note,
			"status":   u.Status,
			"is_admin": u.IsAdmin,
		})
	}
	return mvc.Response{
		Object: iris.Map{
			"total": pagination.TotalCount,
			"data":  data,
		},
	}
}

func (c *UserController) PostLogout() mvc.Result {
	var f api.LoginForm
	err := c.Ctx.ReadJSON(&f)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}
	_, err = c.Db.Where("rustdesk_id = ?", f.RustdeskId).Cols("status").Update(&model.AuthToken{
		Status: 0,
	})
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}
	return mvc.Response{
		Text: "ok",
	}
}
