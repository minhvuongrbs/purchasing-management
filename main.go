package main

import (
	"os"
	"purchasing-management/cmd"

	"github.com/urfave/cli/v2"
)

func main() {
	appCli := appCommandLineInterface()
	if err := appCli.Run(os.Args); err != nil {
		panic(err)
	}
}

func appCommandLineInterface() *cli.App {
	appCli := cli.NewApp()
	appCli.Action = cmd.ServerRun
	appCli.Commands = []*cli.Command{
		{
			Name:   "server",
			Usage:  "start grpc/http server",
			Action: cmd.ServerRun,
		},
	}
	return appCli
}
