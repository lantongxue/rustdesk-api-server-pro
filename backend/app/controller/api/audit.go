package api

import (
	"io"
	"rustdesk-api-server-pro/app/model"
	"time"

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
	// 重要字段解析
	// type：0=控制会话，1=文件传输会话，2=TCP隧道会话

	body, err := io.ReadAll(c.Ctx.Request().Body)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}

	rustdeskId := gjson.GetBytes(body, "id").String()
	sessionId := gjson.GetBytes(body, "session_id").String()

	if gjson.GetBytes(body, "note").Exists() { // 只更新备注
		c.Db.Where("session_id = ?", sessionId).Update(&model.Audit{
			Note: gjson.GetBytes(body, "note").String(),
		})
		return mvc.Response{}
	}

	connId := int(gjson.GetBytes(body, "conn_id").Int())
	uuid := gjson.GetBytes(body, "uuid").String()

	actionResult := gjson.GetBytes(body, "action")
	if actionResult.Exists() {
		action := actionResult.String()
		if action == "new" {
			ip := gjson.GetBytes(body, "ip").String()
			c.Db.Insert(&model.Audit{
				ConnId:     connId,
				RustdeskId: rustdeskId,
				IP:         ip,
				SessionId:  sessionId,
				Uuid:       uuid,
			})
		}
		if action == "close" {
			c.Db.Where("conn_id = ?", connId).Update(&model.Audit{
				ClosedAt: time.Now(),
			})
		}
		return mvc.Response{}
	}

	peerResult := gjson.GetBytes(body, "peer")
	if peerResult.Exists() {
		t := gjson.GetBytes(body, "type").Int()
		_type := int(t)
		peer := peerResult.String()
		c.Db.Where("conn_id = ?", connId).Update(&model.Audit{
			SessionId: sessionId,
			Type:      _type,
			Peer:      peer,
		})
	}

	return mvc.Response{}
}

func (c *AuditController) PostAuditFile() mvc.Result {
	// {"id":"1235182932","info":"{\"files\":[[\"f1\",170398208],[\"f2\",147870720],[\"f3\",104590630]],\"ip\":\"192.168.100.170\",\"name\":\"mrkin\",\"num\":17778}","is_file":false,"path":"/Users/kali/Downloads/.nuget","peer_id":"182921366","type":1,"uuid":"NzIwQTBFMUYtRjg1OS01NjU0LUJCREUtMkNCMEU5MzQ5QzhF"}
	// {"id":"1235182932","info":"{\"files\":[[\"\",89]],\"ip\":\"192.168.100.170\",\"name\":\"mrkin\",\"num\":1}","is_file":true,"path":"/Users/kali/Downloads/.wslconfig","peer_id":"182921366","type":1,"uuid":"NzIwQTBFMUYtRjg1OS01NjU0LUJCREUtMkNCMEU5MzQ5QzhF"}
	// {"id":"1235182932","info":"{\"files\":[[\"\",34335]],\"ip\":\"192.168.100.170\",\"name\":\"mrkin\",\"num\":1}","is_file":true,"path":"/Users/kali/Downloads/weixin.jpg","peer_id":"182921366","type":0,"uuid":"NzIwQTBFMUYtRjg1OS01NjU0LUJCREUtMkNCMEU5MzQ5QzhF"}
	// 这个记录和conn的没关系
	// 重要字段解析
	// type：0=从被控端传输到主控端，1=从主控端传输到被控端

	body, err := io.ReadAll(c.Ctx.Request().Body)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}

	c.Db.Insert(&model.FileTransfer{
		RustdeskId: gjson.GetBytes(body, "id").String(),
		Info:       gjson.GetBytes(body, "info").String(),
		IsFile:     gjson.GetBytes(body, "is_file").Bool(),
		Path:       gjson.GetBytes(body, "path").String(),
		PeerId:     gjson.GetBytes(body, "peer_id").String(),
		Type:       int(gjson.GetBytes(body, "type").Int()),
		Uuid:       gjson.GetBytes(body, "type").String(),
	})

	return mvc.Response{}
}

func (c *AuditController) PostAuditAlarm() mvc.Result {

	return mvc.Response{
		Object: iris.Map{
			"error": "11",
		},
	}
}
