package api

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"time"
)

type SystemController struct {
}

func (c *SystemController) PostHeartbeat() mvc.Result {
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
