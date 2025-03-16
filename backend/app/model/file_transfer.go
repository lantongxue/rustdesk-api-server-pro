package model

import "time"

type FileTransfer struct {
	Id         int       `xorm:"'id' int notnull pk autoincr"`
	RustdeskId string    `xorm:"'rustdesk_id' varchar(100)"`
	Info       string    `xorm:"'info' text"`
	IsFile     bool      `xorm:"'is_file' tinyint"`
	Path       string    `xorm:"'path' text"`
	PeerId     string    `xorm:"'peer_id' varchar(100)"`
	Type       int       `xorm:"'type' tinyint"`
	Uuid       string    `xorm:"'uuid' varchar(255)"`
	CreatedAt  time.Time `xorm:"'created_at' datetime created"`
	UpdatedAt  time.Time `xorm:"'updated_at' datetime updated"`
}

func (m *FileTransfer) TableName() string {
	return "file_transfer"
}
