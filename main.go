package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "Health Check",
		Version: "1.0.0",
		Usage:   "A tool to check if a website is running or not",

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Aliases:  []string{"d"},
				Usage:    "Domain name to check",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Usage:   "Port number to check",
				Value:   "80", // Default value
			},
		},

		Action: func(c *cli.Context) error {
			domain := c.String("domain")
			port := c.String("port")

			status := Check(domain, port)
			fmt.Println(status)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
