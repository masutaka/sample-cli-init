package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "sample-cli-init"
	app.Version = Version
	app.Usage = ""
	app.Author = ""
	app.Email = ""
	app.Commands = Commands

	app.Run(os.Args)
}
