package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"ungraded-3/entity"

	"github.com/julienschmidt/httprouter"
)

type NewInventoryHandler struct {
	*sql.DB
}

var Msg = entity.Message{Message: "Invalid body input"}

func PanicFunc(w http.ResponseWriter, r *http.Request, i interface{}) {
	fmt.Fprintf(w, "Error occured: %v\n", i)
}

func (handler *NewInventoryHandler) GetInventories(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	
	rows, err := handler.Query(`
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
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(inventories)
}

func (handler *NewInventoryHandler) GetInventoriesId(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")

	param := p.ByName("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Msg)
		return
	}

	row := handler.QueryRow(`
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
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(v)
}

func (handler *NewInventoryHandler) PostInventory(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var newInvent entity.PostInventory
	if err := decoder.Decode(&newInvent); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Msg)
	}

	result, err := handler.Exec(`
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

func (handler *NewInventoryHandler) PutInventory(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

	res, err1 := handler.Exec(`
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

func (handler *NewInventoryHandler) DeleteInventory(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	
	param := p.ByName("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
	}

	res, err1 := handler.Exec(`
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

	json.NewEncoder(w).Encode(map[string]any{
		"status": fmt.Sprintf("Affected rows: %v", affectedRows),
		"values": fmt.Sprintf("Column ID: %v", id),
	})
}