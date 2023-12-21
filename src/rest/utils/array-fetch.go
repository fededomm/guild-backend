package utils

import (
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"
)


func FetchArray(db *sql.DB, fetchArr []string, tablename string) ([]string, error){
	query := fmt.Sprintf("SELECT name FROM %s", tablename)
	rows, err := db.Query(query)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return nil, err
	}
	
	var arr []string
	for rows.Next() {
		var field string
		err := rows.Scan(&field)
		if err != nil {
			log.Err(err).Msg(err.Error())
			return nil, err
		}
		arr = append(arr, field)
	}
	
	err = rows.Err()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return nil, err
	}
	
	return arr, nil
}