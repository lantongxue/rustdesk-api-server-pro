package app

import (
	"github.com/kataras/iris/v12/mvc"
	"rustdesk-api-server-pro/app/controller/api"
)

func init() {
	apiParty := app.Party("/api")
	apiMvc := mvc.New(apiParty)
	apiMvc.Handle(new(api.SystemController))

	apiMvc.Handle(new(api.LoginController))

	apiMvc.Handle(new(api.UserController))

	apiMvc.Handle(new(api.PeerController))

	apiMvc.Handle(new(api.AddressBookController))

	apiMvc.Handle(new(api.AuditController))
}
