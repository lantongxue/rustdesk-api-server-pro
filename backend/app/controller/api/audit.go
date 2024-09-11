package api

import (
	"io"
	"rustdesk-api-server-pro/app/model"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/tidwall/gjson"
)

type AuditController struct {
	basicController
}

func (c *AuditController) PostAuditConn() mvc.Result {

	// [INFO] 2024/09/07 16:33 ▶ POST:/api/audit/conn
	// Accept-Encoding: [gzip]
	// Content-Length: [70]
	// Content-Type: [application/json]
	// Accept: [*/*]
	// {"id":"1139987256","note":"21312aaa","session_id":1511599550828610773}
	// {"action":"new","conn_id":762,"id":"182921366","ip":"103.156.242.225","session_id":0,"uuid":"xxx"}
	// {"action":"close","conn_id":762,"id":"182921366","session_id":17409556129324805845,"uuid":"xxx"}
	// {"conn_id":762,"id":"182921366","peer":["1139987256","SYSTEM"],"session_id":17409556129324805845,"type":0,"uuid":"xxx"}

	body, err := io.ReadAll(c.Ctx.Request().Body)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}

	bodys := string(body)

	rustdeskId := gjson.Get(bodys, "id").String()
	sessionId := gjson.Get(bodys, "session_id").String()

	if gjson.Get(bodys, "note").Exists() { // 只更新备注
		c.Db.Where("rustdesk_id = ? and session_id = ?", rustdeskId, sessionId).Update(&model.Audit{
			Note: gjson.Get(bodys, "note").String(),
		})
		return mvc.Response{}
	}

	connId := int(gjson.Get(bodys, "conn_id").Int())
	uuid := gjson.Get(bodys, "uuid").String()

	if gjson.Get(bodys, "action").Exists() {
		action := gjson.Get(bodys, "action").String()
		if action == "new" {
			ip := gjson.Get(bodys, "ip").String()
			c.Db.Insert(&model.Audit{
				Action:     action,
				ConnId:     connId,
				RustdeskId: rustdeskId,
				IP:         ip,
				SessionId:  sessionId,
				Uuid:       uuid,
			})
		}
		if action == "close" {
			c.Db.Insert(&model.Audit{
				Action:     action,
				ConnId:     connId,
				RustdeskId: rustdeskId,
				SessionId:  sessionId,
				Uuid:       uuid,
			})
		}
		return mvc.Response{}
	}

	if gjson.Get(bodys, "peer").Exists() {
		t := gjson.Get(bodys, "type").Int()
		typ := int(t)
		peer := gjson.Get(bodys, "peer").String()
		c.Db.Insert(&model.Audit{
			ConnId:     connId,
			RustdeskId: rustdeskId,
			SessionId:  sessionId,
			Uuid:       uuid,
			Type:       typ,
			Peer:       peer,
		})
	}

	return mvc.Response{}
}
