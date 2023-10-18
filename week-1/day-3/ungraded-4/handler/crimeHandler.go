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


func (d DbHandler) GetCrimeReports(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	rows, err := d.QueryContext(ctx, `
		SELECT 
			cr.ID, cr.Description, cr.Date, 
			h.ID, h.Name, u2.Name, h.Skill, h.ImageURL,
			v.ID, v.Name, u1.Name, v.ImageURL
		FROM CrimeReports cr
		JOIN Heroes h ON h.ID = cr.Hero_id
		JOIN Villains v ON v.ID = cr.Villain_id
		JOIN Universe u1 ON v.Universe_id = u1.ID
		JOIN Universe u2 ON h.Universe_id = u2.ID
	`)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"message": err.Error(),
		})
		return
	}
	defer rows.Close()

	var crimes []entity.GetCrimeReport
	for rows.Next() {
		var crime entity.GetCrimeReport

		err := rows.Scan(
			&crime.ID, &crime.Description, &crime.Date, 
			&crime.H_ID, &crime.H_Name, &crime.H_Universe, &crime.H_Skill, &crime.H_ImageURL,
			&crime.V_ID, &crime.V_Name, &crime.V_Universe, &crime.V_ImageURL,
		)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"message": err.Error(),
			})
			return
		}
		crimes = append(crimes, crime)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(crimes)
}


func (d DbHandler) GetCrimeReportsId(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
		SELECT 
			cr.ID, cr.Description, cr.Date, 
			h.ID, h.Name, u2.Name, h.Skill, h.ImageURL,
			v.ID, v.Name, u1.Name, v.ImageURL
		FROM CrimeReports cr
		JOIN Heroes h ON h.ID = cr.Hero_id
		JOIN Villains v ON v.ID = cr.Villain_id
		JOIN Universe u1 ON v.Universe_id = u1.ID
		JOIN Universe u2 ON h.Universe_id = u2.ID
		WHERE cr.ID = ?
	`, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"message": err.Error(),
		})
		return
	}

	var crime entity.GetCrimeReport
	err = row.Scan(
		&crime.ID, &crime.Description, &crime.Date, 
		&crime.H_ID, &crime.H_Name, &crime.H_Universe, &crime.H_Skill, &crime.H_ImageURL,
		&crime.V_ID, &crime.V_Name, &crime.V_Universe, &crime.V_ImageURL,
	)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(crime)
}


func (d DbHandler) PostCrimeReport(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	
	decoder := json.NewDecoder(r.Body)
	var newCrime entity.PostCrimeReport
	if err := decoder.Decode(&newCrime); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	_, err := d.ExecContext(ctx, `
		INSERT INTO CrimeReports (Hero_id, Villain_id, Description, Date)
		VALUES (?,?,?,?);
	`, newCrime.Hero_id, newCrime.Villain_id, newCrime.Description, newCrime.Date)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
}


func (d DbHandler) PutCrimeReport(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

	decoder := json.NewDecoder(r.Body)
	var putCrime entity.PostCrimeReport
	if err := decoder.Decode(&putCrime); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	result, err1 := d.ExecContext(ctx, `
		UPDATE CrimeReports
		SET 
			Description = ?,
			Date = ?
		WHERE ID = ?
	`, putCrime.Description, putCrime.Date, id)
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err1.Error(),
		})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (d DbHandler) DeleteCrimeReport(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

	result, err1 := d.ExecContext(ctx, `
		DELETE FROM CrimeReports
		WHERE ID = ?
	`, id)
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err1.Error(),
		})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}