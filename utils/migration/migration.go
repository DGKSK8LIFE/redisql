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

// Migrate takes an SQL table and converts its rows into Redis hashes
/* Todo:
- use parameters instead of preset values for SQL and Redis methods
- allow for reverse migration through Migrate()
*/
func Migrate(user, database, table string, toRedis bool) error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s@/%s", user, database))
	if err != nil {
		return err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if toRedis {
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
	}
	return nil
}
