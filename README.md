# tytan-api

### Migrations
```bash
goose create -dir ./migrations name_of_migration sql
```
```bash
goose -dir ./migrations sqlite3 ./database.db up
```
```bash
goose down -dir ./migrations sqlite3 ./database.db down
```