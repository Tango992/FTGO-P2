package handler

import (
	"encoding/json"
	"net/http"
	"preview-week2/entity"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type StoreHandler struct {
	*DbHandler
}

func NewStoreHandler(dbHandler DbHandler) *StoreHandler {
	return &StoreHandler{
		DbHandler: &dbHandler,
	}
}

func (sh StoreHandler) GetBranches(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	branches, err := sh.DbHandler.FindAllBranches()
	if err != nil {
		JsonWriter(&w, http.StatusInternalServerError, entity.Error{
			Code: http.StatusInternalServerError,
			Message: "Internal server error",
		})
		return
	}

	JsonWriter(&w, http.StatusOK, branches)
}

func (sh StoreHandler) GetBranchById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	param := p.ByName("id")

	id, err := strconv.Atoi(param)
	if err != nil {
		JsonWriter(&w, http.StatusBadRequest, entity.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid syntax",
		})
		return
	}

	branch, customError := sh.DbHandler.FindBranch(id)
	if customError != nil {
		JsonWriter(&w, customError.Code, *customError)
		return
	}
	JsonWriter(&w, http.StatusOK, branch)
}

func (sh StoreHandler) PostBranch(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var branch entity.Branch

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&branch); err != nil {
		JsonWriter(&w, http.StatusBadRequest, entity.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid syntax",
		})
		return
	}

	if customError := sh.DbHandler.AddBranchToDb(branch); customError != nil {
		JsonWriter(&w, customError.Code, *customError)
		return
	}

	JsonWriter(&w, http.StatusCreated, map[string]string{
		"message": "Branch posted",
	})
}


func (sh StoreHandler) PutBranch(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var branch entity.Branch
	var err error

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&branch); err != nil {
		JsonWriter(&w, http.StatusBadRequest, entity.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid syntax",
		})
		return
	}

	param := p.ByName("id")
	if branch.Branch_id, err = strconv.Atoi(param); err != nil {
		JsonWriter(&w, http.StatusBadRequest, entity.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid syntax",
		})
		return
	}

	if customError := sh.DbHandler.UpdateBranchToDb(branch); customError != nil {
		JsonWriter(&w, customError.Code, *customError)
		return
	}

	JsonWriter(&w, http.StatusOK, map[string]string{
		"message": "Branch updated",
	})
}

func (sh StoreHandler) DeleteBranch(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	param := p.ByName("id")

	id, err := strconv.Atoi(param)
	if err != nil {
		JsonWriter(&w, http.StatusBadRequest, entity.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid syntax",
		})
		return
	}

	if customError := sh.DbHandler.DeleteBranchInDb(id); customError != nil {
		JsonWriter(&w, customError.Code, *customError)
		return
	}

	JsonWriter(&w, http.StatusOK, map[string]string{
		"message": "Branch deleted",
	})
}
