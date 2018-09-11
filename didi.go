package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

type Identity struct {
	Name             string
	Version          string
	ID               string
	IDLike           string
	VersionCodename  string
	VersionID        string
	PrettyName       string
	AnsiColor        string
	CpeName          string
	HomeURL          string
	SupportURL       string
	BugReportURL     string
	PrivatePolicyURL string
	BuildID          string
	Variant          string
	VariantID        string
}

func Before(c *cli.Context) error {
	_, err := os.Stat("/etc/os-release")
	return err
}

/*
func Action(c *cli.Context) error {
	if c.Bool("all") {
		fmt.Println("all")
	} else {
		fmt.Println("none")
	}
	return nil
}
*/

// Create new application.
func App() *cli.App {
	app := cli.NewApp()
	app.Name = "didi"
	app.Usage = "Destinguish distribution"
	app.Version = "0.9.0"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "all, a",
			Usage: "print all information",
		},
	}

	app.Before = Before
	// app.Action = Action

	return app
}

func main() {
	app := App()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
