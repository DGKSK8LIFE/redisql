package migration

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Celebrity struct {
	Name string `sql:"name"`
	Age  uint8  `sql:"age"`
}

// Migrate takes an SQL table's rows and converts them into Redis hashes
func Migrate() error {
	db, err := sql.Open("mysql", "root:password@(localhost:3306)/celebrities")
	if err != nil {
		return err
	}

	rows, err := db.Query(`SELECT * FROM celebrity;`)
	if err != nil {
		return err
	}

	defer rows.Close()

	var celebrities []Celebrity
	for rows.Next() {
		var name string
		var age uint8
		err := rows.Scan(&name, &age)
		if err != nil {
			return err
		}

		fmt.Println(name, age)
		celebrities = append(celebrities, Celebrity{Name: name, Age: age})
	}
	if err := rows.Err(); err != nil {
		return err
	}
	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "",
	// 	DB:       0,
	// })
	return nil
}
