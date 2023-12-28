package admin

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"rustdesk-api-server-pro/helper/captcha"
)

type AuthController struct {
	basicController
}

func (c *AuthController) PostAuthLogin() mvc.Result {
	return mvc.Response{}
}

func (c *AuthController) GetAuthCaptcha() mvc.Result {
	id, img := captcha.CreateCaptcha()
	return c.Success(iris.Map{
		"captchaID":  id,
		"captchaImg": img,
	}, "")
}
