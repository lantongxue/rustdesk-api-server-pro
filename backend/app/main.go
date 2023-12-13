package app

import (
	"github.com/kataras/iris/v12"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"
)

func newApp(cfg *config.ServerConfig) (*iris.Application, error) {
	app := iris.Default()

	dbEngine, err := db.NewEngine(cfg.Db)
	if err != nil {
		app.Logger().Fatal("Db Engine create error:", err)
		return nil, err
	}
	app.RegisterDependency(dbEngine)

	return app, nil
}

func StartServer() (bool, error) {
	cfg := config.GetServerConfig()

	app, err := newApp(cfg)
	if err != nil {
		return false, err
	}

	err = app.Listen(cfg.Port)
	if err != nil {
		return false, err
	}
	return true, nil
}
