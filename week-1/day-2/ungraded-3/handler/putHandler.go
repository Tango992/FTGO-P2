package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"ungraded-3/entity"

	"github.com/julienschmidt/httprouter"
)

func (app *App) PutInventory(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")

	param := p.ByName("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(r.Body)
	var values entity.PutInventory
	if err := decoder.Decode(&values); err != nil {
		panic(err)
	}

	res, err1 := app.Exec(`
		UPDATE Inventories
		SET  Stock = ?, Status_id = ?
		WHERE Id = ?
	`, values.Stock, values.Status_id, id)
	if err1 != nil {
		panic(err1)
	}

	affectedRows, err2 := res.RowsAffected()
	if err2 != nil {
		panic(err2)
	}

	json.NewEncoder(w).Encode(map[string]any{
		"status": fmt.Sprintf("Affected rows: %v", affectedRows),
		"values": values,
	})
}