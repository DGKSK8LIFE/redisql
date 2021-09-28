package utils

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// CTX is the global context for Redis
var CTX = context.Background()

// OpenRedis opens a redis connection with a desired address and password
func OpenRedis(redisAddress, redisPassword string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: redisPassword,
		DB:       0,
	})
	return rdb
}

// OpenMySQL opens a MySQL connection with a desired user, password, database name, host, and port
func OpenMySQL(user, password, database, host, port string) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// OpenPostgres opens a PostgreSQL connection with a desired user, password database name, host and port
func OpenPostgres(user, password, database, host, port string) (*sql.DB, error) {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, database)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}

// Convert is an internal function for Copy methods
func Convert(redisType, sqlUser, sqlPassword, sqlDatabase, sqlHost, sqlPort, sqlTable, redisAddr, redisPass, sqlType string) error {
	var db *sql.DB
	var err error

	switch sqlType {
	case "mysql":
		db, err = OpenMySQL(sqlUser, sqlPassword, sqlDatabase, sqlHost, sqlPort)
		if err != nil {
			return err
		}
	case "postgres":
		db, err = OpenPostgres(sqlUser, sqlPassword, sqlDatabase, sqlHost, sqlPort)
		if err != nil {
			return err
		}
	default:
		return errors.New("unsupported sql database type")
	}

	rdb := OpenRedis(redisAddr, redisPass)

	defer db.Close()
	defer rdb.Close()

	rows, err := db.Query(fmt.Sprintf(`SELECT * FROM %s`, sqlTable))
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

	index := 0
	switch redisType {
	case "string":
		for rows.Next() {
			if err = rows.Scan(scanArgs...); err != nil {
				return err
			}
			for i, col := range values {
				id := fmt.Sprintf("%s:%d:%s", sqlTable, index, columns[i])
				err := rdb.Set(CTX, id, string(col), 0).Err()
				if err != nil {
					return err
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
			id := fmt.Sprintf("%s:%d", sqlTable, index)
			err := rdb.RPush(CTX, id, fields).Err()
			if err != nil {
				return err
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
			id := fmt.Sprintf("%s:%d", sqlTable, index)
			err := rdb.HSet(CTX, id, rowMap).Err()
			if err != nil {
				return err
			}
			index += 1
		}
		if err = rows.Err(); err != nil {
			return err
		}
	}
	return nil
}
