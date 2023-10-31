package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)


func RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Please log in to access this page")
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("failed to verify token signature")
			}

			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())

		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			timeNow := time.Now().Unix()
			expiredTime := int64(claims["exp"].(float64))
			if timeNow > expiredTime {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token has expired, please log in again")

			}
			
			c.Set("user", map[string]any{
				"id": claims["id"],
				"username": claims["username"],
			})
			return next(c)
		}
		return echo.NewHTTPError(http.StatusUnauthorized, "Please log in to access this page")
	}
}