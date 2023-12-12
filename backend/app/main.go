package app

import "github.com/kataras/iris/v12"

func newApp() *iris.Application {
	app := iris.Default()

	return app
}

func StartServer() (bool, error) {

	return true, nil
}
