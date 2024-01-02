package admin

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type UserController struct {
	basicController
}

func (c *UserController) GetUserInfo() mvc.Result {
	user := c.GetUser()
	return c.Success(iris.Map{
		"userId":   user.Id,
		"userName": user.Name,
	}, "ok")
}
