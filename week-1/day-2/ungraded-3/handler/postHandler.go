package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ungraded-3/entity"

	"github.com/julienschmidt/httprouter"
)

func (app *App) PostInventory(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var newInvent entity.PostInventory
	decoder.Decode(&newInvent)

	result, err := app.Exec(`
		INSERT INTO Inventories (Name, Stock, Description, Status_id) 
		VALUES (?, ?, ?, ?)
	`, newInvent.Name, newInvent.Stock, newInvent.Description, newInvent.Status_id)
	if err != nil {
		panic(err)
	}
	affectedRows, _ := result.RowsAffected()

	json.NewEncoder(w).Encode(map[string]any{
		"status": fmt.Sprintf("Affected rows: %v", affectedRows),
		"values": newInvent,
	})
}