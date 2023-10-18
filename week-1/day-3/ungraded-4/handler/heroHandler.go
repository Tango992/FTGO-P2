package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"ungraded-4/entity"

	"github.com/julienschmidt/httprouter"
)


func (d DbHandler) GetHeroes(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	rows, err := d.QueryContext(ctx, `
		SELECT h.ID, h.Name, u.Name, h.Skill, h.ImageURL
		FROM Heroes h
		JOIN Universe u ON u.ID = h.Universe_id
	`)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"message": err.Error(),
		})
		return
	}
	defer rows.Close()

	var heroes []entity.Hero
	for rows.Next() {
		var hero entity.Hero
		
		if err := rows.Scan(&hero.H_ID, &hero.H_Name, &hero.H_Universe, &hero.H_Skill, &hero.H_ImageURL); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"message": err.Error(),
			})
			return
		}
		heroes = append(heroes, hero)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(heroes)
}

func (d DbHandler) GetHeroById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")

	param := p.ByName("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"message": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	row := d.QueryRowContext(ctx, `
		SELECT h.ID, h.Name, u.Name, h.Skill, h.ImageURL
		FROM Heroes h
		JOIN Universe u ON u.ID = h.Universe_id
		WHERE h.ID = ?
	`, id)

	var hero entity.Hero
	if err := row.Scan(&hero.H_ID, &hero.H_Name, &hero.H_Universe, &hero.H_Skill, &hero.H_ImageURL); err != nil {
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(map[string]any{
			"message": err.Error(),
		})
		return
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(hero)
}