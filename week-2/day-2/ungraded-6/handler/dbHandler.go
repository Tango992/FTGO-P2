package handler

import (
	"context"
	"database/sql"
	"net/http"
	"time"
	"ungraded-6/entity"
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

func (db DbHandler) FindUserInDb(credential *entity.Credential) (entity.User, *entity.Response) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := db.QueryRowContext(ctx, `
		SELECT id, email, password, name, age, occupation, role
		FROM users
		WHERE email = ?
	`, credential.Email)

	var user entity.User
	if err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.Age, &user.Occupation, &user.Role); err != nil {
		return entity.User{}, &entity.Response{
			Code: http.StatusUnauthorized,
			Message: "Invalid credentials",
			Data: nil,
		}
	}

	return user, nil
}

func (db DbHandler) FindRecipeInDb(data *entity.Recipe) (*entity.Response) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := db.QueryRowContext(ctx, `
		SELECT name, description, duration, rating
		FROM recipes
		WHERE id = ?
	`, data.Id)

	if err := row.Scan(&data.Name, &data.Description, &data.Duration, &data.Rating); err != nil {
		return &entity.Response{
			Code: http.StatusNotFound,
			Message: err.Error(),
			Data: nil,
		}
	}
	return nil
}

func (db DbHandler) FindAllRecipesInDb() ([]entity.Recipe, *entity.Response) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
		SELECT id, name, description, duration, rating
		FROM recipes
	`)
	if err != nil {
		return []entity.Recipe{}, &entity.Response{
			Code: http.StatusInternalServerError,
			Message: "Internal server error",
			Data: nil,
		}
	}
	defer rows.Close()

	var recipes []entity.Recipe
	for rows.Next() {
		var r entity.Recipe
		if err := rows.Scan(&r.Id, &r.Name, &r.Description, &r.Duration, &r.Rating); err != nil {
			return []entity.Recipe{}, &entity.Response{
				Code: http.StatusInternalServerError,
				Message: "Internal server error",
				Data: nil,
			}
		}

		recipes = append(recipes, r)
	}
	return recipes, nil
}

func (db DbHandler) InsertRecipeToDb(data entity.Recipe) *entity.Response {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx, `
		INSERT INTO recipes (name, description, duration, rating)
		VALUES (?,?,?,?)
	`, data.Name, data.Description, data.Duration, data.Rating)
	if err != nil {
		return &entity.Response{
			Code: http.StatusInternalServerError,
			Message: "Failed inserting into database",
			Data: nil,
		}
	}

	return nil
}

func (db DbHandler) DeleteRecipeFromDb(id int) *entity.Response {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := db.ExecContext(ctx, `
		DELETE FROM recipes
		WHERE id = ?
	`, id)
	if err != nil {
		return &entity.Response{
			Code: http.StatusInternalServerError,
			Message: "Internal server error",
			Data: nil,
		}
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return &entity.Response{
			Code: http.StatusNotFound,
			Message: "Recipe not found",
			Data: nil,
		}
	}
	return nil
}

func (db DbHandler) UpdateRecipeFromDb(data entity.Recipe) *entity.Response {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := db.ExecContext(ctx, `
		UPDATE recipes
		SET name = ?, description = ?, duration = ?, rating = ?
		WHERE id = ?
	`, data.Name, data.Description, data.Duration, data.Rating, data.Id)
	if err != nil {
		return &entity.Response{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
			Data: nil,
		}
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return &entity.Response{
			Code: http.StatusNotFound,
			Message: "Recipe not found",
			Data: nil,
		}
	}
	return nil
}