package api

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"rustdesk-api-server-pro/app/form/api"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"time"
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
			"modified_at": time.Now().Unix(),
		},
	}
}

func (c *SystemController) PostSysinfo() mvc.Result {
	return mvc.Response{
		Text: "SYSINFO_UPDATED",
	}
}
