package app

import (
	"github.com/kataras/iris/v12"
	"rustdesk-api-server-pro/helper/config"
)

func newApp() *iris.Application {
	app := iris.Default()

	return app
}

func StartServer() (bool, error) {
	cfg := config.GetServerConfig()

	app := newApp()
	err := app.Listen(cfg.Port)
	if err != nil {
		return false, err
	}
	return true, nil
}
