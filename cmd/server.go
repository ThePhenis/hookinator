package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/codegangsta/negroni"
	"github.com/fntlnz/hookinator/client"
	"github.com/fntlnz/hookinator/module/middleware"
	"github.com/fntlnz/hookinator/route"
	"github.com/gorilla/mux"
)

var CmdServer = cli.Command{
	Name:        "server",
	Usage:       "hookinator server",
	Description: "Run the hookinator built-in webserver",
	Action:      runServer,
	Flags:       []cli.Flag{},
}

func runServer(context *cli.Context) {
	r := mux.NewRouter()
	n := negroni.Classic()

	n.Use(middleware.NewStaticLoader(client.Asset))
	n.UseHandler(r)

	r.HandleFunc("/example", route.ExampleHandler)

	n.Run(":3000")
}
