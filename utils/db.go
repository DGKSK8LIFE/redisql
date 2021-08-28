package utils

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// openRedis opens a redis connection with a desired address and password
func openRedis(redisAddress, redisPassword string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: redisPassword,
		DB:       0,
	})
	return rdb
}

// openSQL opens a MySQL connection with a desired user, password, and database name
func openSQL(user, password, database string) (*sql.DB, error) {
	switch password {
	case " ":
		db, err := sql.Open("mysql", fmt.Sprintf("%s@/%s", user, database))
		if err != nil {
			return nil, err
		}
		return db, nil
	default:
		db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, password, database))
		if err != nil {
			return nil, err
		}
		return db, nil
	}
}

// Convert is an internal function for Copy methods
func Convert(redisType, sqluser, sqlpassword, sqldatabase, sqltable, redisaddr, redispass string, log bool) error {
	db, err := openSQL(sqluser, sqlpassword, sqldatabase)
	if err != nil {
		return err
	}
	rdb := openRedis(redisaddr, redispass)

	defer db.Close()
	defer rdb.Close()

	rows, err := db.Query(fmt.Sprintf(`SELECT * FROM %s`, sqltable))
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
		fmt.Println("\nRedis Keys: \n")
	}
	index := 0
	switch redisType {
	case "string":
		for rows.Next() {
			if err = rows.Scan(scanArgs...); err != nil {
				return err
			}
			for i, col := range values {
				id := fmt.Sprintf("%s:%d:%s", sqltable, index, columns[i])
				err := rdb.Set(ctx, id, string(col), 0).Err()
				if err != nil {
					return err
				}
				if log {
					printKey(id, string(col))
				}
			}
			index += 1
		}
	case "list":
		for rows.Next() {
			if err = rows.Scan(scanArgs...); err != nil {
				return err
			}
			fields := []string{}
			for _, col := range values {
				fields = append(fields, string(col))
			}
			id := fmt.Sprintf("%s:%d", sqltable, index)
			err := rdb.RPush(ctx, id, fields).Err()
			if err != nil {
				return err
			}
			if log {
				printKey(id, fields)
			}
			index += 1
		}
	case "hash":
		for rows.Next() {
			if err = rows.Scan(scanArgs...); err != nil {
				return err
			}
			rowMap := make(map[string]string)
			for i, col := range values {
				rowMap[columns[i]] = string(col)
			}
			id := fmt.Sprintf("%s:%d", sqltable, index)
			err := rdb.HSet(ctx, id, rowMap).Err()
			if err != nil {
				return err
			}
			if log {
				printKey(id, rowMap)
			}
			index += 1
		}
		if err = rows.Err(); err != nil {
			return err
		}
	}
	fmt.Println("\nCopying Complete!")
	return nil
}

// AutoSync automatically calls Convert() if there is a change in the desired MySQL table
func AutoSync() {
}
