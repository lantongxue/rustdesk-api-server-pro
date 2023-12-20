package db

import (
	_ "github.com/go-sql-driver/mysql"
	_ "modernc.org/sqlite"
	"rustdesk-api-server-pro/config"
	"xorm.io/xorm"
)

func NewEngine(cfg *config.DbConfig) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine(cfg.Driver, cfg.Dsn)
	if err != nil {
		return nil, err
	}
	engine.ShowSQL(cfg.ShowSql)
	engine.SetMaxIdleConns(100)
	engine.SetMaxOpenConns(100)
	return engine, nil
}
