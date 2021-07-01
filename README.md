# redis-sql 

Migrate data from MySQL to Redis. 

## Example Usage

### Cli

```bash
go run main.go migrate --user=josh --password=joshmicheal15 --database=celebrities --table=celebrity
```

## Tech Stack 

- Go 
- Redis
- MySQL 

## Current Functionality and Limitations

- [x] Manual migration of MySQL tables to Redis via CLI
- [ ] Support for other SQL servers
- [ ] Support for migration of relational schema 
- [ ] Auto sync data
- [ ] Scheduling migrations
- [ ] TTL Support
- [ ] Migrations logs