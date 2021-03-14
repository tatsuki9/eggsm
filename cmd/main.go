package main

import (
	"log"
	"os"

	"eggsm/internal"
	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "profile",
				Usage: "Enter the characters",
			},
			&cli.StringFlag{
				Name:  "prefix",
				Usage: "Enter the characters",
			},
			&cli.StringFlag{
				Name:  "env",
				Usage: "Enter the characters",
			},
		},
		Action: func(c *cli.Context) error {
			sv, err := internal.GetSecretValues(c)
			if err != nil {
				return err
			}
			internal.Output(sv)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Println(err)
	}
}
