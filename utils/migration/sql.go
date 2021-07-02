package migration

import (
	"database/sql"
	"fmt"
)

func OpenSQL(db *sql.DB, user, password, database string, err error) error {
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
