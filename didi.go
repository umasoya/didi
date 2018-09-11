package main

import (
	"bufio"
	"log"
	"os"

	"github.com/urfave/cli"
)

// Identity is structure of /etc/os-release
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

// Parse is parsed /etc/os-release
func Parse() (Identity, error) {
	identity := Identity{}

	// check os-release file
	f, err := os.Open("/etc/os-release")
	if err != nil {
		return identity, err
	}

	defer f.Close()

	// scan os-release file
	rows := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}
	if serr := scanner.Err(); serr != nil {
		return identity, serr
	}

	return identity, nil
}

// Before is check /etc/os-release exist
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

// App is create new application.
func App() (*cli.App, error) {
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

	_, err := Parse()
	// identity, err := Parse()

	// app.Action = Action

	return app, err
}

func main() {
	app, err := App()
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
