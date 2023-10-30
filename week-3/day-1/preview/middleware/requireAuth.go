package middleware

import (
	"fmt"
	"preview-week3/helpers"
	"preview-week3/utils"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth() gin.HandlerFunc{
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			helpers.ErrJsonWriter(c, utils.ErrUnauthorized, "Please log in to access this page")
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("failed to verify token signature")
			}

			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			helpers.ErrJsonWriter(c, utils.ErrUnauthorized, err.Error())
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {			
			c.Set("user", map[string]any{
				"id": claims["id"],
				"email": claims["email"],
				"username": claims["username"],
			})

			c.Next()
			return
		}
		helpers.ErrJsonWriter(c, utils.ErrUnauthorized, "Please log in to access this page")
		c.Abort()
	}
}