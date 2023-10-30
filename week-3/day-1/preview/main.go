package main

import (
	"log"
	"preview-week3/config"
	"preview-week3/entity"
	"preview-week3/handler"
	"preview-week3/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)


func main() {
	db := config.SetupDb()
	db.AutoMigrate(&entity.User{}, &entity.SocialMedia{}, &entity.Photo{}, &entity.Comment{})

	dbHandler := handler.NewDbHandler(db)
	userController := handler.NewUserController(dbHandler)
	photoController := handler.NewPhotoController(dbHandler)

	r := gin.Default()
	user := r.Group("/users")
	{
		user.POST("/register", userController.Register)
		user.POST("/login", userController.Login)
	}

	photos := r.Group("/photos")
	photos.Use(middleware.RequireAuth())
	{
		photos.GET("", photoController.GetPhotos)
		photos.POST("", photoController.PostPhoto)
	}


	log.Fatal(r.Run(":8080"))
}