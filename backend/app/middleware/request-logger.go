package middleware

import (
	"fmt"
	"github.com/kataras/iris/v12"
)

func RequestLogger() iris.Handler {
	return func(context iris.Context) {
		requestInfo := fmt.Sprintf("▶ %s:%s", context.Method(), context.Request().RequestURI)
		body, _ := context.GetBody()
		context.Application().Logger().Info(requestInfo)
		for header, value := range context.Request().Header {
			fmt.Println(header+":", value)
		}
		fmt.Println(string(body))
		context.Next()
	}
}
