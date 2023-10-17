package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *App) DeleteInventory(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	param := p.ByName("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
	}

	res, err1 := app.Exec(`
		DELETE FROM Inventories
		WHERE Id = ?
	`, id)
	if err1 != nil {
		panic(err1)
	}

	affectedRows, err2 := res.RowsAffected()
	if err2 != nil {
		panic(err2)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status": fmt.Sprintf("Affected rows: %v", affectedRows),
		"values": fmt.Sprintf("Column ID: %v", id),
	})
}