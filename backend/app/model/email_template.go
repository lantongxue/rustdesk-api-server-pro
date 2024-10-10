package model

import "time"

const (
	TPL_TYPE_LOGIN_CHECK = "login_check"
)

type EmailTemplate struct {
	Id        int       `xorm:"'id' int notnull pk autoincr"`
	Name      string    `xorm:"'name' varchar(255)"`
	Type      int       `xorm:"'type' varchar(20)"` // login_code
	Subject   string    `xorm:"'subject' varchar(255)"`
	Contents  string    `xorm:"'contents' mediumtext"`
	CreatedAt time.Time `xorm:"'created_at' datetime created"`
	UpdatedAt time.Time `xorm:"'updated_at' datetime updated"`
}

func (m *EmailTemplate) TableName() string {
	return "email_template"
}
