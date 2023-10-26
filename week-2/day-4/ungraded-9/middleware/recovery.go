package middleware

import (
	"fmt"
	"ugc-9/handler"
	"ugc-9/utils"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func(){
			if r := recover(); r != nil {
				fmt.Println("Recovery middleware caught panic")

				handler.WriteJsonErr(c, &utils.ErrInternalServer)
				c.Abort()
			}
		}()
		c.Next()
	}
}