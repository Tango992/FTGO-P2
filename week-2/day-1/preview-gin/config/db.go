package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func SetDb() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/GameStore")
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}
	return db
}