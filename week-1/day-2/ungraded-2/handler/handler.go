package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"ungraded-2/entity"
)

type App struct {
	*sql.DB
}

func (app *App) Heroes(w http.ResponseWriter, r *http.Request) {
	rows, err := app.Query(`
		SELECT h.ID, h.Name, u.Name, h.Skill, h.ImageURL
		FROM Heroes h
		JOIN Universe u ON u.ID = h.Universe_id
	`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var heroes []entity.Heroes
	for rows.Next() {
		hero := entity.Heroes{}
		
		if err := rows.Scan(&hero.ID, &hero.Name, &hero.Universe, &hero.Skill, &hero.ImageURL); err != nil {
			panic(err)
		}
		heroes = append(heroes, hero)
	}

	json.NewEncoder(w).Encode(heroes)
}

func (app *App) Villains(w http.ResponseWriter, r *http.Request) {
	rows, err := app.Query(`
		SELECT v.ID, v.Name, u.Name, v.ImageURL
		FROM Villains v
		JOIN Universe u ON u.ID = v.Universe_id
	`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var villains []entity.Villains
	for rows.Next() {
		villain := entity.Villains{}

		if err := rows.Scan(&villain.ID, &villain.Name, &villain.Universe, &villain.ImageURL); err != nil {
			panic(err)
		}
		villains = append(villains, villain)
	}

	json.NewEncoder(w).Encode(villains)
}
