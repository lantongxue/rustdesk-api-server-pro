package api

import (
	"encoding/json"
	"rustdesk-api-server-pro/app/form/api"
	"rustdesk-api-server-pro/app/model"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type AddressBookController struct {
	basicController
}

func (c *AddressBookController) GetAb() mvc.Result {
	user := c.GetUser()
	tagList := make([]model.Tags, 0)
	err := c.Db.Where("user_id = ?", user.Id).Find(&tagList)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}
	tags := make([]string, 0)
	tagColors := make(map[string]int64)
	for _, tag := range tagList {
		tags = append(tags, tag.Tag)
		colorCode, err := strconv.ParseInt(tag.Color, 10, 64)
		if err != nil {
			continue
		}
		tagColors[tag.Tag] = colorCode
	}

	peerList := make([]model.Peer, 0)
	err = c.Db.Where("user_id = ?", user.Id).Find(&peerList)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}
	peers := make([]iris.Map, 0)
	for _, peer := range peerList {
		var peerTags []string
		err := json.Unmarshal([]byte(peer.Tags), &peerTags)
		if err != nil {
			continue
		}
		peers = append(peers, iris.Map{
			"id":       peer.RustdeskId,
			"hash":     peer.Hash,
			"username": peer.Username,
			"hostname": peer.Hostname,
			"platform": peer.Platform,
			"alias":    peer.Alias,
			"tags":     peerTags,
		})
	}

	tagColorsJson, err := json.Marshal(tagColors)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}

	dataJson, err := json.Marshal(iris.Map{
		"tags":       tags,
		"peers":      peers,
		"tag_colors": string(tagColorsJson),
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
			"licensed_devices": user.LicensedDevices,
			"data":             string(dataJson),
		},
	}
}

func (c *AddressBookController) PostAb() mvc.Result {
	var abForm api.AbForm
	err := c.Ctx.ReadJSON(&abForm)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}
	var abData api.AbData
	err = json.Unmarshal([]byte(abForm.Data), &abData)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}
	var tagColors map[string]int64
	err = json.Unmarshal([]byte(abData.TagColors), &tagColors)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}

	session := c.Db.NewSession()
	defer session.Close()
	err = session.Begin()
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}
	user := c.GetUser()
	_, err = session.Where("user_id = ?", user.Id).Delete(&model.Tags{})
	if err != nil {
		_ = session.Rollback()
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}

	_, err = session.Where("user_id = ?", user.Id).Delete(&model.Peer{})
	if err != nil {
		_ = session.Rollback()
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}

	tags := make([]*model.Tags, 0)
	for _, tag := range abData.Tags {
		tags = append(tags, &model.Tags{
			UserId: user.Id,
			Tag:    tag,
			Color:  strconv.FormatInt(tagColors[tag], 10),
		})
	}
	if len(tags) > 0 {
		_, err = session.Insert(tags)
		if err != nil {
			_ = session.Rollback()
			return mvc.Response{
				Object: iris.Map{
					"error": err.Error(),
				},
			}
		}
	}

	peers := make([]*model.Peer, 0)
	for _, peer := range abData.Peers {
		peerTags := ""
		b, err := json.Marshal(peer.Tags)
		if err == nil {
			peerTags = string(b)
		}
		peers = append(peers, &model.Peer{
			UserId:     user.Id,
			RustdeskId: peer.Id,
			Hash:       peer.Hash,
			Username:   peer.Username,
			Hostname:   peer.Hostname,
			Platform:   peer.Platform,
			Alias:      peer.Alias,
			Tags:       peerTags,
		})
	}
	if len(peers) > 0 {
		_, err = session.Insert(peers)
		if err != nil {
			_ = session.Rollback()
			return mvc.Response{
				Object: iris.Map{
					"error": err.Error(),
				},
			}
		}
	}

	err = session.Commit()
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}

	return mvc.Response{}
}
