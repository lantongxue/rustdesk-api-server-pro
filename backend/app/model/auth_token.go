package model

import "time"

type AuthToken struct {
	Id         int       `xorm:"'id' int notnull pk autoincr"`
	UserId     int       `xorm:"'user_id' int"`
	RustdeskId string    `xorm:"'rustdesk_id' varchar(255)"`
	Uuid       string    `xorm:"'uuid' varchar(255)"`
	DeviceOs   string    `xorm:"'device_os' varchar(10)"`
	DeviceType string    `xorm:"'device_type' varchar(10)"`
	DeviceName string    `xorm:"'device_name' varchar(255)"`
	Token      string    `xorm:"'token' varchar(255)"`
	Expired    time.Time `xorm:"'expired' datetime"`
	IsAdmin    bool      `xorm:"'is_admin' tinyint"`
	Status     int       `xorm:"'status' tinyint"`
	CreatedAt  time.Time `xorm:"'created_at' datetime created"`
	UpdatedAt  time.Time `xorm:"'updated_at' datetime updated"`
}

func (m *AuthToken) TableName() string {
	return "auth_token"
}
