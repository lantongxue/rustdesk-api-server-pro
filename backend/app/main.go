package app

import "github.com/kataras/iris/v12"

func newApp() *iris.Application {
	app := iris.Default()

	return app
}

func StartServer(port string) (bool, error) {

	app := newApp()

	err := app.Listen(port)
	if err != nil {
		return false, err
	}

	return true, nil
}
