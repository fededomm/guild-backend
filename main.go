package main

import (
	"guild-be/src/database"
	"guild-be/docs"
	"guild-be/src/rest/controller"

	"github.com/rs/zerolog/log"
)

func main() {
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server guild server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v1/"
	docs.SwaggerInfo.Schemes = []string{"http"}

	conf, err := ReadConfig()
	if err != nil {
		log.Fatal().Err(err).Msgf("Error reading config: %q", err)
	}

	dbService := database.DBService{
		DB: database.Init(&conf.DataBaseConfig),
	}
	if err != nil {
		log.Fatal().Err(err).Msgf("Error handling database connection: %q", err)
	}
	log.Info().Msg("Database connection established")
	controller.Router(&dbService, conf.ValidArray.Rank, conf.ValidArray.Class)
}

