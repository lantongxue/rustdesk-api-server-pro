package api

import (
	"rustdesk-api-server-pro/app/form/api"
	"rustdesk-api-server-pro/app/model"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type SystemController struct {
	basicController
}

func (c *SystemController) PostHeartbeat() mvc.Result {
	var form api.HeartbeatForm
	err := c.Ctx.ReadJSON(&form)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}

	var device model.Device
	has, err := c.Db.Where("rustdesk_id = ?", form.RustdeskId).Get(&device)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}

	if !has {
		device.RustdeskId = form.RustdeskId
		device.Uuid = form.Uuid
		device.Conns = form.Conns
		device.IsOnline = true
		_, err = c.Db.Insert(&device)
		if err != nil {
			return mvc.Response{
				Object: iris.Map{
					"error": err.Error(),
				},
			}
		}
	}

	_, err = c.Db.Where("rustdesk_id = ?", form.RustdeskId).Cols("is_online", "conns").Update(&model.Device{
		IsOnline: true,
		Conns:    form.Conns,
	})
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}

	return mvc.Response{
		Object: iris.Map{
			"modified_at": time.Now().Unix(),
			//"disconnect":  0,
			//"strategy":    iris.Map{},
		},
	}
}

func (c *SystemController) PostSysinfo() mvc.Result {
	var form api.DeviceForm
	err := c.Ctx.ReadJSON(&form)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}

	var device model.Device
	has, err := c.Db.Where("rustdesk_id = ?", form.RustdeskId).Get(&device)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}

	if !has {
		return mvc.Response{
			Text: "ID_NOT_FOUND",
		}
	}

	device.Cpu = form.Cpu
	device.Hostname = form.Hostname
	device.RustdeskId = form.RustdeskId
	device.Memory = form.Memory
	device.Os = form.Os
	device.Username = form.Username
	device.Uuid = form.Uuid
	device.Version = form.Version

	c.Db.Where("id = ?", device.Id).Update(&device)

	return mvc.Response{
		Text: "SYSINFO_UPDATED",
	}
}
