package main

import (
	"ungraded-6/config"
	"ungraded-6/handler"
	"ungraded-6/initializers"
	"ungraded-6/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
}

func main() {
	db := config.SetupDb()
	defer db.Close()

	dbHandler := handler.NewDbHandler(db)
	userController := handler.NewUserHandler(dbHandler)
	recipeController := handler.NewRecipeHandler(dbHandler)

	r := gin.Default()
	
	router := r.Group("/")
	router.Use(middleware.Logger)
	{
		router.POST("register", userController.Register)
		router.POST("login", userController.Login)

		auth := router.Group("/")
		auth.Use(middleware.RequireAuth)
		{
			superAdmin := auth.Group("")
			superAdmin.Use(middleware.RequireSuperAdmin)
			{
				superAdmin.POST("recipes", recipeController.PostRecipe)
				superAdmin.DELETE("recipes/:id", recipeController.DeleteRecipe)
			}
			auth.GET("recipes", recipeController.GetAllRecipes)
		}
	}

	r.Run()
}