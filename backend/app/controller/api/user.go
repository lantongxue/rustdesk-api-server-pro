package api

import "github.com/kataras/iris/v12/mvc"

type UserController struct {
}

func (c *UserController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/currentUser", "HandleCurrentUser")
}

func (c *UserController) HandleCurrentUser() mvc.Result {
	return mvc.Response{}
}

func (c *UserController) GetUsers() mvc.Result {

	return mvc.Response{}
}
