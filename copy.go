package redisql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/DGKSK8LIFE/redisql/utils"
	_ "github.com/go-sql-driver/mysql"
	uuid "github.com/satori/go.uuid"
)

// Configuration struct for redisql
type Config struct {
	SQLUser     string
	SQLPassword string
	SQLDatabase string
	SQLTable    string
	RedisAddr   string
	RedisPass   string
}

var ctx = context.Background()

// Copy reads a desired SQL table's rows and writes them to Redis hashes
func (c Config) Copy(log bool) error {
	db, err := utils.OpenSQL(c.SQLUser, c.SQLPassword, c.SQLDatabase)
	if err != nil {
		return err
	}
	defer db.Close()

	rdb := utils.OpenRedis(c.RedisAddr, c.RedisPass)
	defer rdb.Close()

	rows, err := db.Query(fmt.Sprintf(`SELECT * FROM %s;`, c.SQLTable))
	if err != nil {
		return err
	}

	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	if log {
		fmt.Println("Redis Hashes:\n")
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return err
		}

		rowMap := make(map[string]string)
		for i, col := range values {
			rowMap[columns[i]] = string(col)
		}
		id := (uuid.NewV4()).String()
		rdb.HSet(ctx, id, rowMap)
		if log {
			utils.PrintRow(id, rowMap)
		}
	}
	if err := rows.Err(); err != nil {
		return err
	}
	fmt.Println("\nMigration Complete!")
	return nil
}
