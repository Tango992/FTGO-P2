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
	controller := handler.NewUserHandler(dbHandler)

	r := gin.Default()
	
	router := r.Group("/")
	router.Use(middleware.Logger)
	{
		router.POST("/register", controller.Register)
		router.POST("/login", controller.Login)
	}

	r.Run()
}