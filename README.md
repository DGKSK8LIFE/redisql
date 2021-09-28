<div align="center">
	<img src="./banner.png" width="100%"/> 
</div>

<div align="center">
	<a href="https://pkg.go.dev/github.com/DGKSK8LIFE/redisql">
		<img src="http://godoc.org/github.com/DGKSK8LIFE/redisql?status.svg"/>
	</a>
	<a href="https://discord.gg/RfZU9dp9ts">
		<img src="https://img.shields.io/discord/884452044566577182?color=738ADB&label=discord&style=flat-square" />
	</a>
</div>

---

## Example Usage

## CLI

### Installation and Configuration: 

```bash
go install github.com/DGKSK8LIFE/redisql/redisql
```

Create a YAML file with the following structure:

```yaml
sqltype:
sqluser: 
sqlpassword: 
sqldatabase:
sqlhost:
sqlport:
sqltable:
redisaddr:
redispass:
```

### Usage:

```bash
# copy to redis string
redisql copy -type=string -config=pathtofile.yml 

# copy to redis list
redisql copy -type=list -config=pathtofile.yml

# copy to redis hash
redisql copy -type=hash -config=pathtofile.yml
```

## Library

### Installation:

```bash
go get github.com/DGKSK8LIFE/redisql
```
### Usage:

```go
package main

import (
    "github.com/DGKSK8LIFE/redisql"
)

func main() {
	config := redisql.Config{
		SQLType:     "mysql",
		SQLUser:     "root",
		SQLPassword: "password",
		SQLDatabase: "users",
		SQLHost:     "localhost",
		SQLPort:     "3306",
		SQLTable:    "user",
		RedisAddr:   "localhost:6379",
		RedisPass:   "",
	}
	err := config.CopyToString()
	if err != nil {
		panic(err)
	}
}
```

### Other Methods:

```go
// copy to redis list
config.CopyToList()

// copy to redis hash
config.CopyToHash()
```

## Contributing 

Check out [CONTRIBUTING](CONTRIBUTING.md)

## Current Functionality and Limitations

- [x] Simple copying of entire MySQL tables to Redis via CLI and Go Module 
- [x] Improved logs (optional CLI output, improved formatting)	
- [x] Support for most commonly used Redis data types (strings, lists, hashes)
- [X] Support for Postgres
- [ ] Autosync