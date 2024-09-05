package model

import "time"

const PersonalAddressBookName = "My address book"
const MaxPeer = 0
const LegacyAddressBookName = "Legacy address book"

type AddressBook struct {
	Id        int       `xorm:"'id' int notnull pk autoincr"`
	UserId    int       `xorm:"'user_id' int"`
	Guid      string    `xorm:"'guid' varchar(64)"`
	Name      string    `xorm:"'name' varchar(255)"`
	Owner     string    `xorm:"'owner' varchar(255)"` // 归属人：用户名
	Note      string    `xorm:"'note' varchar(255)"`
	Rule      int       `xorm:"'rule' int"`       // 1 read 2 read&write 3 fullcontrol
	MaxPeer   int       `xorm:"'max_peer' int"`   // 一个ab最大能存peer数量
	Shared    bool      `xorm:"'shared' tinyint"` // 设置是否共享地址簿
	CreatedAt time.Time `xorm:"'created_at' datetime created"`
	UpdatedAt time.Time `xorm:"'updated_at' datetime updated"`
}

func (m *AddressBook) TableName() string {
	return "address_book"
}

type AddressBookTag struct {
	Id        int       `xorm:"'id' int notnull pk autoincr"`
	UserId    int       `xorm:"'user_id' int"`
	AbId      int       `xorm:"'ab_id' int"`
	Name      string    `xorm:"'name' varchar(255)"`
	Color     int64     `xorm:"'color' int"`
	CreatedAt time.Time `xorm:"'created_at' datetime created"`
	UpdatedAt time.Time `xorm:"'updated_at' datetime updated"`
}

func (m *AddressBookTag) TableName() string {
	return "address_book_tag"
}
