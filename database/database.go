package database

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func Init(config *DbInfo) *sql.DB {
	intPort, err := strconv.Atoi(config.Port)
	if err != nil {
		log.Fatal().Err(err).Msgf("Error convert Port to int: %q", err)
	}

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host,
		intPort,
		config.User,
		config.Password,
		config.Dbname,
		config.Sslmode,
	)
	
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Err(err).Msgf("Error opening database: %q", err)
	}

	if err = db.Ping(); err != nil {
		log.Err(err).Msgf("Error pinging database: %q", err)
	}
	return db
}
