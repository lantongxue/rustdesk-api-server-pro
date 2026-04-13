package api

import (
	"rustdesk-api-server-pro/app/form/api"
	"rustdesk-api-server-pro/app/model"
	versionhelper "rustdesk-api-server-pro/helper/version"
	"strconv"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type SystemController struct {
	basicController
}

func ResolveHeartbeatRustdeskID(id, uuid string) string {
	if id != "" {
		return id
	}
	return uuid
}

func NormalizeReportedVersion(version string, ver int64) string {
	if version != "" {
		return version
	}
	if ver > 0 {
		return strconv.FormatInt(ver, 10)
	}
	return ""
}

func (c *SystemController) PostHeartbeat() mvc.Result {
	// {"conns":[762],"id":"182921366","modified_at":1725698100,"uuid":"xxx","ver":1002070}
	var form api.HeartbeatForm
	err := c.Ctx.ReadJSON(&form)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}
	form.RustdeskId = ResolveHeartbeatRustdeskID(form.RustdeskId, form.Uuid)
	if form.RustdeskId == "" {
		return mvc.Response{
			Object: iris.Map{
				"error": "missing id",
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
		device.Conns = len(form.Conns)
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
		Conns:    len(form.Conns),
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
			"strategy":    versionhelper.ResolveCapability(NormalizeReportedVersion(form.Version, form.Ver)),
			//"disconnect":  0,
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
	form.RustdeskId = ResolveHeartbeatRustdeskID(form.RustdeskId, form.Uuid)
	if form.RustdeskId == "" {
		return mvc.Response{
			Text: "ID_NOT_FOUND",
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
	device.Version = NormalizeReportedVersion(form.Version, form.Ver)

	c.Db.Where("id = ?", device.Id).Update(&device)

	return mvc.Response{
		Text: "SYSINFO_UPDATED",
	}
}
