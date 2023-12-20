package api

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/util"
	"xorm.io/xorm"
)

type LoginForm struct {
	Username   string     `json:"username"`
	Password   string     `json:"password"`
	Id         string     `json:"id"`
	Uuid       string     `json:"uuid"`
	AutoLogin  bool       `json:"autoLogin"`
	Type       string     `json:"type"`
	DeviceInfo DeviceInfo `json:"deviceInfo"`
}

type DeviceInfo struct {
	OS   string `json:"os"`
	Type string `json:"type"`
	Name string `json:"name"`
}

type LoginController struct {
	Ctx iris.Context
	Db  *xorm.Engine
}

func (c *LoginController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/login-options", "HandleLoginOptions")
}

func (c *LoginController) PostLogin() mvc.Result {
	var loginForm LoginForm
	err := c.Ctx.ReadJSON(&loginForm)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}
	var user model.User
	get, err := c.Db.Where("username = ?", loginForm.Username).Get(&user)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}

	if !get {
		return mvc.Response{
			Object: iris.Map{
				"error": "user not exists",
			},
		}
	}

	if !util.PasswordVerify(loginForm.Password, user.Password) {
		return mvc.Response{
			Object: iris.Map{
				"error": "username or password error",
			},
		}
	}

	return mvc.Response{}
}

func (c *LoginController) PostLogout() mvc.Result {
	return mvc.Response{}
}

func (c *LoginController) HandleLoginOptions() mvc.Result {
	return mvc.Response{}
}
