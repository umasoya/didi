package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "didi"
	app.Usage = "Destinguish distribution"
	app.Version = "0.9.0"

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
