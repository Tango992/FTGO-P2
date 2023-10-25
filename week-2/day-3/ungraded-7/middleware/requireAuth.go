package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("Authorization")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"message": "Unauthorized access",
			})
			return
		}
	
		token, parseErr := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, err
			}
			return []byte(os.Getenv("SECRET")), nil
		})
	
		if parseErr != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"message": "Unauthorized access",
			})
			return
		}
	
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"code": http.StatusUnauthorized,
					"message": "Unauthorized access",
				})
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
	
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"message": "Unauthorized access",
		})
	}
}