package cmd

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"rustdesk-api-server-pro/helper/github"
	"rustdesk-api-server-pro/helper/rustdesk"
	"rustdesk-api-server-pro/util"
	"strings"

	"github.com/spf13/cobra"
)

var rustdeskServerCmd = &cobra.Command{
	Use:   "rustdesk [command]",
	Short: "About rustdesk-server command",
}

var rustdeskInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Download and run rustdesk-server",
	Long:  "This command will be download rustdesk-server from https://github.com/rustdesk/rustdesk-server/releases and run it.",
	Run: func(cmd *cobra.Command, args []string) {
		hbbr, hbbs := rustdesk.GetRustdeskServerBin()
		if util.FileExists(hbbr) && util.FileExists(hbbs) {
			fmt.Println("The rustdesk-server has been initialized.")
			os.Exit(0)
		}

		repo := "rustdesk/rustdesk-server"
		release := &github.Release{}
		rustdeskServerVersion := cmd.Flag("version").Value.String()
		if rustdeskServerVersion == "latest" {
			release = github.GetLatestRelease(repo)
		} else {
			release = github.GetReleaseByTag(repo, rustdeskServerVersion)
		}
		matchedAsset := github.Asset{}
		arch := runtime.GOARCH
		for _, asset := range release.Assets {
			if runtime.GOOS == "windows" {
				if strings.Contains(asset.Name, "windows") {
					matchedAsset = asset
					arch = "x86_64"
					break
				}
			}
			if runtime.GOOS == "linux" {
				if asset.Name == "rustdesk-server-linux-amd64.zip" {
					matchedAsset = asset
					arch = "amd64"
					break
				}
			}
		}
		if matchedAsset.Name == "" {
			fmt.Println("Your operating system is not supported, only support windows and linux")
			os.Exit(0)
		}

		// Unzip the rustdesk-server zip if it already exists locally, otherwise download it from github
		if !util.FileExists(matchedAsset.Name) {
			proxyServer := cmd.Flag("proxy").Value.String()
			util.SetHttpProxy(proxyServer)
			err := util.DownloadFile(matchedAsset.BrowserDownloadURL, matchedAsset.Name, true)
			if err != nil {
				fmt.Println("rustdesk-server download error: ", err)
				os.Exit(0)
			}
		}

		fmt.Println("unzipping", matchedAsset.Name)
		err := util.Unzip(matchedAsset.Name, rustdesk.GetRustdeskServerBinDir())
		if err != nil {
			fmt.Println(matchedAsset.Name, "unzipped error: ", err)
			os.Exit(0)
		}
		src := path.Join(rustdesk.GetRustdeskServerBinDir(), arch)
		util.MoveFiles(src, rustdesk.GetRustdeskServerBinDir())
		os.Remove(matchedAsset.Name)
		os.RemoveAll(src)
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

		_, err := rustdesk.StartServer()
		if err != nil {
			fmt.Println("rustdesk-server start failed:", err.Error())
			return
		}
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

var rustdeskListCmd = &cobra.Command{
	Use:                   "list",
	Short:                 "List the releases rustdesk-server",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		proxyServer := cmd.Flag("proxy").Value.String()
		util.SetHttpProxy(proxyServer)
		releases := github.GetReleases("rustdesk/rustdesk-server")
		fmt.Printf("%-20s%s\n", "Version", "Published")
		for _, release := range *releases {
			fmt.Printf("%-20s%s\n", release.TagName, release.PublishedAt)
		}
	},
}

func init() {
	rustdeskInstallCmd.Flags().StringP("proxy", "p", "", "Setting up a proxy to download rustdesk-server program (e.g [http|https|socks5]://proxy-host:port)")
	rustdeskInstallCmd.Flags().StringP("version", "v", "latest", "Setting the rustdesk-server program version")
	rustdeskServerCmd.AddCommand(rustdeskInstallCmd)
	rustdeskServerCmd.AddCommand(rustdeskStartCmd)
	rustdeskServerCmd.AddCommand(rustdeskStopCmd)
	rustdeskServerCmd.AddCommand(rustdeskRestartCmd)
	rustdeskServerCmd.AddCommand(rustdeskStatusCmd)
	rustdeskServerCmd.AddCommand(rustdeskKeysCmd)

	rustdeskListCmd.Flags().StringP("proxy", "p", "", "Setting up a proxy to download rustdesk-server program (e.g [http|https|socks5]://proxy-host:port)")
	rustdeskServerCmd.AddCommand(rustdeskListCmd)
	RootCmd.AddCommand(rustdeskServerCmd)
}
