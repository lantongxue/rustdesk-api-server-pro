package admin

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"xorm.io/xorm"
)

type basicController struct {
	Ctx iris.Context
	Db  *xorm.Engine
}

func (c *basicController) Success(data interface{}, message string) mvc.Result {
	return mvc.Response{
		Object: iris.Map{
			"code":    200,
			"message": message,
			"data":    data,
		},
	}
}
