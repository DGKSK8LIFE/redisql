package migration

import (
	"database/sql"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
)

type Celebrity struct {
	Name string `sql:"name"`
	Age  uint8  `sql:"age"`
}


// Migrate takes an SQL table's rows and converts them into Redis hashes
func Migrate() error {
	db, err := sql.Open("mysql", "root:password@/celebrities")
	if err != nil {
		return err
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return nil
}
