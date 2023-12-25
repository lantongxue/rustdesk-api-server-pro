package api

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"rustdesk-api-server-pro/app/form"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/util"
	"time"
	"xorm.io/xorm"
)

type LoginController struct {
	Ctx iris.Context
	Db  *xorm.Engine
	Cfg *config.ServerConfig
}

func (c *LoginController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/login-options", "HandleLoginOptions")
}

func (c *LoginController) PostLogin() mvc.Result {
	var loginForm form.LoginForm
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

	// make other tokens expired
	_, err = c.Db.Where("user_id = ? and my_id = ? and uuid = ? and status = 1", user.Id, loginForm.Id, loginForm.Uuid).Cols("status").Update(&model.AuthToken{
		Status: 0,
	})

	signStr := loginForm.Id + loginForm.Uuid + user.Username + time.Now().String()
	token := util.HmacSha256(signStr, c.Cfg.SignKey)
	expired := 90 * 24 * time.Hour // 3 months

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

func (c *LoginController) HandleLoginOptions() mvc.Result {
	// this api returns about Open ID Connect options
	// returns like [ "common-oidc/[{'name':'<oidc name>', 'icon': '<oidc icon>'}]" ]
	// or [ "oidc/<oidc name>" ]
	// for rustdesk, it supports gitHub,gitlab,apple,auth0,azure,facebook,google,okta authentication
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
