package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerCountry(c *gin.Context) {
	url := "https://country-by-api-ninjas.p.rapidapi.com/v1/country?name=United%20States"

	req, _ := http.NewRequest("GET", url, nil)
	
	req.Header.Add("X-RapidAPI-Key", "7c8eb49c3bmsh29afe95d269234cp19aff5jsn7986cf16f9cb")
	req.Header.Add("X-RapidAPI-Host", "country-by-api-ninjas.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	defer res.Body.Close()
	
	var resData any
	if err := json.NewDecoder(res.Body).Decode(&resData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"message": resData,
	})
}