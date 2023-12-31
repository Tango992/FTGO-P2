package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"ungraded-4/entity"

	"github.com/julienschmidt/httprouter"
)



func (d DbHandler) GetVillains(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	
	rows, err := d.QueryContext(ctx, `
		SELECT v.ID, v.Name, u.Name, v.ImageURL
		FROM Villains v
		JOIN Universe u ON u.ID = v.Universe_id
	`)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.Response{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	defer rows.Close()

	var villains []entity.Villain
	for rows.Next() {
		var villain entity.Villain

		if err := rows.Scan(&villain.V_ID, &villain.V_Name, &villain.V_Universe, &villain.V_ImageURL); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(entity.Response{
				Code: http.StatusBadRequest,
				Message: err.Error(),
			})
		}
		villains = append(villains, villain)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(entity.Response{
		Code: http.StatusOK,
		Message: "Get all villains",
		Data: villains,
	})
}

func (d DbHandler) GetVillainById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")

	param := p.ByName("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.Response{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	row := d.QueryRowContext(ctx, `
		SELECT v.ID, v.Name, u.Name, v.ImageURL
		FROM Villains v
		JOIN Universe u ON u.ID = v.Universe_id
		WHERE v.ID = ?
	`, id)

	var villain entity.Villain
	if err := row.Scan(&villain.V_ID, &villain.V_Name, &villain.V_Universe, &villain.V_ImageURL); err != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(entity.Response{
			Code: http.StatusOK,
			Message: "Column out of range",
		})
		return
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(entity.Response{
		Code: http.StatusOK,
		Message: fmt.Sprintf("Get villain by id = %v", id),
		Data: villain,
	})
}