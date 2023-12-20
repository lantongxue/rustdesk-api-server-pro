package middleware

import (
	"fmt"
	"github.com/kataras/iris/v12"
)

func RequestLogger() iris.Handler {
	return func(context iris.Context) {
		info := fmt.Sprintf("▶ %s:%s", context.Method(), context.Request().RequestURI)
		body, _ := context.GetBody()
		fmt.Println(info, string(body))
	}
}
