package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "rpm",
				Usage: "manage rpm repository",
				Subcommands: []*cli.Command{
					{
						Name: "upload",
						Usage: "upload a package",
						Action: func(c *cli.Context) error {
							fmt.Println("Uploading an rpm")
							return nil
						},
					},
				},
			},
			{
				Name: "deb",
				Usage: "manage deb repository",
				Subcommands: []*cli.Command{
					{
						Name: "upload",
						Usage: "upload a package",
						Action: func(c *cli.Context) error {
							fmt.Println("Uploading a deb")
							return nil
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
