package main

import (
	"ugc-9/config"
	"ugc-9/handler"
	"ugc-9/initializers"
	"ugc-9/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
}

func main() {
	db := config.SetupDb()

	dbHandler := handler.NewDbHandler(db)
	storeController := handler.NewStoreHandler(dbHandler)
	productController := handler.NewProductHandler(dbHandler)

	r := gin.Default()
	
	r.Use(middleware.Recovery())
	router := r.Group("/users")
	{
		router.POST("/register", storeController.Register)
		router.POST("/login", storeController.Login)
	}

	auth := r.Group("/products")
	auth.Use(middleware.RequireAuth())
	{
		auth.PUT("/:id", productController.UpdateProductById)
		auth.DELETE("/:id", productController.DeleteProductById)
		auth.GET("/:id", productController.GetProductById)
		auth.GET("", productController.GetAllProducts)
		auth.POST("", productController.PostProduct)
	}

	r.Run()
}