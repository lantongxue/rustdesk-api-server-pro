package app

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"rustdesk-api-server-pro/app/middleware"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"
	"time"
)

func newApp(cfg *config.ServerConfig) (*iris.Application, error) {
	app := iris.Default()

	location, _ := time.LoadLocation(cfg.TimeZone)

	dbEngine, err := db.NewEngine(cfg.Db)
	if err != nil {
		app.Logger().Fatal("Db Engine create error:", err)
		return nil, err
	}
	dbEngine.TZLocation = location
	dbEngine.DatabaseTZ = location
	app.RegisterDependency(dbEngine, cfg)

	app.OnErrorCode(iris.StatusNotFound, func(context iris.Context) {
		requestInfo := fmt.Sprintf("(404)▶ %s:%s", context.Method(), context.Request().RequestURI)
		body, _ := context.GetBody()
		context.Application().Logger().Info(requestInfo)
		for header, value := range context.Request().Header {
			fmt.Println(header+":", value)
		}
		fmt.Println(string(body))
	})

	app.Use(iris.Compression)
	app.Use(middleware.RequestLogger())

	SetRoute(app)

	return app, nil
}

func StartServer() (bool, error) {
	cfg := config.GetServerConfig()

	app, err := newApp(cfg)
	if err != nil {
		return false, err
	}

	err = app.Listen(cfg.Port, iris.WithoutBodyConsumptionOnUnmarshal)
	if err != nil {
		return false, err
	}
	return true, nil
}
