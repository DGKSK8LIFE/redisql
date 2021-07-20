# redisql 
[![GoDoc](http://godoc.org/github.com/DGKSK8LIFE/redisql?status.svg)](http://godoc.org/github.com/DGKSK8LIFE/redisql) 

MySQL to Redis caching made easy

## Example Usage

### CLI

#### Installation: 

```bash
go install github.com/DGKSK8LIFE/redisql/redisql
```

#### Configuration:

Create a YAML file with the following structure:

```yaml
sqluser: 
sqlpassword: 
sqldatabase:
sqltable:
redisaddr:
redispass:
log:
```

#### Usage:

```bash
# copy to a redis string
redisql copy -type=string -config=pathtofile.yml 

# copy to a redis list
redisql copy -type=list -config=pathtofile.yml

# copy to a redis hash
redisql copy -type=hash -config=pathtofile.yml
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
		Log:         true,
	}
	err := config.CopyToString()
	if err != nil {
		panic(err)
	}
}
```

#### Other Methods:

```go
// copy to a list
config.CopyToList()

// copy to a hash
config.CopyToHash()
```

## Current Functionality and Limitations

- [x] Simple copying of entire MySQL tables to Redis via CLI and Go Module 
- [x] Improved logs (optional CLI output, improved formatting)	
- [x] Support for most commonly used Redis data types (strings, lists, hashes)
- [ ] Support for other SQL servers
