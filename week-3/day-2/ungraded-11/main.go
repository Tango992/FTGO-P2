package main

import (
	"net/http"
	"ungraded-11/config"
	"ungraded-11/controller"
	"ungraded-11/handler"

	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
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

	log := logrus.New()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.WithFields(logrus.Fields{
				"URI":    values.URI,
				"status": values.Status,
			}).Info("request")

			return nil
		},
	}))

	users := e.Group("/users")
	{
		users.POST("/register", userController.Register)
		users.POST("/login", userController.Login)
	}

	e.Logger.Fatal(e.Start(":1323"))
}
