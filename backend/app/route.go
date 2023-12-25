package app

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"rustdesk-api-server-pro/app/controller/api"
	"rustdesk-api-server-pro/app/middleware"
)

func SetRoute(app *iris.Application) {
	apiParty := app.Party("/api")
	apiMvc := mvc.New(apiParty)
	apiMvc.Handle(new(api.SystemController))
	apiMvc.Handle(new(api.LoginController))

	apiWithAuthParty := app.Party("/api")
	apiWithAuthParty.Use(middleware.Auth(app))
	{
		apiWithAuthMvc := mvc.New(apiWithAuthParty)
		apiWithAuthMvc.Handle(new(api.UserController))
		apiWithAuthMvc.Handle(new(api.PeerController))
		apiWithAuthMvc.Handle(new(api.AddressBookController))
		apiWithAuthMvc.Handle(new(api.AuditController))
	}
}
