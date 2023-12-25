package api

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	"github.com/kataras/iris/v12/mvc"
	"rustdesk-api-server-pro/app/form"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"xorm.io/xorm"
)

type UserController struct {
	Ctx    iris.Context
	Db     *xorm.Engine
	Verify *jwt.Verifier
}

func (c *UserController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/currentUser", "HandleCurrentUser")
}

func (c *UserController) HandleCurrentUser() mvc.Result {
	user := c.Ctx.Values().Get(config.CurrentUserKey).(*model.User)
	return mvc.Response{
		Object: iris.Map{
			"name":     user.Name,
			"email":    user.Email,
			"note":     user.Note,
			"status":   user.Status,
			"is_admin": false,
		},
	}
}

func (c *UserController) GetUsers() mvc.Result {

	return mvc.Response{}
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
	user := c.Ctx.Values().Get(config.CurrentUserKey).(*model.User)
	token := c.Ctx.Values().Get(config.CurrentAuthTokenString).(string)
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
