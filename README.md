# redisql 
[![GoDoc](http://godoc.org/github.com/DGKSK8LIFE/redisql?status.svg)](http://godoc.org/github.com/DGKSK8LIFE/redisql) 

MySQL to Redis caching made easy

## Example Usage

### CLI

#### Installation: 

```bash
go install github.com/DGKSK8LIFE/redisql/redisql
```

#### Usage:

```bash
redisql copy -user=josh1 -password=joshmark52 -database=celebrities -table=celebrity -redisaddr=localhost:6379 -redispass=joshmark52
```

### Library

#### Installation:

```bash
go get github.com/DGKSK8LIFE/redisql
```

#### Usage:

```go
package main

import (
    "github.com/DGKSK8LIFE/redisql"
)

func main() {
	config := redisql.Config{
		SQLUser:     "josh",
		SQLPassword: "joshmark52",
		SQLDatabase: "celebrities",
		SQLTable:    "celebrity",
		RedisAddr:   "localhost:6379",
		RedisPass:   "joshmark52",
	}
	err := config.Copy()
	if err != nil {
		panic(err)
	}
}
```

## Current Functionality and Limitations

- [x] Simple copying of entire MySQL tables to Redis via CLI and Go Module 
- [ ] Improved logs (log files and proper customization of formatting)
- [ ] Support for multiple Redis data types (lists, sets, etc)
- [ ] Support for custom SQL queries
- [ ] Auto sync
- [ ] Support for other SQL servers
- [ ] Representation of SQL relations within Redis (limited) 