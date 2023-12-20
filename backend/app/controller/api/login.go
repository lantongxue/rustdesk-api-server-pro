package api

import "github.com/kataras/iris/v12/mvc"

type LoginController struct {
}

func (c *LoginController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/login-options", "HandleLoginOptions")
}

func (c *LoginController) PostLogin() mvc.Result {

	return mvc.Response{}
}

func (c *LoginController) PostLogout() mvc.Result {
	return mvc.Response{}
}

func (c *LoginController) HandleLoginOptions() mvc.Result {
	return mvc.Response{}
}
