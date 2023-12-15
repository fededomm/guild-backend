package main

import (
	"apocalypse/database"
	"apocalypse/rest"

	"github.com/rs/zerolog/log"
)

func main() {
	conf, err := ReadConfig()
	if err != nil {
		log.Fatal().Err(err).Msgf("Error reading config: %q", err)
	}

	db := database.Init(&conf.DataBaseConfig)
	if err != nil {
		log.Fatal().Err(err).Msgf("Error handling database connection: %q", err)
	}
	log.Info().Msg("Database connection established")
	rest.Router(db)
}

