package handler

import (
	"context"
	"database/sql"
	"net/http"
	"time"
	"ungraded-5/entity"
)

type DbHandler struct {
	*sql.DB
}

func NewDbHandler(db *sql.DB) *DbHandler {
	return &DbHandler{
		DB: db,
	}
}

func (db DbHandler) AddUserToDb(u entity.User) *entity.Response {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	_, err := db.ExecContext(ctx, `
		INSERT INTO users (email, password, name, age, occupation, role)
		VALUES (?,?,?,?,?,?)
	`, u.Email, u.Password, u.Name, u.Age, u.Occupation, u.Role)
	if err != nil {
		return &entity.Response{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
			Data: nil,
		}
	}

	return nil
}

func (db DbHandler) FindHashInDb(credential *entity.Credential) (string, *entity.Response) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := db.QueryRowContext(ctx, `
		SELECT password
		FROM users
		WHERE email = ?
	`, credential.Email)

	var hash string
	if err := row.Scan(&hash); err != nil {
		return hash, &entity.Response{
			Code: http.StatusUnauthorized,
			Message: "Invalid credentials",
			Data: nil,
		}
	}

	return hash, nil
}
