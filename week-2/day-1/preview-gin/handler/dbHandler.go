package handler

import (
	"database/sql"
	"net/http"
	"preview-week2-gin/entity"
)

type DbHandler struct {
	*sql.DB
}

func NewDbHandler(db *sql.DB) *DbHandler {
	return &DbHandler{
		DB: db,
	}
}

func (db DbHandler) FindAllBranches() ([]entity.Branch, error) {
	ctx, cancel := AddContext()
	defer cancel()

	rows, err := db.QueryContext(ctx, `
		SELECT branch_id, name, location FROM branches
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var branches []entity.Branch
	for rows.Next() {
		var branch entity.Branch

		if err := rows.Scan(&branch.Branch_id, &branch.Name, &branch.Location); err != nil {
			return nil, err
		}
		branches = append(branches, branch)
	}
	return branches, nil
}

func (db DbHandler) FindBranch(id int) (entity.Branch, *entity.Error) {
	ctx, cancel := AddContext()
	defer cancel()

	row := db.QueryRowContext(ctx, `
		SELECT branch_id, name, location FROM branches
		WHERE branch_id = ?
	`, id)

	var branch entity.Branch
	if err := row.Scan(&branch.Branch_id, &branch.Name, &branch.Location); err != nil {
		return entity.Branch{}, &entity.Error{
			Code: http.StatusNotFound,
			Message: err.Error(),
		}
	}
	return branch, nil
}

func (db DbHandler) AddBranchToDb(branch entity.Branch) *entity.Error {
	ctx, cancel := AddContext()
	defer cancel()

	_, err := db.ExecContext(ctx, `
		INSERT INTO branches (name, location)
		VALUES (?,?)
	`, branch.Name, branch.Location)
	if err != nil {
		return &entity.Error{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func (db DbHandler) UpdateBranchToDb(branch entity.Branch) *entity.Error {
	ctx, cancel := AddContext()
	defer cancel()

	res, err := db.ExecContext(ctx, `
		UPDATE branches
		SET name = ?, location = ?
		WHERE branch_id = ?
	`, branch.Name, branch.Location, branch.Branch_id)
	if err != nil {
		return &entity.Error{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return &entity.Error{
			Code: http.StatusNotFound,
			Message: "Not found",
		}
	}
	return nil
}

func (db DbHandler) DeleteBranchInDb(id int) *entity.Error {
	ctx, cancel := AddContext()
	defer cancel()

	res, err := db.ExecContext(ctx, `
		DELETE FROM branches
		WHERE branch_id = ?
	`, id)
	if err != nil {
		return &entity.Error{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return &entity.Error{
			Code: http.StatusNotFound,
			Message: "Not found",
		}
	}
	return nil
}