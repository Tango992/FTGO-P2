package handler

import "database/sql"

type DbHandler struct {
	*sql.DB
}

func NewDbHandler(db *sql.DB) *DbHandler {
	return &DbHandler{
		DB: db,
	}
}