package utils

import (
	"database/sql"
	"fmt"
)

// OpenSQL opens a MySQL connection with a desired user, password, and database name
func OpenSQL(user, password, database string) (*sql.DB, error) {
	switch password {
	case " ":
		db, err := sql.Open("mysql", fmt.Sprintf("%s@/%s", user, database))
		if err != nil {
			return nil, err
		}
		return db, nil
	default:
		db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, password, database))
		if err != nil {
			return nil, err
		}
		return db, nil
	}
}
