package model

import "time"

type Device struct {
	Id         int       `xorm:"'id' int notnull pk autoincr"`
	UserId     int       `xorm:"'user_id' int"`
	Cpu        string    `xorm:"'cpu' varchar(255)"`
	Hostname   string    `xorm:"'hostname' varchar(255)"`
	RustdeskId string    `xorm:"'rustdesk_id' varchar(255)"`
	Memory     string    `xorm:"'memory' varchar(50)"`
	Os         string    `xorm:"'os' varchar(255)"`
	Username   string    `xorm:"'username' varchar(255)"`
	Uuid       string    `xorm:"'uuid' varchar(255)"`
	Version    string    `xorm:"'version' varchar(255)"`
	CreatedAt  time.Time `xorm:"'created_at' datetime created"`
	UpdatedAt  time.Time `xorm:"'updated_at' datetime updated"`
}

func (m *Device) TableName() string {
	return "device"
}
