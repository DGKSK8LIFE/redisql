package migration

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/DGKSK8LIFE/redisql/utils"
	_ "github.com/go-sql-driver/mysql"
	uuid "github.com/satori/go.uuid"
)

var ctx = context.Background()

// Migrate takes an SQL table and converts its rows into Redis hashes
func Migrate(user, password, database, table, redisAddress, redisPassword string) error {
	db, err := utils.OpenSQL(user, password, database)
	if err != nil {
		return err
	}
	defer db.Close()

	rdb := utils.OpenRedis(redisAddress, redisPassword)
	defer rdb.Close()

	rows, err := db.Query(fmt.Sprintf(`SELECT * FROM %s;`, table))
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

	fmt.Println("Redis Keys:\n")
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return err
		}

		rowMap := make(map[string]interface{})
		for i, col := range values {
			rowMap[columns[i]] = string(col)
		}
		id := uuid.NewV4()
		fmt.Println(id)
		rdb.HSet(ctx, id.String(), rowMap)
	}
	if err := rows.Err(); err != nil {
		return err
	}
	fmt.Println("\nMigration Complete!")
	return nil
}
