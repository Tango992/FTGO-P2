### Postgres Url
```bash
postgres://postgres:secret@localhost:5432/p2_ungraded_10?sslmode=disable
```

### Migrations
```bash
migrate -database "postgres://postgres:secret@localhost:5432/p2_ungraded_10?sslmode=disable" -path db/migrations up
```
```bash
migrate -database "postgres://postgres:secret@localhost:5432/p2_ungraded_10?sslmode=disable" -path db/migrations down
```