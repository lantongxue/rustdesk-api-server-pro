package main

import (
	"rustdesk-api-server-pro/cmd"
)

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
