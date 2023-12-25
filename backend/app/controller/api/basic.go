package api

import (
	"github.com/kataras/iris/v12"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"xorm.io/xorm"
)

type basicController struct {
	Ctx iris.Context
	Db  *xorm.Engine
}

func (c *basicController) GetUser() *model.User {
	return c.Ctx.Values().Get(config.CurrentUserKey).(*model.User)
}

func (c *basicController) GetToken() string {
	return c.Ctx.Values().Get(config.CurrentAuthTokenString).(string)
}

func (c *basicController) GetAuthToken() *model.AuthToken {
	return c.Ctx.Values().Get(config.CurrentAuthToken).(*model.AuthToken)
}
