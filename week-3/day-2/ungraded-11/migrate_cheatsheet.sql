'postgres://postgres:secret@localhost:5432/p2_ungraded_11?sslmode=disable'

migrate create -ext sql -dir migrations -seq insert_products_table

migrate -database ${POSTGRESQL_URL} -path migrations down

migrate -database ${POSTGRESQL_URL} -path migrations up
