package admin

import (
	"rustdesk-api-server-pro/helper/captcha"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
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
		"id":  id,
		"img": img,
	}, "")
}
