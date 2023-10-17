package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"ungraded-3/entity"

	"github.com/julienschmidt/httprouter"
)

func (app *App) GetInventories(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	
	rows, err := app.Query(`
		SELECT i.Id, i.Name, i.Stock, i.Description, s.Name
		FROM Inventories i
		JOIN Status s ON I.Status_id = s.Id
	`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var inventories []entity.Inventory
	for rows.Next() {
		var v entity.Inventory

		if err := rows.Scan(&v.Id, &v.Name, &v.Stock, &v.Description, &v.Status); err != nil {
			panic(err)
		}
		inventories = append(inventories, v)
	}
	
	json.NewEncoder(w).Encode(inventories)
}

func (app *App) GetInventoriesId(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")

	param := p.ByName("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Msg)
		return
	}

	row := app.QueryRow(`
		SELECT i.Id, i.Name, i.Stock, i.Description, s.Name
		FROM Inventories i
		JOIN Status s ON I.Status_id = s.Id
		WHERE i.Id = ?
	`, id)

	var v entity.Inventory
	if err := row.Scan(&v.Id, &v.Name, &v.Stock, &v.Description, &v.Status); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"message": err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(v)
}