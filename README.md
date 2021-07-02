# redisql 
[![GoDoc](http://godoc.org/github.com/DGKSK8LIFE/redisql?status.svg)](http://godoc.org/github.com/DGKSK8LIFE/redisql) 

Easily migrate data from MySQL to Redis. 

## Example Usage

### Cli

Installation: 

```bash
go install github.com/DGKSK8LIFE/redisql/redisql
```

Usage:

```bash
redisql migrate -user=josh1 -password=joshmark52 -database=celebrities -table=celebrity -redisaddr=localhost:6379 -redispass=joshmark52
```

### Library

Installation:

```bash
go get github.com/DGKSK8LIFE/redisql
```

Usage:

```go
package main

import (
    "github.com/DGKSK8LIFE/redisql"
)

func main() {
    err := redisql.Migrate("josh1", "joshmark52", "celebrities", "celebrity", "localhost:6379", "joshmark52")
    if err != nil {
        panic(err)
    }
}
```

## Tech Stack 

- Go 
- Redis
- MySQL 

## Current Functionality and Limitations

- [x] Manual migration of MySQL tables to Redis via CLI
- [ ] Support for migration of relational schema 
- [ ] Auto syncing data
- [ ] Scheduling migrations
- [ ] TTL argument for migrating data 
- [ ] Improved migration logs (log files and proper customization of formatting)
- [ ] Support for other SQL servers such as PostgreSQL and Microsoft SQL Server