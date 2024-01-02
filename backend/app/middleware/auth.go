package middleware

import (
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/helper"
	"time"

	"github.com/golang-module/carbon/v2"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	"xorm.io/xorm"
)

func ApiAuth(app *iris.Application) iris.Handler {
	return func(context iris.Context) {
		db := helper.GetAppDependency(app, "*xorm.Engine").(*xorm.Engine)
		token := jwt.FromHeader(context)

		var authToken model.AuthToken
		get, err := db.Where("token = ? and expired > ? and status = 1 and is_admin = 0", token, time.Now().Format(config.TimeFormat)).Get(&authToken)
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

func AdminAuth(app *iris.Application) iris.Handler {
	return func(context iris.Context) {
		db := helper.GetAppDependency(app, "*xorm.Engine").(*xorm.Engine)
		token := context.GetHeader("Authorization")

		var authToken model.AuthToken
		get, err := db.Where("token = ? and expired > ? and status = 1 and is_admin = 1", token, time.Now().Format(config.TimeFormat)).Get(&authToken)
		if !get || err != nil {
			context.StopWithText(iris.StatusUnauthorized, "Unauthorized")
			return
		}

		var user model.User
		get, err = db.Where("id = ? and status > 0 and is_admin = 1", authToken.UserId).Get(&user)
		if !get || err != nil {
			context.StopWithText(iris.StatusNotAcceptable, "NotAcceptable")
			return
		}

		s := carbon.Now().DiffInSeconds(carbon.Parse(authToken.Expired.Format(config.TimeFormat)))
		if s <= 60*5 {
			authToken.Expired = carbon.Parse(authToken.Expired.Format(config.TimeFormat)).AddHours(2).ToStdTime()
			db.Where("id = ?", authToken.Id).Cols("expired").Update(&authToken)
		}

		context.Values().Set(config.AdminUserKey, &user)
		context.Values().Set(config.AdminAuthTokenString, token)
		context.Values().Set(config.AdminAuthToken, &authToken)
		context.Next()
	}
}
