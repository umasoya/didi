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

func main() {
	app := cli.NewApp()
	app.Name = "didi"
	app.Usage = "Destinguish distribution"
	app.Version = "0.9.0"

	// Check if /etc/os-release exists
	app.Before = func(c *cli.Context) error {
		_, err := os.Stat("/etc/os-release")
		return err
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
