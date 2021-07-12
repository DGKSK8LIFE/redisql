package redisql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/DGKSK8LIFE/redisql/utils"
	_ "github.com/go-sql-driver/mysql"
)

// Configuration struct for redisql
type Config struct {
	SQLUser     string `yaml:"sqluser"`
	SQLPassword string `yaml:"sqlpassword"`
	SQLDatabase string `yaml:"sqldatabase"`
	SQLTable    string `yaml:"sqltable"`
	RedisAddr   string `yaml:"redisaddr"`
	RedisPass   string `yaml:"redispass"`
	Log         bool   `yaml:"log"`
}

var ctx = context.Background()

// Copy reads a desired SQL table's rows and writes them to Redis hashes
func (c Config) Copy() error {
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

	if c.Log {
		fmt.Println("\nRedis Hashes:\n")
	}
	index := 0
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return err
		}

		rowMap := make(map[string]string)
		for i, col := range values {
			rowMap[columns[i]] = string(col)
		}
		id := fmt.Sprintf("%s:%d", c.SQLTable, index)
		rdb.HSet(ctx, id, rowMap)
		if c.Log {
			utils.PrintRow(id, rowMap)
		}
		index += 1
	}
	if err := rows.Err(); err != nil {
		return err
	}
	fmt.Println("Migration Complete!")
	return nil
}
