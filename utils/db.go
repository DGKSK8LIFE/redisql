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
		Addr:        redisAddress,
		Password:    redisPassword,
		DialTimeout: 1000000,
		DB:          0,
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

func OpenDB(cfg Config) (*sql.DB, error) {
	var db *sql.DB
	var err error

	switch cfg.SQLType {
	case "mysql":
		db, err = OpenMySQL(cfg.SQLUser, cfg.SQLPassword, cfg.SQLDatabase, cfg.SQLHost, cfg.SQLPort)
		if err != nil {
			return nil, err
		}
	case "postgres":
		db, err = OpenPostgres(cfg.SQLUser, cfg.SQLPassword, cfg.SQLDatabase, cfg.SQLHost, cfg.SQLPort)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("unsupported sql database type")
	}
	return db, nil
}
