### Migration

```bash
migrate create -ext sql -dir migrations create_table_users
```

### Execute

```bash
migrate -database "postgres://postgres:postgree@127.0.0.1:5433/ngc10?sslmode=disable" -path migrations up
```
