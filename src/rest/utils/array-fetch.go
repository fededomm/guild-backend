package utils

import (
	"database/sql"
	"fmt"
)

func FetchArrayByName(db *sql.DB, fetchArr []string, tablename string) ([]string, error) {
	query := fmt.Sprintf("SELECT name FROM %s", tablename)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	var arr []string
	for rows.Next() {
		var field string
		if err := rows.Scan(&field); err != nil {
			return nil, err
		}
		arr = append(arr, field)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return arr, nil
}
