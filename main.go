package main

import (
	"apocalypse/database"
	"apocalypse/rest"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
)

func main() {
	conf, err := ReadConfig()
	if err != nil {
		log.Fatal().Err(err).Msgf("Error reading config: %q", err)
	}

	intPort, err := strconv.Atoi(conf.DataBaseConfig.Port)
	if err != nil {
		log.Fatal().Err(err).Msgf("Error convert Port to int: %q", err)
	}

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		conf.DataBaseConfig.Host,
		intPort,
		conf.DataBaseConfig.User,
		conf.DataBaseConfig.Password,
		conf.DataBaseConfig.Dbname,
		conf.DataBaseConfig.Sslmode,
	)

	db := database.Init(connectionString)
	if err != nil {
		log.Fatal().Err(err).Msgf("Error handling database connection: %q", err)
	}
	log.Info().Msg("Database connection established")
	log.Info().Msgf("Database connection established, connection string: %s", connectionString)
	rest.Router(db)
}

func handleDBConnection(db *sql.DB) (*sql.DB, error) {
	defer func() {
		if err := db.Close(); err != nil {
			log.Error().Err(err).Msg("Error closing database connection")
			return
		} else {
			log.Info().Msg("Database connection closed")
		}
	}()
	return db, nil
}
