package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/helper"
	"time"
	"xorm.io/xorm"
)

func ApiAuth(app *iris.Application) iris.Handler {
	return func(context iris.Context) {
		db := helper.GetAppDependency(app, "*xorm.Engine").(*xorm.Engine)
		token := jwt.FromHeader(context)

		var authToken model.AuthToken
		get, err := db.Where("token = ? and expired > ? and status = 1", token, time.Now().Format(config.TimeFormat)).Get(&authToken)
		if !get || err != nil {
			context.StopWithText(iris.StatusUnauthorized, "Unauthorized")
			return
		}

		var user model.User
		get, err = db.Where("id = ? and status > 0", authToken.UserId).Get(&user)
		if !get || err != nil {
			context.StopWithText(iris.StatusNotAcceptable, "NotAcceptable")
			return
		}

		context.Values().Set(config.CurrentUserKey, &user)
		context.Values().Set(config.CurrentAuthTokenString, token)
		context.Values().Set(config.CurrentAuthToken, &authToken)
		context.Next()
	}
}
