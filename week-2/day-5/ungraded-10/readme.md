### Postgres URL
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

### Versioning

#### 000001
![v1](./dbdiagram/v1.svg)
#### 000002
![v2](./dbdiagram/v2.svg)
#### 000003 - 000004
![v3-4](./dbdiagram/v3-4.svg)