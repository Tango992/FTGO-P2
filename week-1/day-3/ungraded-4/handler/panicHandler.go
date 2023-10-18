package handler

import (
	"fmt"
	"net/http"
)

func PanicFunc(w http.ResponseWriter, r *http.Request, i interface{}) {
	fmt.Fprintf(w, "Error occured: %v\n", i)
}