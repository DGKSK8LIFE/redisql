package migration

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	// uuid "github.com/satori/go.uuid"
)

var ctx = context.Background()

// Migrate takes an SQL table and converts its rows into Redis hashes
func Migrate(user, password, database, table string) error {
	var db *sql.DB
	var err error

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

	defer db.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	defer rdb.Close()

	rows, err := db.Query(fmt.Sprintf(`SELECT * FROM %s;`, table))
	if err != nil {
		return err
	}

	defer rows.Close()

	columns, err := rows.Columns()
	for rows.Next() {
		if err != nil {
			return err
		}
		values := make([]sql.RawBytes, len(columns))
		err = rows.Scan(&columns)
		if err != nil {
			return err
		}

		// id := uuid.NewV4()
		// fmt.Println(id)
		// rdb.HSet(ctx, id.String(), map[string]interface{}{"name": name, "age": age})
	}
	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}
