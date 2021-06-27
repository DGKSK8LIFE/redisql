package migration

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	uuid "github.com/satori/go.uuid"
)

var ctx = context.Background()

/* Todo:
Add inferred/preset values for Migrate function to support any SQL schema (can do this via a describe query)
*/
// Migrate takes an SQL table and converts its rows into Redis hashes
func Migrate(user, password, database, table string) error {
	var db *sql.DB
	var err error
	if password != "" || password != " " {
		db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, password, database))
		if err != nil {
			return err
		}
	} else {
		db, err = sql.Open("mysql", fmt.Sprintf("%s@/%s", user, database))
		if err != nil {
			return err
		}
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	rows, err := db.Query(fmt.Sprintf(`SELECT * FROM %s;`, table))
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var age uint8
		err := rows.Scan(&name, &age)
		if err != nil {
			return err
		}

		id := uuid.NewV4()
		fmt.Println(id)
		rdb.HSet(ctx, id.String(), map[string]interface{}{"name": name, "age": age})
	}
	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}
