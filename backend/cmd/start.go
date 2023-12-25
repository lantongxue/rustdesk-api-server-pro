package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"rustdesk-api-server-pro/app"
)

var startCmd = &cobra.Command{
	Use:                   "start",
	Short:                 "Start the api-server",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := app.StartServer()
		if err != nil {
			fmt.Println("api-server failed to start:", err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(startCmd)
}
