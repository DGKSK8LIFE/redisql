package redisql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

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

// CopyToHash reads a desired SQL table's rows and writes them to Redis hashes
func (c Config) CopyToHash() error {
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
		fmt.Println("\nRedis Keys:\n")
	}
	index := 0
	for rows.Next() {
		if err = rows.Scan(scanArgs...); err != nil {
			return err
		}

		rowMap := make(map[string]string)
		for i, col := range values {
			rowMap[columns[i]] = string(col)
		}

		id := fmt.Sprintf("%s:%d", c.SQLTable, index)
		rdb.HSet(ctx, id, rowMap)
		if c.Log {
			utils.PrintKey(id, rowMap)
		}
		index += 1
	}
	if err = rows.Err(); err != nil {
		return err
	}
	fmt.Println("Copying Complete!")
	return nil
}

// CopyToString reads a desired SQL table's rows and writes them to Redis strings
func (c Config) CopyToString(ttl uint) error {
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
		fmt.Println("\nRedis Keys:\n")
	}
	index := 0
	for rows.Next() {
		if err = rows.Scan(scanArgs...); err != nil {
			return err
		}

		for i, col := range values {
			id := fmt.Sprintf("%s:%d:%s", c.SQLTable, index, columns[i])
			rdb.Set(ctx, id, string(col), time.Duration(ttl)*time.Second)
			if c.Log {
				utils.PrintKey(id, string(col))
			}
		}
		index += 1
	}
	if err = rows.Err(); err != nil {
		return err
	}
	fmt.Println("Copying Complete!")
	return nil
}

// CopyToList reads a desired SQL table's rows and writes them to Redis lists
func (c Config) CopyToList() error {
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
		fmt.Println("\nRedis Keys:\n")
	}
	index := 0
	for rows.Next() {
		if err = rows.Scan(scanArgs...); err != nil {
			return err
		}

		fields := []string{}
		for _, col := range values {
			fields = append(fields, string(col))

		}
		id := fmt.Sprintf("%s:%d", c.SQLTable, index)
		rdb.RPush(ctx, id, fields)
		if c.Log {
			utils.PrintKey(id, fields)
		}
		index += 1
	}
	if err = rows.Err(); err != nil {
		return err
	}
	fmt.Println("Copying Complete!")
	return nil
}
