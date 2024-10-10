package model

import "time"

type SystemSettings struct {
	Id        int       `xorm:"'id' int notnull pk autoincr"`
	Name      string    `xorm:"'name' varchar(255)"`
	Key       string    `xorm:"'key' varchar(255)"`
	Value     string    `xorm:"'value' varchar(255)"`
	CreatedAt time.Time `xorm:"'created_at' datetime created"`
	UpdatedAt time.Time `xorm:"'updated_at' datetime updated"`
}

func (m *SystemSettings) TableName() string {
	return "system_settings"
}
