package main

import (
	"ugc-7/config"
	"ugc-7/handler"
	"ugc-7/initializers"
	"ugc-7/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
}

func main() {
	db := config.SetupDb()
	defer db.Close()

	dbHandler := handler.NewDbHandler(db)
	storeController := handler.NewStoreHandler(dbHandler)
	productController := handler.NewProductHandler(dbHandler)

	r := gin.Default()
	
	router := r.Group("/users")
	{
		router.POST("/register", storeController.Register)
		router.POST("/login", storeController.Login)
	}

	auth := r.Group("/products")
	auth.Use(middleware.RequireAuth())
	{
		auth.GET("", productController.Hello)
	}

	r.Run()
}