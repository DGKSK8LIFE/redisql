package migration

import (
	"database/sql"
	"fmt"
)

func OpenSQL(user, password, database string, db *sql.DB, err error) error {
	switch password {
	case " ":
		db, err = sql.Open("mysql", fmt.Sprintf("%s@/%s", user, database))
		if err != nil {
			return err
		}
	default:
		db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, password, database))
		if err != nil {
			return err
		}
	}
	return nil
}
