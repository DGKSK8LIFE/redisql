package utils

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql" //mysql driver neccessary to establish connection
	_ "github.com/lib/pq"              //postgresql driver necessary to establish connection
)

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

// openPostgres opens a PostgreSQL connection with a desired user, password database name, host and port
func OpenPostgres(user, password, database, host, port string) (*sql.DB, error) {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=prefer", user, password, host, port, database)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	//ping to check the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}

// Convert is an internal function for Copy methods
func Convert(redisType, sqluser, sqlpassword, sqldatabase, sqlhost, sqlport, sqltable, redisaddr, redispass, sqlType string, log bool) error {
	var db *sql.DB
	var err error

	switch sqlType {
	case "mysql":
		db, err = OpenMySQL(sqluser, sqlpassword, sqldatabase, sqlhost, sqlport)
		if err != nil {
			return err
		}
	case "postgres":
		db, err = OpenPostgres(sqluser, sqlpassword, sqldatabase, sqlhost, sqlport)
		if err != nil {
			return err
		}
	default:
		return errors.New("Sql database type not known!")
	}

	rdb := OpenRedis(redisaddr, redispass)

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
		fmt.Printf("\nRedis Keys: \n")
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
				err := rdb.Set(CTX, id, string(col), 0).Err()
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
			err := rdb.RPush(CTX, id, fields).Err()
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
			err := rdb.HSet(CTX, id, rowMap).Err()
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
	if log {
		fmt.Println("\nCopying Complete!")
	}
	return nil
}

// AutoSync automatically calls Convert() if there is a change in the desired MySQL table
func AutoSync(redisType, sqltype, sqluser, sqlpassword, sqldatabase, sqlhost, sqlport, sqltable, redisaddr, redispass string, log bool) error {
	for {
		Convert(redisType, sqltype, sqluser, sqlpassword, sqldatabase, sqlhost, sqlport, sqltable, redisaddr, redispass, log)
		time.Sleep(time.Second * 5)
	}
}
