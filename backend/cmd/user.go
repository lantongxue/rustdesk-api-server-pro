package cmd

import (
	"fmt"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"
	"rustdesk-api-server-pro/util"

	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "User management",
}

var isAdmin bool

var userAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add user [ add username password][--admin]",
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		password, _ := util.Password(args[1])

		user := &model.User{
			Username:        username,
			Password:        password,
			Name:            username,
			LicensedDevices: 0,
			LoginVerify:     model.LOGIN_ACCESS_TOKEN,
			IsAdmin:         isAdmin,
			Status:          1,
		}
		cfg := config.GetServerConfig()
		engine, err := db.NewEngine(cfg.Db)
		if err != nil {
			fmt.Println("Db Engine create error:", err)
			return
		}
		_, err = engine.Insert(user)
		if err != nil {
			fmt.Println("Add error:", err)
			return
		}
		fmt.Println("Add success")
	},
}

func init() {
	userAddCmd.Flags().BoolVarP(&isAdmin, "admin", "a", false, "Set user admin")
	userCmd.AddCommand(userAddCmd)
	RootCmd.AddCommand(userCmd)
}
