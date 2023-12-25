package api

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"rustdesk-api-server-pro/app/form"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/db"
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
	var f form.LoginForm
	err := c.Ctx.ReadJSON(&f)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}
	user := c.GetUser()
	token := c.GetToken()
	query := c.Db.Where("my_id = ? and uuid = ? and token = ?", f.Id, f.Uuid, token)
	if user != nil {
		query.Where("user_id = ?", user.Id)
	}
	var authToken model.AuthToken
	_, err = query.Get(&authToken)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}

	authToken.Status = 0
	_, err = c.Db.Where("id = ?", authToken.Id).Cols("status").Update(&authToken)
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
