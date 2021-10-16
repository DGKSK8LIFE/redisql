package utils

import (
	"database/sql"
	"fmt"

	"github.com/DGKSK8LIFE/redisql/utils/logging"
)

// CopyToString reads a desired SQL table's rows and writes them to Redis strings
func (c Config) CopyToString() error {
	logging.Log("Starting CopyToString", 1)
	if err := copyTable(c, "string"); err != nil {
		return err
	}
	return nil
}

// CopyToList reads a desired SQL table's rows and writes them to Redis lists
func (c Config) CopyToList() error {
	logging.Log("Starting CopyToList", 1)
	if err := copyTable(c, "list"); err != nil {
		return err
	}
	return nil
}

// CopyToHash reads a desired SQL table's rows and writes them to Redis hashes
func (c Config) CopyToHash() error {
	logging.Log("Starting CopyToHash", 1)
	if err := copyTable(c, "hash"); err != nil {
		return err
	}
	return nil
}

// copyTable is an internal function for Copy methods
func copyTable(cfg Config, redisType string) error {

	db, err := OpenDB(cfg)
	if err != nil {
		return err
	}
	logging.Log("SQL connection open", 1)

	rdb := OpenRedis(cfg.RedisAddr, cfg.RedisPass)
	logging.Log("Redis connection open", 1)

	defer db.Close()
	defer rdb.Close()

	rows, err := db.Query(fmt.Sprintf(`SELECT * FROM %s`, cfg.SQLTable))
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

	hook := PipelineHook{}
	rdb.AddHook(hook)

	pipe := rdb.Pipeline()
	defer pipe.Close()

	switch redisType {
	case "string":
		for rows.Next() {
			if err = rows.Scan(scanArgs...); err != nil {
				return err
			}
			for i, col := range values {
				id := fmt.Sprintf("%s:%d:%s", cfg.SQLTable, index, columns[i])
				res := pipe.Set(CTX, id, string(col), 0)
				if res.Err() != nil {
					return res.Err()
				}
			}
			index += 1
		}
		_, err = pipe.Exec(CTX)
		return err
	case "list":
		for rows.Next() {
			if err = rows.Scan(scanArgs...); err != nil {
				return err
			}
			fields := []string{}
			for _, col := range values {
				fields = append(fields, string(col))
			}
			id := fmt.Sprintf("%s:%d", cfg.SQLTable, index)
			err := pipe.RPush(CTX, id, fields).Err()
			if err != nil {
				return err
			}
			index += 1
		}
		_, err = pipe.Exec(CTX)
		return err
	case "hash":
		for rows.Next() {
			if err = rows.Scan(scanArgs...); err != nil {
				return err
			}
			rowMap := make(map[string]string)
			for i, col := range values {
				rowMap[columns[i]] = string(col)
			}
			id := fmt.Sprintf("%s:%d", cfg.SQLTable, index)
			err := pipe.HSet(CTX, id, rowMap).Err()
			if err != nil {
				return err
			}
			index += 1
		}
		_, err = pipe.Exec(CTX)
		if err != nil {
			return err
		}

		if err = rows.Err(); err != nil {
			return err
		}
	}
	logging.Log("Copying done", 1)
	return nil
}
