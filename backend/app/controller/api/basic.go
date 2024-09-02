package api

import (
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"

	"github.com/kataras/iris/v12"
	"xorm.io/xorm"
)

type basicController struct {
	Ctx iris.Context
	Db  *xorm.Engine
}

func (c *basicController) GetUser() *model.User {
	user := c.Ctx.Values().Get(config.CurrentUserKey)
	if user != nil {
		return user.(*model.User)
	}
	return nil
}

func (c *basicController) GetToken() string {
	return c.Ctx.Values().Get(config.CurrentAuthTokenString).(string)
}

func (c *basicController) GetAuthToken() *model.AuthToken {
	return c.Ctx.Values().Get(config.CurrentAuthToken).(*model.AuthToken)
}
