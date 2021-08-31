package redisql

import (
	"os"
	"testing"

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
	config := Config{
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
		t.Fail()
	}
	// _, err = db.Exec(`CREATE DATABASE IF NOT EXISTS ?`, config.SQLDatabase)
	// if err != nil {
	// 	t.Error(err)
	// 	t.Fail()
	// }
	// _, err = db.Exec(`
	// CREATE TABLE IF NOT EXISTS user (
	// 	id INT AUTO_INCREMENT PRIMARY KEY,
	// 	name VARCHAR(255) NOT NULL,
	// 	uuid VARCHAR(64) NOT NULL,
	// 	height VARCHAR(5) NOT NULL,
	// 	shoesize TINYINT NOT NULL,
	// 	age TINYINT NOT NULL,
	// 	bio TEXT NOT NULL,
	// 	friends_count TINYINT NOT NULL,
	// 	favorite_animal VARCHAR(20) NOT NULL,
	// 	favorite_color VARCHAR(10) NOT NULL,
	// 	favorite_food VARCHAR(20) NOT NULL,
	// 	mobile_phone VARCHAR(50) NOT NULL
	// );
	// `)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	for i := 0; i < 1000000; i++ {
		_, err = db.Exec(`INSERT INTO user VALUES (NULL, "martin", "f8d1c837-719f-42a9-9a37-0e2ed7c0e458",  "5'9", "9", 15, "Student and Developer", 100, "horse", "red", "apple", "555-555-5555")`)
		if err != nil {
			t.Error(err)
			t.Fail()
			os.Exit(1)
		}
	}
	// _, err = db.Exec(`
	// 	CREATE PROCEDURE million_row_insert()
	// 		BEGIN
	// 			SET @i = 0;
	// 			REPEAT
	// 				INSERT INTO $? (NULL, "martin", "f8d1c837-719f-42a9-9a37-0e2ed7c0e458", "5'9", "9", 15, "Student and Developer", 100, "horse", "red", "apple", "555-555-5555");
	// 				SET @i = @i + 1;
	// 			UNTIL @i = 1000000 END REPEAT;
	// 		END
	// `, config.SQLTable)
	// if err != nil {
	// 	t.Error(err)
	// 	t.Fail()
	// }
	// // _, err = db.Exec(`CALL million_row_insert()`)
	// if err != nil {
	// 	t.Error(err)
	// 	t.Fail()
	// }
	err = config.CopyToString()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
