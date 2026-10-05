package a

import (
	"context"
	"database/sql"
)

func _(ctx context.Context, db *sql.DB, term string) error {
	rows, err := db.QueryContext(ctx, "SELECT a FROM t WHERE t MATCH ?", term) // want "rows.Err must be checked"
	if err != nil {
		rows, err = db.QueryContext(ctx, "SELECT a FROM t WHERE a LIKE ?", term) // want "rows.Err must be checked"
		if err != nil {
			return err
		}
	}

	for rows.Next() {
		var a sql.NullString
		if err := rows.Scan(&a); err != nil {
			return err
		}
	}

	return rows.Close()
}
