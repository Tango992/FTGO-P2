'postgres://postgres:secret@localhost:5432/p2_ungraded_11?sslmode=disable'

migrate -database ${POSTGRESQL_URL} -path migrations down   