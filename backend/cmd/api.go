package cmd

import (
	"fmt"
	"rustdesk-api-server-pro/app"

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
		port := cmd.Flag("port").Value.String()
		ret, err := app.StartServer(port)
		if ret {
			fmt.Println("api-server started")
		} else {
			fmt.Println("api-server failed to start:", err.Error())
		}
	},
}

func init() {
	apiStartCmd.Flags().StringP("port", "p", ":8080", "Setting the listen port")
	apiCmd.AddCommand(apiStartCmd)
	RootCmd.AddCommand(apiCmd)
}
