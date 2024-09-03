package db

import (
	"rustdesk-api-server-pro/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "modernc.org/sqlite"
	"xorm.io/xorm"
)

func NewEngine(cfg *config.DbConfig) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine(cfg.Driver, cfg.Dsn)
	if err != nil {
		return nil, err
	}
	location, _ := time.LoadLocation(cfg.TimeZone)
	engine.TZLocation = location
	engine.DatabaseTZ = location
	engine.ShowSQL(cfg.ShowSql)
	engine.SetMaxIdleConns(100)
	engine.SetMaxOpenConns(100)
	return engine, nil
}
