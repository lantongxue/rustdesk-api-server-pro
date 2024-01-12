package admin

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"
	"xorm.io/xorm"
)

type UsersController struct {
	basicController
}

func (c *UsersController) GetUserList() mvc.Result {
	currentPage := c.Ctx.URLParamIntDefault("page", 1)
	pageSize := c.Ctx.URLParamIntDefault("pageSize", 10)

	query := func() *xorm.Session {
		q := c.Db.Table(&model.User{})
		return q
	}

	pagination := db.NewPagination(currentPage, pageSize)
	userList := make([]model.User, 0)
	err := pagination.Paginate(query, &model.User{}, &userList)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}

	list := make([]iris.Map, 0)
	for _, u := range userList {
		list = append(list, iris.Map{
			"id":               u.Id,
			"username":         u.Username,
			"name":             u.Name,
			"email":            u.Email,
			"licensed_devices": u.LicensedDevices,
			"note":             u.Note,
			"status":           u.Status,
			"is_admin":         u.IsAdmin,
			"created_at":       u.CreatedAt.Format(config.TimeFormat),
		})
	}
	return c.Success(iris.Map{
		"total": pagination.TotalCount,
		"list":  list,
	}, "ok")
}
