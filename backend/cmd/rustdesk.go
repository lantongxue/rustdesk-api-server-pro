package cmd

import (
	"fmt"
	"os"
	"runtime"
	"rustdesk-api-server-pro/helper/rustdesk"
	"rustdesk-api-server-pro/util"

	"github.com/spf13/cobra"
)

var rustdeskServerCmd = &cobra.Command{
	Use:   "rustdesk [command]",
	Short: "About rustdesk-server command",
}

var rustdeskInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Download and run rustdesk-server",
	Long:  "This command will be download rustdesk-server from https://github.com/rustdesk/rustdesk-server and run it.",
	Run: func(cmd *cobra.Command, args []string) {
		hbbr, hbbs := rustdesk.GetRustdeskServerBin()
		if util.FileExists(hbbr) && util.FileExists(hbbs) {
			fmt.Println("The rustdesk-server has been initialized.")
			os.Exit(0)
		}

		var serverProgramZip string = ""
		switch runtime.GOOS {
		case "windows":
			serverProgramZip = "rustdesk-server-windows-x86_64.zip"
		case "linux":
			serverProgramZip = "rustdesk-server-linux-amd64.zip"
		default:
			fmt.Println("Your operating system is not supported")
			os.Exit(0)
		}

		// Unzip the rustdesk-server zip if it already exists locally, otherwise download it from github
		if !util.FileExists(serverProgramZip) {
			rustdeskServerVersion := cmd.Flag("version").Value.String()
			serverProgramUrl := fmt.Sprintf("https://github.com/rustdesk/rustdesk-server/releases/download/%s/", rustdeskServerVersion)

			serverProgramUrl += serverProgramZip
			proxyServer := cmd.Flag("proxy").Value.String()
			err := util.DownloadFile(serverProgramUrl, serverProgramZip, proxyServer, true)
			if err != nil {
				fmt.Println("rustdesk-server download error: ", err)
				os.Exit(0)
			}
		}

		fmt.Println("unzipping", serverProgramZip)
		err := util.Unzip(serverProgramZip, "rustdesk-server")
		if err != nil {
			fmt.Println(serverProgramZip, "unzipped error: ", err)
			os.Exit(0)
		}
		fmt.Println("The rustdesk-server has been initialized.")
	},
}

var rustdeskStartCmd = &cobra.Command{
	Use:                   "start",
	Short:                 "Start the rustdesk-server",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		ret, err := rustdesk.StartServer()
		if ret {
			fmt.Println("rustdesk-server started")
		} else {
			fmt.Println("rustdesk-server failed to start:", err.Error())
		}
	},
}

var rustdeskStopCmd = &cobra.Command{
	Use:                   "stop",
	Short:                 "Stop the rustdesk-server",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		rustdesk.StopServer()
		fmt.Println("rustdesk-server stopped")
	},
}

var rustdeskRestartCmd = &cobra.Command{
	Use:                   "restart",
	Short:                 "Restart the rustdesk-server",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		rustdesk.StopServer()
		fmt.Println("rustdesk-server stopped")

		rustdesk.StartServer()
		fmt.Println("rustdesk-server started")
	},
}

var rustdeskStatusCmd = &cobra.Command{
	Use:                   "status",
	Short:                 "Show rustdesk-server status",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		hbbrIsRunning, hbbsIsRunning := rustdesk.Status()

		fmt.Println("ServerName\t", "IsRunning")
		fmt.Println("hbbr\t\t", hbbrIsRunning)
		fmt.Println("hbbs\t\t", hbbsIsRunning)
	},
}

var rustdeskKeysCmd = &cobra.Command{
	Use:                   "keys",
	Short:                 "Show rustdesk-server keys",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		public, private := rustdesk.Keys()

		fmt.Println("public keys:")
		fmt.Println(public)

		fmt.Println("")

		fmt.Println("private keys:")
		fmt.Println(private)
	},
}

func init() {
	rustdeskInitCmd.Flags().StringP("proxy", "p", "", "Setting up a proxy to download rustdesk-server program (e.g [http|https|socks5]://proxy-host:port)")
	rustdeskInitCmd.Flags().StringP("version", "v", "1.1.9", "Setting the rustdesk-server program version")
	rustdeskServerCmd.AddCommand(rustdeskInitCmd)
	rustdeskServerCmd.AddCommand(rustdeskStartCmd)
	rustdeskServerCmd.AddCommand(rustdeskStopCmd)
	rustdeskServerCmd.AddCommand(rustdeskRestartCmd)
	rustdeskServerCmd.AddCommand(rustdeskStatusCmd)
	rustdeskServerCmd.AddCommand(rustdeskKeysCmd)
	RootCmd.AddCommand(rustdeskServerCmd)
}
