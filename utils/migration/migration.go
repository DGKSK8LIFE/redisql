package migration

import (
	"context"
	"database/sql"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
)

var ctx = context.Background()

// Migrate takes an SQL table's rows and converts them into Redis hashes
func Migrate() error {
	db, err := sql.Open("mysql", "root@/celebrities")
	if err != nil {
		return err
	}

	rows, err := db.Query(`SELECT * FROM celebrity;`)
	if err != nil {
		return err
	}
	defer rows.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	for rows.Next() {
		var name string
		var age uint8
		err := rows.Scan(&name, &age)
		if err != nil {
			return err
		}
		rdb.HSet(ctx, "celebrity", map[string]interface{}{"name": name, "age": age})
	}
	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}
