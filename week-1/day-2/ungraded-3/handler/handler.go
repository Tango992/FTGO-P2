package handler

import (
	"database/sql"
	"fmt"
	"net/http"
)

type App struct {
	*sql.DB
}

func PanicHandler(w http.ResponseWriter, r *http.Request, i interface{}) {
	fmt.Fprintf(w, "Error occured: %v\n", i)
}