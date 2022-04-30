package main

import (
	"github.com/rs/zerolog/log"
)

var Version string
var BuildTime string

func main() {
	log.Info().Msgf("version: %s / time: %s", Version, BuildTime)
}
