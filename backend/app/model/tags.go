package model

import "time"

type Tags struct {
	Id        int       `xorm:"'id' int notnull pk autoincr"`
	UserId    int       `xorm:"'user_id' int"`
	Tag       string    `xorm:"'tag' varchar(255)"`
	Color     string    `xorm:"'color' varchar(255)"`
	CreatedAt time.Time `xorm:"'created_at' datetime created"`
	UpdatedAt time.Time `xorm:"'updated_at' datetime updated"`
}

func (m *Tags) TableName() string {
	return "tag"
}
