package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"ungraded-13/dto"
	"ungraded-13/repository"

	"github.com/labstack/echo/v4"
)

type StoreController struct {
	repository.DbHandler
}

func NewStoreController(dbHandler repository.DbHandler) StoreController {
	return StoreController{
		DbHandler: dbHandler,
	}
}

// Stores        godoc
// @Summary      View all stores
// @Tags         stores
// @Produce      json
// @Success      200  {object}  dto.StoreResponse
// @Failure      400  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /stores [get]
func (sc StoreController) GetStores(c echo.Context) error {
	stores, dbErr := sc.DbHandler.FindAllStoresInDb()
	if dbErr != nil {
		return dbErr
	}
	return c.JSON(http.StatusOK, dto.Response{
		Message: "Get all stores",
		Data: stores,
	})
}

// Stores by Id  godoc
// @Summary      Get store by Id
// @Tags         stores
// @Produce      json
// @Param        store_id path int true "Store Id"
// @Success      200  {object}  dto.StoreByIdResponse
// @Failure      400  {object}  dto.Error
// @Failure      500  {object}  dto.Error
// @Router       /stores/{store_id} [get]
func (sc StoreController) GetStoreById(c echo.Context) error {
	storeIdTemp := c.Param("id")
	storeId, err := strconv.Atoi(storeIdTemp)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	storeData, dbErr := sc.DbHandler.FindStoreInDb(storeId)
	if dbErr != nil {
		return dbErr
	}
	url := "https://weather-by-api-ninjas.p.rapidapi.com/v1/weather"

	lat := fmt.Sprintf("%v", storeData.Lat)
	long := fmt.Sprintf("%v", storeData.Long)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-RapidAPI-Key", os.Getenv("RAPID_API_KEY"))
	req.Header.Add("X-RapidAPI-Host", "weather-by-api-ninjas.p.rapidapi.com")

	q := req.URL.Query()
	q.Add("lat", lat)
	q.Add("lon", long)
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var weatherData any
	if err := json.NewDecoder(res.Body).Decode(&weatherData); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	storeDetail := dto.StoreDetail{
		StoreWithSales: storeData,
		Weather: weatherData,
	}
	
	response := dto.Response{
		Message: "Get store by ID",
		Data: storeDetail,
	}
	
	return c.JSON(http.StatusOK, response)
}