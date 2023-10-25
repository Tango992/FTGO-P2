package main

import (
	"ugc-8/config"
	"ugc-8/handler"
	"ugc-8/initializers"
	"ugc-8/middleware"

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