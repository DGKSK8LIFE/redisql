package test

import (
	"testing"

	redisql "github.com/DGKSK8LIFE/redisql"
	"github.com/DGKSK8LIFE/redisql/utils"
)

func TestCopyToString(t *testing.T) {
	/*
		- Needs to create database, and table schema from `schema.sql`
		- Then inserts random or preset data 1,000,000 times (1 million table rows)
		- Then runs CopyToString() on the table
		- Goals:
			1. Benchmark various Copy() functions
			2. Look for edgecases and major issues in them
	*/
	config := redisql.Config{
		SQLUser:     "root",
		SQLPassword: "password",
		SQLDatabase: "users",
		SQLTable:    "user",
		RedisAddr:   "localhost:6379",
		RedisPass:   "",
		Log:         true,
	}
	db, err := utils.OpenSQL(config.SQLUser, config.SQLPassword, config.SQLDatabase)
	defer db.Close()
	if err != nil {
		t.Error(err)
	}
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS ?", config.SQLDatabase)
	if err != nil {
		t.Error(err)
	}
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS user (
    	id INT AUTO_INCREMENT PRIMARY KEY,
    	name VARCHAR(255) NOT NULL,
    	uuid VARCHAR(32) NOT NULL,
    	height VARCHAR(5) NOT NULL,
    	shoesize TINYINT NOT NULL,
    	age TINYINT NOT NULL,
    	bio TEXT NOT NULL,
    	friends_count TINYINT NOT NULL,
    	favorite_animal VARCHAR(20) NOT NULL,
    	favorite_color VARCHAR(10) NOT NULL,
    	favorite_food VARCHAR(20) NOT NULL,
    	mobile_phone VARCHAR(50) NOT NULL
	)
	`)
}
