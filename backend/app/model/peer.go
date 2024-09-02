package model

import "time"

type Peer struct {
	Id               int       `xorm:"'id' int notnull pk autoincr"`
	UserId           int       `xorm:"'user_id' int"`
	RustdeskId       string    `xorm:"'rustdesk_id' varchar(255)"`
	Hash             string    `xorm:"'hash' varchar(255)"`
	Username         string    `xorm:"'username' varchar(255)"`
	Hostname         string    `xorm:"'hostname' varchar(255)"`
	Platform         string    `xorm:"'platform' varchar(255)"`
	Alias            string    `xorm:"'alias' varchar(255)"`
	Tags             string    `xorm:"'tags' text"`
	Note             string    `xorm:"'note' text"`
	ForceAlwaysRelay bool      `xorm:"'forceAlwaysRelay' tinyint"`
	RdpPort          string    `xorm:"'rdpPort' varchar(5)"`
	RdpUsername      string    `xorm:"'rdpUsername' varchar(100)"`
	LoginName        string    `xorm:"'loginName' varchar(100)"`
	CreatedAt        time.Time `xorm:"'created_at' datetime created"`
	UpdatedAt        time.Time `xorm:"'updated_at' datetime updated"`
}

func (m *Peer) TableName() string {
	return "peer"
}
