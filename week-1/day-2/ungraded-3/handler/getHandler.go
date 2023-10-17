package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"ungraded-3/entity"

	"github.com/julienschmidt/httprouter"
)

func (app *App) GetInventories(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
	param := p.ByName("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
	}

	row := app.QueryRow(`
	SELECT i.Id, i.Name, i.Stock, i.Description, s.Name
	FROM Inventories i
	JOIN Status s ON I.Status_id = s.Id
	WHERE i.Id = ?
	`, id)

	var v entity.Inventory
	if err := row.Scan(&v.Id, &v.Name, &v.Stock, &v.Description, &v.Status); err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(v)
}