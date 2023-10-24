package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"ungraded-6/entity"

	"github.com/gin-gonic/gin"
)

type RecipeHandler struct {
	*DbHandler
}

func NewRecipeHandler(dbHandler *DbHandler) *RecipeHandler {
	return &RecipeHandler{
		DbHandler: dbHandler,
	}
}

func (rh RecipeHandler) GetAllRecipes(c *gin.Context) {
	recipes, dbErr := rh.DbHandler.FindAllRecipesInDb()
	if dbErr != nil {
		WriteJson(&c, *dbErr)
		return
	}

	WriteJson(&c, entity.Response{
		Code: http.StatusOK,
		Message: "Get all recipes",
		Data: recipes,
	})
}

func (rh RecipeHandler) PostRecipe(c *gin.Context) {
	var data entity.Recipe
	if err := c.ShouldBindJSON(&data); err != nil {
		WriteJson(&c, entity.Response{
			Code: http.StatusBadRequest,
			Message: "Invalid syntax",
			Data: nil,
		})
		return
	}

	if reflectErr := ValidateStruct(data); reflectErr != nil {
		WriteJson(&c, *reflectErr)
		return
	}

	if dbErr := rh.DbHandler.InsertRecipeToDb(data); dbErr != nil {
		WriteJson(&c, *dbErr)
		return
	}

	WriteJson(&c, entity.Response{
		Code: http.StatusOK,
		Message: "Recipe posted",
		Data: data,
	})
}

func (rh RecipeHandler) DeleteRecipe(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		WriteJson(&c, entity.Response{
			Code: http.StatusBadRequest,
			Message: "Invalid syntax",
			Data: nil,
		})
		return
	}

	if dbErr := rh.DbHandler.DeleteRecipeFromDb(id); dbErr != nil {
		WriteJson(&c, *dbErr)
		return
	}

	WriteJson(&c, entity.Response{
		Code: http.StatusOK,
		Message: fmt.Sprintf("Recipe deleted on id = %d", id),
	})
}