package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "db",
				Value: "judge.db",
				Usage: "database path",
			},
		},
		Commands: []*cli.Command{{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "start web server",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "listen",
					Value: "[::]:8080",
					Usage: "listen host",
				},
				&cli.StringFlag{
					Name:  "static",
					Value: "web",
					Usage: "static file path",
				},
			},
			Action: func(cCtx *cli.Context) error {
				s := server{
					Listen:         cCtx.String("listen"),
					Database:       cCtx.String("db"),
					StaticFilePath: cCtx.String("static"),
				}
				s.init()
				s.serve()
				return nil
			},
		}, {
			Name:    "importPlayer",
			Aliases: []string{"ip"},
			Usage:   "import players",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "data",
					Value: "./player.xlsx",
					Usage: "data file path",
				},
			},
			Action: func(cCtx *cli.Context) error {
				return importPlayer(cCtx.String("data"), cCtx.String("db"))
			},
		}, {
			Name:    "importReferee",
			Aliases: []string{"ir"},
			Usage:   "simport referees",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "data",
					Value: "./referee.xlsx",
					Usage: "data file path",
				},
			},
			Action: func(cCtx *cli.Context) error {
				return importReferee(cCtx.String("data"), cCtx.String("db"))
			},
		}},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
