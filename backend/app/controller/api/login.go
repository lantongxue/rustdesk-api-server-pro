package api

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	"github.com/kataras/iris/v12/mvc"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/util"
	"time"
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
	Cfg *config.ServerConfig
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
	expired := 7 * 24 * time.Hour // 7 days
	if loginForm.AutoLogin {      // for long time
		expired = 36500 * 24 * time.Hour // 100 years
	}
	tokenBytes, err := jwt.Sign(jwt.HS256, []byte(c.Cfg.JWT.SignKey), user, jwt.MaxAge(expired))
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}
	token := string(tokenBytes)

	authToken := &model.AuthToken{
		UserId:  user.Id,
		MyId:    loginForm.Id,
		Uuid:    loginForm.Uuid,
		Token:   token,
		Expired: time.Now().Add(expired),
		Status:  1,
	}

	_, err = c.Db.Insert(authToken)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}

	return mvc.Response{
		Object: iris.Map{
			"access_token": token,
			"type":         "access_token",
			"user": iris.Map{
				"name":     user.Name,
				"email":    user.Email,
				"note":     user.Note,
				"status":   user.Status,
				"is_admin": false,
			},
		},
	}
}

func (c *LoginController) PostLogout() mvc.Result {
	var form LoginForm
	err := c.Ctx.ReadJSON(&form)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}
	user := jwt.Get(c.Ctx).(*model.User)
	query := c.Db.Where("my_id = ? and uuid = ?", form.Id, form.Uuid)
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
	_, err = c.Db.Where("id = ?", authToken.Id).Update(&authToken)
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

func (c *LoginController) HandleLoginOptions() mvc.Result {
	// this api returns about Open ID Connect options
	// returns like [ "common-oidc/[{'name':'<oidc name>', 'icon': '<oidc icon>'}]" ]
	// or [ "oidc/<oidc name>" ]
	// for rustdesk, it supports github,gitlab,apple,auth0,azure,facebook,google,okta authentication
	// example:
	//oidc := []string{
	//	"oidc/github",
	//	"oidc/gitlab",
	//	"oidc/apple",
	//	"oidc/auth0",
	//	"oidc/azure",
	//	"oidc/facebook",
	//	"oidc/google",
	//	"oidc/okta",
	//}
	//return mvc.Response{
	//	Object: oidc,
	//}
	return mvc.Response{}
}
