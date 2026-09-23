package db

import (
	"rustdesk-api-server-pro/config"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "modernc.org/sqlite"
	"xorm.io/xorm"
)

var (
	DbEngine *xorm.Engine
	engineMu sync.Mutex
)

func NewEngine(cfg *config.DbConfig) (*xorm.Engine, error) {
	if DbEngine != nil {
		return DbEngine, nil
	}

	engineMu.Lock()
	defer engineMu.Unlock()

	if DbEngine != nil {
		return DbEngine, nil
	}

	engine, err := xorm.NewEngine(cfg.Driver, cfg.Dsn)
	if err != nil {
		return nil, err
	}
	location, _ := time.LoadLocation(cfg.TimeZone)
	engine.TZLocation = location
	engine.DatabaseTZ = location
	engine.ShowSQL(cfg.ShowSql)
	engine.SetMaxIdleConns(25)
	engine.SetMaxOpenConns(25)
	DbEngine = engine
	return engine, nil
}
