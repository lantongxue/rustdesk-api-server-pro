package app

import (
	"rustdesk-api-server-pro/app/controller/admin"
	"rustdesk-api-server-pro/app/controller/api"
	"rustdesk-api-server-pro/app/middleware"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func SetRoute(app *iris.Application) {
	apiParty := app.Party("/api")
	apiMvc := mvc.New(apiParty)
	apiMvc.Handle(new(api.SystemController))
	apiMvc.Handle(new(api.LoginController))
	apiMvc.Handle(new(api.AuditController))

	apiWithAuthParty := app.Party("/api")
	apiWithAuthParty.Use(middleware.ApiAuth(app))
	{
		apiWithAuthMvc := mvc.New(apiWithAuthParty)
		apiWithAuthMvc.Handle(new(api.UserController))
		apiWithAuthMvc.Handle(new(api.PeerController))
		apiWithAuthMvc.Handle(new(api.AddressBookController))
		apiWithAuthMvc.Handle(new(api.AddressBookPeerController))
		apiWithAuthMvc.Handle(new(api.AddressBookTagController))
	}

	adminParty := app.Party("/admin")
	adminMvc := mvc.New(adminParty)
	adminMvc.Handle(new(admin.AuthController))

	adminWithAuthParty := app.Party("/admin")
	adminWithAuthParty.Use(middleware.AdminAuth(app))
	{
		adminWithAuthMvc := mvc.New(adminWithAuthParty)
		adminWithAuthMvc.Handle(new(admin.IndexController))
		adminWithAuthMvc.Handle(new(admin.DashboardController))
		adminWithAuthMvc.Handle(new(admin.UsersController))
		adminWithAuthMvc.Handle(new(admin.SessionsController))
		adminWithAuthMvc.Handle(new(admin.AuditController))
		adminWithAuthMvc.Handle(new(admin.MailTemplateController))
		adminWithAuthMvc.Handle(new(admin.MaiLogsController))
	}
}
