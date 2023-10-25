package handler

import (
	"context"
	"database/sql"
	"net/http"
	"time"
	"ugc-7/dto"
	"ugc-7/entity"
)

type DbHandler struct {
	*sql.DB
}

func NewDbHandler(db *sql.DB) *DbHandler {
	return &DbHandler{
		DB: db,
	}
}

func (db DbHandler) AddStoreToDb(u entity.Store) *dto.Response {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	_, err := db.ExecContext(ctx, `
		INSERT INTO stores (email, password, name, type)
		VALUES (?,?,?,?)
	`, u.Email, u.Password, u.Name, u.Type)
	if err != nil {
		return &dto.Response{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
			Data: nil,
		}
	}
	return nil
}

func (db DbHandler) FindUserInDb(credential *dto.Credential) (entity.Store, *dto.Response) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := db.QueryRowContext(ctx, `
		SELECT id, email, password, name, type
		FROM stores
		WHERE email = ?
	`, credential.Email)

	var store entity.Store
	if err := row.Scan(&store.Id, &store.Email, &store.Password, &store.Name, &store.Type); err != nil {
		return entity.Store{}, &dto.Response{
			Code: http.StatusUnauthorized,
			Message: "Invalid credentials",
			Data: nil,
		}
	}
	return store, nil
}
