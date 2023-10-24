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
			auth.GET("recipes", recipeController.GetAllRecipes)
			auth.POST("recipes", recipeController.PostRecipe)
		}
	}

	r.Run()
}