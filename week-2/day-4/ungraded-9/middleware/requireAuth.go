package middleware

import (
	"os"
	"time"
	"ugc-9/handler"
	"ugc-9/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("Authorization")
		if err != nil {
			handler.WriteJsonErr(c, &utils.ErrUnauthorized, "Requires log in")
			c.Abort()
			return
		}
	
		token, parseErr := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, err
			}
			return []byte(os.Getenv("SECRET")), nil
		})
	
		if parseErr != nil {
			handler.WriteJsonErr(c, &utils.ErrUnauthorized, "Requires log in")
			c.Abort()
			return
		}
	
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				handler.WriteJsonErr(c, &utils.ErrUnauthorized, "Requires log in")
				c.Abort()
				return
			}
	
			c.Set("user", gin.H{
				"id": claims["id"],
				"name": claims["name"],
				"email": claims["email"],
				"type": claims["type"],
			})
	
			c.Next()
			return
		}

		handler.WriteJsonErr(c, &utils.ErrUnauthorized, "Requires log in")
		c.Abort()
	}
}