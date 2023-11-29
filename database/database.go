package database

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func Init(cn string) *sql.DB {
	db, err := sql.Open("postgres", cn)
	if err != nil {
		log.Err(err).Msgf("Error opening database: %q", err)
	}

	if err = db.Ping(); err != nil {
		log.Err(err).Msgf("Error pinging database: %q", err)
	}
	return db
}
