package api

import (
	"rustdesk-api-server-pro/app/form/api"
	"rustdesk-api-server-pro/app/service"
	"rustdesk-api-server-pro/config"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type LoginController struct {
	basicController
	Cfg *config.ServerConfig
}

func (c *LoginController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/login-options", "HandleLoginOptions")
}

func (c *LoginController) PostLogin() mvc.Result {
	var loginForm api.LoginForm
	err := c.Ctx.ReadJSON(&loginForm)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}

	userService := service.NewUserService()

	// {"type":"email_code","verificationCode":"666666","secret":""} // email
	if loginForm.Type == "email_code" && loginForm.VerificationCode != "" && loginForm.TfaCode == "" { // email_code 开始验证
		result := userService.LoginVerifyByEmailCode(loginForm)
		return mvc.Response{
			Object: result,
		}
	}

	// {"type":"email_code","verificationCode":"747332","tfaCode":"747332","secret":""} // 2fa
	if loginForm.Type == "email_code" && loginForm.TfaCode != "" && loginForm.VerificationCode == loginForm.TfaCode { // tfa_code 开始验证
		result := userService.LoginVerifyBy2FACode(loginForm)
		return mvc.Response{
			Object: result,
		}
	}

	result := userService.Login(loginForm)
	return mvc.Response{
		Object: result,
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
