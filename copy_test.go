package redisql

import (
	"database/sql"
	"flag"
	"testing"

	"os"

	"fmt"

	"github.com/DGKSK8LIFE/redisql/utils"
	"github.com/go-redis/redis/v8"
)

var createTableMySQL = `
	CREATE TABLE IF NOT EXISTS user (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		uuid VARCHAR(64) NOT NULL,
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
`

var createTablePostgres = `
	CREATE TABLE IF NOT EXISTS user_table (
    		id SERIAL PRIMARY KEY,
    		name character varying NOT NULL,
    		uuid uuid NOT NULL,
    		height character varying NOT NULL,
    		shoesize integer NOT NULL,
    		age integer NOT NULL,
		favorite_animal character varying(20) NOT NULL,
    		friends_count integer NOT NULL,
    		favorite_color character varying(50) NOT NULL,
    		favorite_food character varying(50) NOT NULL,
    		mobile_phone character varying(50) NOT NULL,
    		bio character varying NOT NULL
	)
`

var config Config
var rdb *redis.Client

func TestMain(m *testing.M) {
	config = Config{
		SQLType:     "",
		SQLUser:     "root",
		SQLPassword: "password",
		SQLDatabase: "users",
		SQLHost:     "localhost",
		SQLPort:     "3306",
		SQLTable:    "user",
		RedisAddr:   "localhost:6379",
		RedisPass:   "",
	}
	var rows int
	flag.StringVar(&config.SQLType, "db", "mysql", "postgres or mysql")
	flag.IntVar(&rows, "rows", 1000, "number of rows to insert before redisql tests run")
	flag.Parse()
	fmt.Println("Preparing Test...")

	var db *sql.DB
	var err error
	switch config.SQLType {
	case "mysql":
		db, err = utils.OpenMySQL(config.SQLUser, config.SQLPassword, config.SQLDatabase, config.SQLHost, config.SQLPort)
		if err != nil {
			panic(err)
		}
		_, err = db.Exec(createTableMySQL)
		if err != nil {
			panic(err)
		}
	case "postgres":
		config.SQLPort = "5432"
		config.SQLTable = "user_table"
		db, err = utils.OpenPostgres(config.SQLUser, config.SQLPassword, config.SQLDatabase, config.SQLHost, config.SQLPort)
		if err != nil {
			panic(err)
		}
		_, err = db.Exec(createTablePostgres)
		if err != nil {
			panic(err)
		}
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf(`DELETE FROM %s`, config.SQLTable))
	if err != nil {
		panic(err)
	}

	for i := 0; i < rows; i++ {
		_, err = db.Exec(fmt.Sprintf(`INSERT INTO %s (name, uuid, height, shoesize, age, bio, friends_count, favorite_animal, favorite_color, favorite_food, mobile_phone) VALUES ('martin', 'f8d1c837-719f-42a9-9a37-0e2ed7c0e458',  '5,9', 9, 15, 'Student and Developer', 100, 'horse', 'red', 'apple', '555-555-5555')`, config.SQLTable))
		if err != nil {
			panic(err)
		}
	}

	rdb = utils.OpenRedis(config.RedisAddr, config.RedisPass)
	defer rdb.Close()
	rdb.FlushAll(utils.CTX)

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestCopyToString(t *testing.T) {
	err := config.CopyToString()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestCopyToList(t *testing.T) {
	rdb.FlushAll(utils.CTX)
	err := config.CopyToList()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestCopyToHash(t *testing.T) {
	rdb.FlushAll(utils.CTX)
	err := config.CopyToHash()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
