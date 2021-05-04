package main

import (
	"log"
	"os"

	"github.com/tatsuki9/eggsm/internal"
	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "secret_id",
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
