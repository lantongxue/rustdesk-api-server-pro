package model

import "time"

type Audit struct {
	Id         int       `xorm:"'id' int notnull pk autoincr"`
	Action     string    `xorm:"'action' varchar(100)"`
	ConnId     int       `xorm:"'conn_id' int"`
	RustdeskId string    `xorm:"'rustdesk_id' varchar(100)"`
	IP         string    `xorm:"'ip' varchar(15)"`
	SessionId  string    `xorm:"'session_id' varchar(50)"`
	Peer       string    `xorm:"'peer' text"`
	Uuid       string    `xorm:"'uuid' varchar(255)"`
	Note       string    `xorm:"'note' varchar(255)"`
	Type       int       `xorm:"'type' tinyint"`
	CreatedAt  time.Time `xorm:"'created_at' datetime created"`
	UpdatedAt  time.Time `xorm:"'updated_at' datetime updated"`
}

func (m *Audit) TableName() string {
	return "audit"
}
