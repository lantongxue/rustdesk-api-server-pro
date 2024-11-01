package model

import "time"

type MailTemplate struct {
	Id        int       `xorm:"'id' int notnull pk autoincr"`
	Name      string    `xorm:"'name' varchar(255)"`
	Type      int       `xorm:"'type' tinyint"` // login_code
	Subject   string    `xorm:"'subject' varchar(255)"`
	Contents  string    `xorm:"'contents' mediumtext"`
	CreatedAt time.Time `xorm:"'created_at' datetime created"`
	UpdatedAt time.Time `xorm:"'updated_at' datetime updated"`
}

func (m *MailTemplate) TableName() string {
	return "mail_template"
}
