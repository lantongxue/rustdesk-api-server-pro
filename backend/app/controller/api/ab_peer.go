package api

import (
	"encoding/json"
	"rustdesk-api-server-pro/app/form/api"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/db"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"xorm.io/xorm"
)

type AddressBookPeerController struct {
	basicController
}

func (c *AddressBookPeerController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "ab/peer/add/{guid:string}", "HandleAbPeerAdd")
	b.Handle("PUT", "ab/peer/update/{guid:string}", "HandleAbPeerUpdate")
	b.Handle("DELETE", "ab/peer/{guid:string}", "HandleAbPeerDelete")
}

func (c *AddressBookPeerController) PostAbPeers() mvc.Result {
	current := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("pageSize", 10)
	abGuid := c.Ctx.URLParamDefault("ab", "")

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
	query := func() *xorm.Session {
		q := c.Db.Table(&model.Peer{}).Where("user_id = ? and ab_id = ?", user.Id, ab.Id)
		return q
	}

	pagination := db.NewPagination(current, pageSize)
	peerList := make([]model.Peer, 0)
	err = pagination.Paginate(query, &model.Peer{}, &peerList)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}
	data := make([]iris.Map, 0)
	for _, peer := range peerList {
		forceAlwaysRelay := "false"
		if peer.ForceAlwaysRelay {
			forceAlwaysRelay = "true"
		}

		var peerTags []string
		err := json.Unmarshal([]byte(peer.Tags), &peerTags)
		if err != nil {
			continue
		}

		data = append(data, iris.Map{
			"id":               peer.RustdeskId,
			"hash":             peer.Hash,
			"password":         peer.Password,
			"username":         peer.Username,
			"hostname":         peer.Hostname,
			"platform":         peer.Platform,
			"alias":            peer.Alias,
			"tags":             peerTags,
			"forceAlwaysRelay": forceAlwaysRelay,
			"rdpPort":          peer.RdpPort,
			"rdpUsername":      peer.RdpUsername,
			"loginName":        peer.LoginName,
			"same_server":      peer.SameServer,
		})
	}

	return mvc.Response{
		Object: iris.Map{
			"total": pagination.TotalCount,
			"data":  data,
		},
	}
}

func (c *AddressBookPeerController) HandleAbPeerAdd() mvc.Result {
	abGuid := c.Ctx.Params().Get("guid")

	var form api.AbPeer
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

	totalPeers, err := c.Db.Where("user_id = ? and ab_id = ?", user.Id, ab.Id).Count(&model.Peer{})
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}

	// 限制单个地址簿最大Peer数量
	if ab.MaxPeer > 0 && totalPeers >= int64(ab.MaxPeer) {
		return mvc.Response{
			Object: iris.Map{
				"error": "exceed_max_devices",
			},
		}
	}

	peerTags := ""
	b, err := json.Marshal(form.Tags)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}
	peerTags = string(b)
	if peerTags == "null" {
		peerTags = "[]"
	}
	forceAlwaysRelay := false
	if form.ForceAlwaysRelay == "true" {
		forceAlwaysRelay = true
	}
	sameServer := false
	if form.SameServer != "" {
		sameServer = true
	}
	peer := model.Peer{
		UserId:           user.Id,
		AbId:             ab.Id,
		RustdeskId:       form.Id,
		Username:         form.Username,
		Hostname:         form.Hostname,
		Platform:         form.Platform,
		Alias:            form.Alias,
		Tags:             peerTags,
		ForceAlwaysRelay: forceAlwaysRelay,
		RdpPort:          form.RdpPort,
		RdpUsername:      form.RdpUsername,
		LoginName:        form.LoginName,
		SameServer:       sameServer,
	}
	_, err = c.Db.Insert(&peer)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}
	return mvc.Response{
		Text: "",
	}
}

func (c *AddressBookPeerController) HandleAbPeerUpdate() mvc.Result {
	abGuid := c.Ctx.Params().Get("guid")

	var fields map[string]interface{}
	err := c.Ctx.ReadBody(&fields)
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

	var peer model.Peer
	has, err := c.Db.Where("user_id = ? and ab_id = ? and rustdesk_id = ?", user.Id, ab.Id, fields["id"]).Get(&peer)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}
	if !has {
		return mvc.Response{
			Object: iris.Map{
				"error": "peer not found",
			},
		}
	}

	// 下面的更新逻辑有点复杂，每次请求只会带上要更新的字段，要考虑和其他字段更新的兼容性
	val, ok := fields["tags"]
	if ok {
		peerTags := ""
		b, err := json.Marshal(val)
		if err != nil {
			return mvc.Response{
				Object: iris.Map{
					"error": err.Error(),
				},
			}
		}
		peerTags = string(b)
		if peerTags == "null" {
			peerTags = "[]"
		}
		peer.Tags = peerTags
	}

	val, ok = fields["alias"]
	if ok {
		peer.Alias = val.(string)
	}

	val, ok = fields["hash"]
	if ok {
		peer.Hash = val.(string)
	}

	val, ok = fields["password"]
	if ok {
		peer.Password = val.(string)
	}

	_, err = c.Db.Where("id = ?", peer.Id).Cols("tags", "alias", "hash", "password").Update(&peer)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}
	return mvc.Response{
		Text: "",
	}
}

func (c *AddressBookPeerController) HandleAbPeerDelete() mvc.Result {
	abGuid := c.Ctx.Params().Get("guid")

	var ids []string
	err := c.Ctx.ReadBody(&ids)
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

	c.Db.Where("user_id = ? and ab_id = ?", user.Id, ab.Id).In("rustdesk_id", ids).Delete(&model.Peer{})

	return mvc.Response{}
}
