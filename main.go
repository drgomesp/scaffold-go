package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var Version string

var Cfg Config

func init() {
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	MustLoadConfig()
}

func main() {
	if Version != "" {
		log.Info().Msgf("Version: %s\t", Version)
	}

	app := &cli.App{
		Name:  "greet",
		Usage: "fight the loneliness!",
		Commands: []*cli.Command{
			{
				Name:    "project",
				Aliases: []string{"p"},
				Usage:   "Scaffolds a new project",
				Action: func(ctx *cli.Context) error {
					spew.Dump(Cfg)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to run")
	}

}

func MustLoadConfig() {
	f, err := ioutil.ReadFile("./scaffold-go.yaml")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}

	err = yaml.Unmarshal(f, &Cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse config")
	}
}
