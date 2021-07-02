# redisql 

Easily migrate data from MySQL to Redis. 

## Example Usage

### Cli

Installation: 

```bash
go install github.com/DGKSK8LIFE/redisql/cli
```

Usage:

### Library

Installation:

```bash
go get github.com/DGKSK8LIFE/redisql
```

Usage:

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