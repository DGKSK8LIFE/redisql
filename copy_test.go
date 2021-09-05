package redisql

import (
	"testing"

	"os"

	"fmt"

	"github.com/DGKSK8LIFE/redisql/utils"
	"github.com/go-redis/redis/v8"
)

var insertString = `
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
var config Config
var rdb *redis.Client

func TestMain(m *testing.M) {
	fmt.Println("Preparing Test...")
	config = Config{
		SQLUser:     "root",
		SQLPassword: "password",
		SQLDatabase: "users",
		SQLTable:    "user",
		RedisAddr:   "localhost:6379",
		RedisPass:   "",
		Log:         false,
	}
	db, err := utils.OpenSQL(config.SQLUser, config.SQLPassword, config.SQLDatabase)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	rdb = utils.OpenRedis(config.RedisAddr, config.RedisPass)
	defer rdb.Close()
	_, err = db.Exec(insertString)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`DELETE FROM user`)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 1000; i++ {
		_, err = db.Exec(`INSERT INTO user VALUES (NULL, "martin", "f8d1c837-719f-42a9-9a37-0e2ed7c0e458",  "5'9", "9", 15, "Student and Developer", 100, "horse", "red", "apple", "555-555-5555")`)
		if err != nil {
			panic(err)
		}
	}
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestCopyToString(t *testing.T) {
	t.Log("Testing CopyToString...")
	err := config.CopyToString()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestCopyToList(t *testing.T) {
	t.Log("Testing CopyToList...")
	err := config.CopyToList()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestCopyToHash(t *testing.T) {
	t.Log("Testing CopyToHash...")
	err := config.CopyToHash()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
