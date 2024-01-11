package utils

import (
	"context"
	"database/sql"
	"fmt"

	"go.opentelemetry.io/otel"
)

var trace = otel.Tracer("db-utils-guild-tracer")

func FetchArrayByName(ctx context.Context,db *sql.DB, fetchArr []string, tablename string) ([]string, error) {
	_, span := trace.Start(ctx, "FetchArrayByName")
	defer span.End()
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
