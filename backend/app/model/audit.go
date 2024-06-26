package model

import "time"

type Audit struct {
	Id        int       `xorm:"'id' int notnull pk autoincr"`
	UserId    int       `xorm:"'user_id' int"`
	Action    string    `xorm:"'action' varchar(100)"`
	ConnId    string    `xorm:"'conn_id' varchar(100)"`
	MyId      string    `xorm:"'my_id' varchar(100)"`
	IP        string    `xorm:"'ip' varchar(15)"`
	SessionId string    `xorm:"'session_id' varchar(100)"`
	Uuid      string    `xorm:"'uuid' varchar(255)"`
	CreatedAt time.Time `xorm:"'created_at' datetime created"`
	UpdatedAt time.Time `xorm:"'updated_at' datetime updated"`
}

func (m *Audit) TableName() string {
	return "audit"
}
