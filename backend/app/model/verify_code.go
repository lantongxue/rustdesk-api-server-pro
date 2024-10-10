package model

import "time"

type VerifyCode struct {
	Id        int       `xorm:"'id' int notnull pk autoincr"`
	UserId    int       `xorm:"'user_id' int"`
	Type      int       `xorm:"'type' tinyint"` // 1=email,2=sms
	Uuid      string    `xorm:"'uuid' varchar(255)"`
	Code      string    `xorm:"'code' varchar(10)"`
	Expired   time.Time `xorm:"'expired' datetime"`
	Status    int       `xorm:"'status' tinyint"` // 1=unused,2=used,3=expired
	CreatedAt time.Time `xorm:"'created_at' datetime created"`
	UpdatedAt time.Time `xorm:"'updated_at' datetime updated"`
}

func (m *VerifyCode) TableName() string {
	return "verify_code"
}
