package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger(c *gin.Context) {
	now := time.Now()
	date := now.Format("2006/01/02")
	time := now.Format("15:04:05")
	method := c.Request.Method
	path := c.Request.URL.Path
	fmt.Printf("%v - %v HTTP request sent to %v%v\n", date, time, method, path)
	c.Next()
}
