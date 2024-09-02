package api

import (
	"rustdesk-api-server-pro/app/form/api"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
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

	var authToken model.AuthToken
	get, err := c.Db.Where("my_id = ? and uuid = ? and expired > ? and status = 1 and is_admin = 0", form.Id, form.Uuid, time.Now().Format(config.TimeFormat)).Get(&authToken)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}

	_, err = c.Db.Where("peer_id = ?", authToken.MyId).Update(&model.Peer{
		IsOnline: get,
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
			"modified_at": time.Now().Add(time.Duration(60) * time.Second).Unix(),
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
	has, err := c.Db.Where("uuid = ?", form.Uuid).Get(&device)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
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

	user := c.GetUser()
	if user != nil {
		device.UserId = user.Id
	}

	if has {
		c.Db.Where("id = ?", device.Id).Update(&device)
	} else {
		c.Db.Insert(&device)
	}

	return mvc.Response{
		Text: "SYSINFO_UPDATED",
	}
}
