package main

import (
	"net/http"
	"ungraded-11/config"
	"ungraded-11/controller"
	"ungraded-11/handler"
	"ungraded-11/middlewares"

	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	db := config.InitDB()

	dbHandler := handler.NewDbHandler(db)
	userController := controller.NewUserController(dbHandler)

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.RequestLoggerWithConfig(middlewares.LogrusConfig()))

	users := e.Group("/users")
	{
		users.POST("/register", userController.Register)
		users.POST("/login", userController.Login)
	}

	e.Logger.Fatal(e.Start(":1323"))
}
