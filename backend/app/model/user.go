package model

import (
	"time"
)

type User struct {
	Id              int       `xorm:"'id' int notnull pk autoincr"`
	Username        string    `xorm:"'username' varchar(50)"`
	Password        string    `xorm:"'password' varchar(255)"`
	Name            string    `xorm:"'name' varchar(100)"`
	Email           string    `xorm:"'email' varchar(255)"`
	Note            string    `xorm:"'note' varchar(255)"`
	LicensedDevices int       `xorm:"'licensed_devices' int"`
	Status          int       `xorm:"'status' tinyint"`
	IsAdmin         bool      `xorm:"'is_admin' tinyint"`
	CreatedAt       time.Time `xorm:"'created_at' datetime created"`
	UpdatedAt       time.Time `xorm:"'updated_at' datetime updated"`
}

func (m *User) TableName() string {
	return "user"
}
