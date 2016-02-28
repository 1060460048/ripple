package main

import (
	"github.com/codegangsta/cli"
	"os"
	"github.com/bmbstack/ripple/fixtures/forum/scripts"
)

func main() {
	app := cli.NewApp()
	app.Name = "forum"
	app.Usage = "A forum application powered by Ripple framework"
	app.Author = "wangmingjob"
	app.Email = "wangmingjob@icloud.com"
	app.Version = "0.0.1"
	app.Commands = scripts.Commands()
	app.Run(os.Args)
}
