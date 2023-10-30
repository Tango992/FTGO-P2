package main

import (
	"log"
	"preview-week3/config"
	"preview-week3/entity"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)


func main() {
	db := config.SetupDb()
	db.AutoMigrate(&entity.User{}, &entity.SocialMedia{}, &entity.Photo{}, &entity.Comment{})

	r := gin.Default()

	log.Fatal(r.Run(":8080"))
}