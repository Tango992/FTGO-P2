package config

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func SetupDb() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("DB"))
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}
	return db
}