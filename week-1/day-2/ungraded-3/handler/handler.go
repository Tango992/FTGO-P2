package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"ungraded-3/entity"
)

type App struct {
	*sql.DB
}

var Msg = entity.Message{Message: "Invalid body input"}

func PanicHandler(w http.ResponseWriter, r *http.Request, i interface{}) {
	fmt.Fprintf(w, "Error occured: %v\n", i)
}