package main

import (
	"net/http"
	"os"
	"preview-w4/config"
	"preview-w4/controller"
	"preview-w4/models"
	"preview-w4/repository"

	"github.com/go-playground/validator"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type CustomValidator struct {
	validator *validator.Validate
}

func main() {
	db := config.InitDb()
	db.AutoMigrate(&models.Player{}, &models.User{})

	dbHandler := repository.NewDbHandler(db)
	playerController := controller.NewPlayerController(dbHandler)
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

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	players := e.Group("/players")
	{
		players.GET("", playerController.GetPlayers)
		players.POST("", playerController.GetPlayers)
		players.PUT("/:id", playerController.GetPlayers)
	}

	e.POST("/register", userController.Register)

	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
