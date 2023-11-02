package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Menu(c echo.Context) error {
	return c.Render(http.StatusOK, "menu.html", map[string]any{
		"forum": "Home",
	})
}


func InventoryChatRoom(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]any{
		"forum": "Inventory",
		"path": "inventory",
	})
}

func CrimesChatRoom(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]any{
		"forum": "Crimes",
		"path": "crime",
	})
}

func RandomChatRoom(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]any{
		"forum": "Random stuff",
		"path": "random",
	})
}