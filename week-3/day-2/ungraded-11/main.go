package main

import (
	"ungraded-11/config"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)

func main() {
    config.InitDB()
    
    e := echo.New()

    e.Logger.Fatal(e.Start(":1323"))
}