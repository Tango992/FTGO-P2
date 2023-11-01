package main

import (
	"api-ninja/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", handler.HandlerCountry)
	
	r.Run(":8082")
}