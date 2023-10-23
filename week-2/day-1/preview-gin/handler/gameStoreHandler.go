package handler

import (
	"net/http"
	"preview-week2-gin/entity"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StoreHandler struct {
	*DbHandler
}

func NewStoreHandler(dbHandler DbHandler) *StoreHandler {
	return &StoreHandler{
		DbHandler: &dbHandler,
	}
}

func (sh StoreHandler) GetBranches(c *gin.Context) {
	branches, err := sh.DbHandler.FindAllBranches()
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Error{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, branches)
}

func (sh StoreHandler) GetBranchById(c *gin.Context) {
	param := c.Param("id")

	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.Error{
			Code: http.StatusBadRequest,
			Message: "Invalid syntax",
		})
		return
	}

	branch, customError := sh.DbHandler.FindBranch(id)
	if customError != nil {
		c.JSON(customError.Code, *customError)
		return
	}
	c.JSON(http.StatusOK, branch)
}

func (sh StoreHandler) PostBranch(c *gin.Context) {
	var branch entity.Branch
	if err := c.ShouldBindJSON(&branch); err != nil {
		c.JSON(http.StatusBadRequest, entity.Error{
			Code: http.StatusBadRequest,
			Message: "Invalid syntax",
		})
		return
	}

	if customError := sh.DbHandler.AddBranchToDb(branch); customError != nil {
		c.JSON(customError.Code, *customError)
		return
	}

	c.JSON(http.StatusCreated, map[string]string{
		"message": "Branch posted",
	})
}

func (sh StoreHandler) PutBranch(c *gin.Context) {
	var branch entity.Branch
	var err error

	if err := c.ShouldBindJSON(&branch); err != nil {
		c.JSON(http.StatusBadRequest, entity.Error{
			Code: http.StatusBadRequest,
			Message: "Invalid syntax",
		})
		return
	}

	param := c.Param("id")

	if branch.Branch_id, err = strconv.Atoi(param); err != nil {
		c.JSON(http.StatusBadRequest, entity.Error{
			Code: http.StatusBadRequest,
			Message: "Invalid syntax",
		})
		return
	}

	if customError := sh.DbHandler.UpdateBranchToDb(branch); customError != nil {
		c.JSON(customError.Code, *customError)
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "Branch updated",
	})
}

func (sh StoreHandler) DeleteBranch(c *gin.Context) {
	param := c.Param("id")

	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.Error{
			Code: http.StatusBadRequest,
			Message: "Invalid syntax",
		})
		return
	}

	if customError := sh.DbHandler.DeleteBranchInDb(id); customError != nil {
		c.JSON(customError.Code, *customError)
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "Branch deleted",
	})
}
