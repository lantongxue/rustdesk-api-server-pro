package model

import "time"

type AuthToken struct {
	Id        int       `xorm:"'id' int notnull pk autoincr"`
	UserId    int       `xorm:"'user_id' int"`
	MyId      string    `xorm:"'my_id' varchar(255)"`
	Uuid      string    `xorm:"'uuid' varchar(255)"`
	Token     string    `xorm:"'token' varchar(255)"`
	Expired   time.Time `xorm:"'expired' datetime"`
	Status    int       `xorm:"'status' tinyint"`
	CreatedAt time.Time `xorm:"'created_at' datetime created"`
	UpdatedAt time.Time `xorm:"'updated_at' datetime updated"`
}

func (m *AuthToken) TableName() string {
	return "auth_token"
}
