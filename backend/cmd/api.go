package cmd

import (
	"fmt"
	"rustdesk-server-pro/app"

	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "About api-server command",
}

var apiStartCmd = &cobra.Command{
	Use:                   "start",
	Short:                 "Start the api-server",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		ret, err := app.StartServer()
		if ret {
			fmt.Println("api-server started")
		} else {
			fmt.Println("api-server failed to start:", err.Error())
		}
	},
}

func init() {
	apiCmd.AddCommand(apiStartCmd)
	RootCmd.AddCommand(apiCmd)
}
