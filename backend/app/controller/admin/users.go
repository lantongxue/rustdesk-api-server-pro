package admin

import (
	"rustdesk-api-server-pro/app/form/admin"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"
	"rustdesk-api-server-pro/util"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"xorm.io/xorm"
)

type UsersController struct {
	basicController
}

func (c *UsersController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/users/list", "HandleList")
	b.Handle("POST", "/users/add", "HandleAdd")
	b.Handle("POST", "/users/edit", "HandleEdit")
	b.Handle("POST", "/users/delete", "HandleDelete")
}

func (c *UsersController) HandleList() mvc.Result {
	currentPage := c.Ctx.URLParamIntDefault("page", 1)
	pageSize := c.Ctx.URLParamIntDefault("pageSize", 10)
	username := c.Ctx.URLParamDefault("username", "")
	name := c.Ctx.URLParamDefault("name", "")
	email := c.Ctx.URLParamDefault("email", "")
	admin_status := c.Ctx.URLParamDefault("admin_status", "")
	status := c.Ctx.URLParamDefault("status", "")

	query := func() *xorm.Session {
		q := c.Db.Table(&model.User{})
		if username != "" {
			q.Where("username = ?", username)
		}
		if name != "" {
			name = "%" + name + "%"
			q.Where("name like ?", name)
		}
		if email != "" {
			q.Where("email = ?", email)
		}
		if admin_status != "" {
			q.Where("is_admin = ?", admin_status)
		}
		if status != "" {
			q.Where("status = ?", status)
		}
		return q
	}

	pagination := db.NewPagination(currentPage, pageSize)
	userList := make([]model.User, 0)
	err := pagination.Paginate(query, &model.User{}, &userList)
	if err != nil {
		return c.Error(nil, err.Error())
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
		"total":   pagination.TotalCount,
		"records": list,
		"current": currentPage,
		"size":    pageSize,
	}, "ok")
}

func (c *UsersController) HandleAdd() mvc.Result {
	var form admin.UserForm
	err := c.Ctx.ReadJSON(&form)
	if err != nil {
		return c.Error(nil, err.Error())
	}

	if form.Username == "" {
		return c.Error(nil, "UsernameEmpty")
	}
	has, _ := c.Db.Where("username = ?", form.Username).Get(&model.User{})
	if has {
		return c.Error(nil, "UserExists")
	}

	if form.Password == "" {
		return c.Error(nil, "PasswordEmpty")
	}

	if form.Name == "" {
		form.Name = form.Username
	}

	if form.LicensedDevices < 0 {
		form.LicensedDevices = 0
	}

	p, err := util.Password(form.Password)
	if err != nil {
		return c.Error(nil, err.Error())
	}

	user := &model.User{
		Username:        form.Username,
		Password:        p,
		Name:            form.Name,
		Email:           form.Email,
		Note:            form.Note,
		LicensedDevices: form.LicensedDevices,
		Status:          form.Status,
		IsAdmin:         form.IsAdmin,
	}

	_, err = c.Db.Insert(user)
	if err != nil {
		return c.Error(nil, err.Error())
	}

	return c.Success(nil, "UserAddSuccess")
}

func (c *UsersController) HandleEdit() mvc.Result {
	var form admin.UserForm
	err := c.Ctx.ReadJSON(&form)
	if err != nil {
		return c.Error(nil, err.Error())
	}

	if form.Id <= 0 {
		return c.Error(nil, "DataError")
	}

	if form.LicensedDevices < 0 {
		form.LicensedDevices = 0
	}

	p := ""
	if form.Password != "" {
		p, err = util.Password(form.Password)
		if err != nil {
			return c.Error(nil, err.Error())
		}
	}

	if form.Name == "" {
		form.Name = form.Username
	}

	user := &model.User{
		Name:            form.Name,
		Email:           form.Email,
		Note:            form.Note,
		LicensedDevices: form.LicensedDevices,
		Status:          form.Status,
		IsAdmin:         form.IsAdmin,
	}

	if p != "" {
		user.Password = p
	}

	_, err = c.Db.Where("id = ?", form.Id).MustCols("licensed_devices", "status", "is_admin").Update(user)
	if err != nil {
		return c.Error(nil, err.Error())
	}

	return c.Success(nil, "UserUpdateSuccess")
}

func (c *UsersController) HandleDelete() mvc.Result {
	type deleteParams struct {
		Ids []int `json:"ids"`
	}
	var params deleteParams
	err := c.Ctx.ReadJSON(&params)
	if err != nil {
		return c.Error(nil, err.Error())
	}
	ids := util.RemoveElement(params.Ids, 1)
	_, err = c.Db.In("id", ids).Delete(&model.User{})
	if err != nil {
		return c.Error(nil, err.Error())
	}
	return c.Success(nil, "UserDeleteSuccess")
}
