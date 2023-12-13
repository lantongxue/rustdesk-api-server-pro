package model

import "time"

type Peer struct {
	Id        int       `xorm:"'id' int notnull pk autoincr"`
	UserId    int       `xorm:"'user_id' int"`
	PeerId    string    `xorm:"'peer_id' varchar(255)"`
	Hash      string    `xorm:"'hash' varchar(255)"`
	Username  string    `xorm:"'username' varchar(255)"`
	Hostname  string    `xorm:"'hostname' varchar(255)"`
	Platform  string    `xorm:"'platform' varchar(255)"`
	Alias     string    `xorm:"'alias' varchar(255)"`
	Tags      string    `xorm:"'tags' text"`
	CreatedAt time.Time `xorm:"'created_at' datetime created"`
	UpdatedAt time.Time `xorm:"'updated_at' datetime updated"`
}

func (m *Peer) TableName() string {
	return "peer"
}
