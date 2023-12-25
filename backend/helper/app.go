package helper

import (
	"github.com/kataras/iris/v12"
	"reflect"
)

func GetAppDependency(app *iris.Application, t string) interface{} {
	for _, d := range app.ConfigureContainer().Container.Dependencies {
		if reflect.TypeOf(d.OriginalValue).String() == t {
			return d.OriginalValue
		}
	}
	return nil
}
