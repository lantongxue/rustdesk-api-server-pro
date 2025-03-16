package cmd

import (
	"fmt"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"

	"github.com/spf13/cobra"
)

var dbSyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "The api-server database synchronization",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetServerConfig()
		engine, err := db.NewEngine(cfg.Db)
		if err != nil {
			fmt.Println("Db Engine create error:", err)
			return
		}
		models := []interface{}{
			new(model.User),
			new(model.Peer),
			new(model.Tags),
			new(model.AuthToken),
			new(model.Audit),
			new(model.FileTransfer),
			new(model.Device),
			new(model.AddressBook),
			new(model.AddressBookTag),
			new(model.MailLogs),
			new(model.VerifyCode),
			new(model.SystemSettings),
			new(model.MailTemplate),
		}
		err = engine.Sync(models...)
		if err != nil {
			fmt.Println("Db init error:", err)
			return
		}
		fmt.Println("Database tables sync success")
	},
}

func init() {
	RootCmd.AddCommand(dbSyncCmd)
}
