package main

import (
	"log"
	"preview-week3/config"
	"preview-week3/entity"
	"preview-week3/handler"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)


func main() {
	db := config.SetupDb()
	db.AutoMigrate(&entity.User{}, &entity.SocialMedia{}, &entity.Photo{}, &entity.Comment{})

	dbHandler := handler.NewDbHandler(db)
	userController := handler.NewUserController(dbHandler)

	r := gin.Default()
	user := r.Group("/users")
	{
		user.POST("/register", userController.Register)
		user.POST("/login", userController.Login)
	}



	log.Fatal(r.Run(":8080"))
}