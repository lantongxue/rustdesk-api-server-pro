package api

import (
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/db"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"xorm.io/xorm"
)

type PeerController struct {
	basicController
}

func (c *PeerController) GetPeers() mvc.Result {
	current := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("pageSize", 10)
	status := c.Ctx.URLParamIntDefault("status", 1)

	user := c.GetUser()

	query := func() *xorm.Session {
		q := c.Db.Table(&model.Peer{}).Where("user_id = ?", user.Id)
		return q.Desc("id")
	}

	pagination := db.NewPagination(current, pageSize)
	peerList := make([]model.Peer, 0)
	err := pagination.Paginate(query, &model.Peer{}, &peerList)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}
	data := make([]iris.Map, 0)
	for _, p := range peerList {
		data = append(data, iris.Map{
			"id": p.RustdeskId,
			"info": iris.Map{
				"username":    p.Username,
				"os":          p.Platform,
				"device_name": p.Hostname,
			},
			"status":    status,
			"user":      user.Username,
			"user_name": p.LoginName,
			"note":      p.Note,
		})
	}
	return mvc.Response{
		Object: iris.Map{
			"total": pagination.TotalCount,
			"data":  data,
		},
	}
}
