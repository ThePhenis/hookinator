package main

import (
	"github.com/codegangsta/cli"
	"github.com/fntlnz/hookinator/cmd"
	"os"
)

const VER = "0.0.1-alpha"

func main() {
	app := cli.NewApp()
	app.Name = "hookinator"
	app.Usage = "We have to write a funny entrance here"
	app.Commands = []cli.Command{
		cmd.CmdServer,
	}
	app.Run(os.Args)
}
