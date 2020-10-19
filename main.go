package main

import (
	"os"

	"github.com/yannuokiss/profile/cmd"

	//_ "github.com/yannuokiss/profile/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
