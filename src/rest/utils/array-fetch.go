package utils

import (
	"database/sql"
	"fmt"
)


func FetchArray(db *sql.DB, fetchArr []string, tablename string) ([]string, error){
	query := fmt.Sprintf("SELECT name FROM %s", tablename)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	
	var arr []string
	for rows.Next() {
		var field string
		err := rows.Scan(&field)
		if err != nil {
			return nil, err
		}
		arr = append(arr, field)
	}
	
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	
	return arr, nil
}