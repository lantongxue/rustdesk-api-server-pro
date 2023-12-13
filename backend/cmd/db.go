package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"
)

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "About api-server db command",
}

var dbInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the database",
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
	dbCmd.AddCommand(dbInitCmd)
	RootCmd.AddCommand(dbCmd)
}
