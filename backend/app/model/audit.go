package model

import "time"

type Audit struct {
	Id         int       `xorm:"'id' int notnull pk autoincr"`
	UserId     int       `xorm:"'user_id' int"`
	Action     string    `xorm:"'action' varchar(100)"`
	ConnId     string    `xorm:"'conn_id' varchar(100)"`
	RustdeskId string    `xorm:"'rustdesk_id' varchar(100)"`
	IP         string    `xorm:"'ip' varchar(15)"`
	SessionId  string    `xorm:"'session_id' varchar(100)"`
	Uuid       string    `xorm:"'uuid' varchar(255)"`
	Note       string    `xorm:"'note' varchar(255)"`
	Type       int       `xorm:"'type' int"`
	CreatedAt  time.Time `xorm:"'created_at' datetime created"`
	UpdatedAt  time.Time `xorm:"'updated_at' datetime updated"`
}

func (m *Audit) TableName() string {
	return "audit"
}
