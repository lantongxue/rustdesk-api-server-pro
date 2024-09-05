package api

import (
	"rustdesk-api-server-pro/app/form/api"
	"rustdesk-api-server-pro/app/model"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type AddressBookTagController struct {
	basicController
}

func (c *AddressBookTagController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "ab/tags/{guid:string}", "HandleAbTags")
	b.Handle("POST", "ab/tag/add/{guid:string}", "HandleAbTagAdd")
	b.Handle("PUT", "ab/tag/update/{guid:string}", "HandleAbTagUpdate")
	b.Handle("PUT", "ab/tag/rename/{guid:string}", "HandleAbTagRename")
}

func (c *AddressBookTagController) HandleAbTags() mvc.Result {
	abGuid := c.Ctx.Params().Get("guid")

	user := c.GetUser()

	var ab model.AddressBook
	_, err := c.Db.Where("user_id = ? and guid = ?", user.Id, abGuid).Get(&ab)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}

	tags := make([]model.AddressBookTag, 0)
	err = c.Db.Where("user_id = ? and ab_id = ?", user.Id, ab.Id).Find(&tags)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}

	data := make([]iris.Map, 0)
	for _, tag := range tags {
		data = append(data, iris.Map{
			"name":  tag.Name,
			"color": tag.Color,
		})
	}

	return mvc.Response{
		Object: data,
	}
}

func (c *AddressBookTagController) HandleAbTagAdd() mvc.Result {
	abGuid := c.Ctx.Params().Get("guid")

	var form api.AbTagForm
	err := c.Ctx.ReadJSON(&form)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}

	user := c.GetUser()

	var ab model.AddressBook
	_, err = c.Db.Where("user_id = ? and guid = ?", user.Id, abGuid).Get(&ab)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}

	abTag := model.AddressBookTag{
		UserId: user.Id,
		AbId:   ab.Id,
		Name:   form.Name,
		Color:  form.Color,
	}
	_, err = c.Db.Insert(&abTag)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}
	return mvc.Response{
		Text: "OK",
	}
}

func (c *AddressBookTagController) HandleAbTagUpdate() mvc.Result {
	abGuid := c.Ctx.Params().Get("guid")

	var form api.AbTagForm
	err := c.Ctx.ReadJSON(&form)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}

	user := c.GetUser()
	var ab model.AddressBook
	_, err = c.Db.Where("user_id = ? and guid = ?", user.Id, abGuid).Get(&ab)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}

	_, err = c.Db.Where("user_id = ? and ab_id = ? and name = ?", user.Id, ab.Id, form.Name).Update(&model.AddressBookTag{
		Color: form.Color,
	})
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}
	return mvc.Response{
		Text: "OK",
	}
}

func (c *AddressBookTagController) HandleAbTagRename() mvc.Result {
	abGuid := c.Ctx.Params().Get("guid")

	var form api.AbTagRenameForm
	err := c.Ctx.ReadJSON(&form)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}

	user := c.GetUser()
	var ab model.AddressBook
	_, err = c.Db.Where("user_id = ? and guid = ?", user.Id, abGuid).Get(&ab)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}

	_, err = c.Db.Where("user_id = ? and ab_id = ? and name = ?", user.Id, ab.Id, form.Old).Update(&model.AddressBookTag{
		Name: form.New,
	})
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}
	return mvc.Response{
		Text: "OK",
	}
}
